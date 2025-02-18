package controllers

import (
	"net/http"
	"warehouse-api/config"
	"warehouse-api/models"

	"github.com/gin-gonic/gin"
)

type ShelfRequest struct {
	ShelfParent string `json:"shelfparent" binding:"required"`
	ShelfChild  string `json:"shelfchild" binding:"required"`
}

func CreateShelf(c *gin.Context) {
	var shelf models.Shelf
	if err := c.ShouldBindJSON(&shelf); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":        "Invalid request format",
			"correct_json": models.Shelf{ShelfParent: "A", ShelfChild: "A2"},
		})
		return
	}

	var existShelf models.Shelf
	if err := config.DB.Where("shelfchild = ?", shelf.ShelfChild).First(&existShelf).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Shelf child already exists",
		})
	}

	if err := config.DB.Create(&shelf).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create shelf"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Shelf created successfully",
		"shelf":   shelf,
	})
}

func GetShelves(c *gin.Context) {
	var shelves []models.Shelf
	config.DB.Find(&shelves)
	c.JSON(http.StatusOK, shelves)
}
