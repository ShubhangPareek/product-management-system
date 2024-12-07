package controllers

import (
	"net/http"
	"net/http/httptest"
	"product-management-system/config"
	"product-management-system/routes"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCreateProduct(t *testing.T) {
	config.ConnectDB()

	router := gin.Default()
	// Pass actual handlers from controllers
	routes.RegisterRoutes(router, CreateProduct, GetProductByID, GetProducts)

	t.Run("Valid Product", func(t *testing.T) {
		req, _ := http.NewRequest("POST", "/products", nil)
		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		if recorder.Code != http.StatusBadRequest {
			t.Errorf("Expected status code 400, but got %d", recorder.Code)
		}
	})
}
func TestGetProducts(t *testing.T) {
	config.ConnectDB()

	router := gin.Default()
	routes.RegisterRoutes(router, CreateProduct, GetProductByID, GetProducts)

	req, _ := http.NewRequest("GET", "/products", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code 200, but got %d", recorder.Code)
	}
}
func TestGetProductByIDWithInvalidID(t *testing.T) {
	config.ConnectDB()

	router := gin.Default()
	routes.RegisterRoutes(router, CreateProduct, GetProductByID, GetProducts)

	req, _ := http.NewRequest("GET", "/products/invalid", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	// Adjust the expected response code based on the actual behavior
	if recorder.Code != http.StatusNotFound {
		t.Errorf("Expected status code 404, got %d", recorder.Code)
	}
}
func TestGetProductsInvalidQueryParams(t *testing.T) {
	config.ConnectDB()
	router := gin.Default()
	routes.RegisterRoutes(router, CreateProduct, GetProductByID, GetProducts)

	// Test with invalid query parameters
	req, _ := http.NewRequest("GET", "/products?min_price=invalid", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected 400, got %d", w.Code)
	}
}
func TestCreateProductInvalidData(t *testing.T) {
	config.ConnectDB()
	router := gin.Default()
	routes.RegisterRoutes(router, CreateProduct, GetProductByID, GetProducts)

	invalidPayload := `{"user_id":"abc", "product_price":"xyz"}`
	req, _ := http.NewRequest("POST", "/products", strings.NewReader(invalidPayload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected 400, got %d", w.Code)
	}
}
