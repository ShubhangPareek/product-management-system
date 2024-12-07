package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRegisterRoutes(t *testing.T) {
	r := gin.Default()

	RegisterRoutes(
		r,
		func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "POST success"}) },
		func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "GET by ID success"}) },
		func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "GET all success"}) },
	)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/products", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}
