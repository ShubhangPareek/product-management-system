package main

import (
	"product-management-system/config"
	"product-management-system/routes"
	"product-management-system/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize MongoDB
	config.ConnectDB()

	// Initialize Logger
	utils.InitLogger()
	utils.Logger.Info("Logger initialized")

	// Start Router
	r := gin.Default()
	routes.RegisterRoutes(r)
	utils.Logger.Info("Server starting on port 8080")

	r.Run(":8080")
}