package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Uttkarsh-Raj/Proxie/controller"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func TestProxyServer(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/", controller.ProxyServer())

	// Test the proxy server's function
	req, _ := http.NewRequest("GET", "/?url=https://www.tsetit.com/", nil)
	req.Header.Add("X-Forwarded-For", "123.123.123.123")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	expectedResponse := "<html><h1>Test Page</h1></html>"
	assert.Equal(t, expectedResponse, w.Body.String())

	// Test the rate limiter
	err := `{"error":"error: Please wait for 5s seconds before next request"}`
	req, _ = http.NewRequest("GET", "/?url=https://www.tsetit.com/", nil)
	req.Header.Add("X-Forwarded-For", "123.123.123.123")

	w = httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, err, w.Body.String())
}
