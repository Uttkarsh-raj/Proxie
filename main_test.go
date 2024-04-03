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

	// Test the proxy servers function
	req, _ := http.NewRequest("GET", "/?url=https://www.tests.com/", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "https://www.tests.com/", w.Header().Get("Location"))

	// Check the Rate Limiter
	err := `{"error":"error: Please wait for 5s seconds before next request"}`
	req, _ = http.NewRequest("GET", "/?url=https://www.tests.com/", nil)
	w = httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, err, w.Body.String())
}
