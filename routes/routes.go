package routes

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up the API endpoints
func RegisterRoutes(
	r *gin.Engine,
	createProduct gin.HandlerFunc,
	getProductByID gin.HandlerFunc,
	getProducts gin.HandlerFunc,
) {
	r.POST("/products", createProduct)
	r.GET("/products/:id", getProductByID)
	r.GET("/products", getProducts)
}
