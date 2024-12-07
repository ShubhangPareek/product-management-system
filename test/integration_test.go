package test

import (
	"product-management-system/routes"
	"testing"

	"github.com/gin-gonic/gin"
)

func mockHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "success"})
}

func TestIntegration(t *testing.T) {
	router := gin.Default()

	// Pass mock handlers to satisfy the updated RegisterRoutes signature
	routes.RegisterRoutes(router, mockHandler, mockHandler, mockHandler)

	t.Log("Integration test executed successfully!")
}
