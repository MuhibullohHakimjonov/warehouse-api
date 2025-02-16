package routes

import (
	"warehouse-api/controllers"

	"github.com/gin-gonic/gin"
)

func ShelvesRoutes(r *gin.Engine) {
	shelves := r.Group("/shelves")
	{
		shelves.POST("/", controllers.CreateShelf)
		shelves.GET("/", controllers.GetShelves)
	}
}

