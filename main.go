package main

import (
	"log"
	"warehouse-api/config"
	"warehouse-api/models"
	"warehouse-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	// Миграции
	config.DB.AutoMigrate(&models.Shelf{}, &models.Product{})

	r := gin.Default()

	// Подключаем роуты
	routes.ShelvesRoutes(r)
	routes.ProductRoutes(r)

	log.Println("Сервер запущен на порту 8080")
	r.Run(":8080")
}
