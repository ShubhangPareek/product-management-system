package routes

import (
	"product-management-system/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/products", controllers.CreateProduct)
	r.GET("/products/:id", controllers.GetProductByID)
	r.GET("/products", controllers.GetProducts)
}