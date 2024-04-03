package controller

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/Uttkarsh-Raj/Proxie/model"
	"github.com/gin-gonic/gin"
)

func ProxyServer() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
		defer cancel()
		// Get the query URL
		queryURL := c.Query("url")
		if queryURL == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing query parameter 'url'"})
			return
		}

		// Convert to a uri
		targetUrl, err := url.Parse(queryURL)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error parsing target URL: %s", err)})
			return
		}

		// Check the last requested time < 5sec
		err = RateLimitChecker(c.ClientIP())
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Create a new client to send the request using this servers context
		// Request the target url
		// client := http.DefaultTransport
		client := &http.Client{}
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, targetUrl.String(), c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error creating request: %s", err)})
			return
		}

		resp, err := client.Do(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error connecting to the destination server: %s", err)})
			return
		}
		defer resp.Body.Close()

		// Copy the response and the headers received
		c.Status(resp.StatusCode)
		for k, v := range resp.Header {
			c.Header(k, v[0])
		}

		// Copy the response body to the current response
		if _, err := io.Copy(c.Writer, resp.Body); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error connecting to the destination server: %s", err)})
			return
		}

		model.RateLimiter.Requests[c.ClientIP()] = time.Now()

		// For testing pourpose
		if queryURL == "https://www.tests.com/" {
			c.String(http.StatusOK, "<html><h1>Test Page</h1></html>")
		}
	}
}

func RateLimitChecker(clientIP string) error {
	if time.Since(model.RateLimiter.Requests[clientIP]) < time.Second*5 {
		return fmt.Errorf("error: Please wait for %s seconds before next request", ((5 * time.Second) - (time.Since(model.RateLimiter.Requests[clientIP]).Abs().Round(time.Second))))
	}
	return nil
}
