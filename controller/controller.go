package controller

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	filename = "./logs.txt"
)

func ProxyServer() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
		defer cancel()
		queryURL := c.Query("url")
		if queryURL == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing query parameter 'url'"})
			return
		}
		targetUrl, err := url.Parse(queryURL)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error parsing target URL: %s", err)})
			return
		}
		client := &http.Client{}
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, targetUrl.String(), nil)
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

		_, err = io.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error reading response body: %s", err)})
			return
		}
		fmt.Println(c.ClientIP())
		fmt.Println(resp.Request.Host)
		fmt.Println(c.Request.UserAgent())
		fmt.Println(resp.Request.Method)
		c.Redirect(http.StatusTemporaryRedirect, queryURL)
	}
}
