package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func BenchmarkGetProductByID(b *testing.B) {
	router := gin.Default()
	router.GET("/products/:id", GetProductByID)

	for i := 0; i < b.N; i++ {
		req, _ := http.NewRequest("GET", "/products/1", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
	}
}
