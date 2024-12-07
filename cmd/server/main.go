package main

import (
	"product-management-system/config"
	"product-management-system/controllers"
	"product-management-system/routes"
	"product-management-system/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Database
	config.ConnectDB()

	// Initialize Logger
	utils.InitLogger()
	utils.Logger.Info("Logger initialized")

	// Set up Gin router
	r := gin.Default()

	// Register routes with explicit controller functions
	routes.RegisterRoutes(
		r,
		controllers.CreateProduct,
		controllers.GetProductByID,
		controllers.GetProducts,
	)

	utils.Logger.Info("Server starting on port 8080")
	r.Run(":8080")
}
