package routes

import (
	"warehouse-api/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine) {
	products := r.Group("/products")
	{
		products.POST("/", controllers.CreateProduct)
		products.GET("/:id", controllers.GetProductLocation)
	}
}
