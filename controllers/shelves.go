package controllers

import (
	"net/http"
	"warehouse-api/config"
	"warehouse-api/models"

	"github.com/gin-gonic/gin"
)

func CreateShelf(c *gin.Context) {
	var shelf models.Shelf

	if err := c.ShouldBindJSON(&shelf); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format", "details": err.Error()})
		return
	}

	config.DB.Create(&shelf)
	c.JSON(http.StatusCreated, shelf)
}

func GetShelves(c *gin.Context) {
	var shelves []models.Shelf
	config.DB.Find(&shelves)
	c.JSON(http.StatusOK, shelves)
}
