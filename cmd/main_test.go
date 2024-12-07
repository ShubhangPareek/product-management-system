package main

import (
	"net/http"
	"net/http/httptest"
	"product-management-system/routes"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMainInitialization(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mockHandler := func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "test"})
	}
	routes.RegisterRoutes(router, mockHandler, mockHandler, mockHandler)

	req, _ := http.NewRequest("GET", "/products", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}
