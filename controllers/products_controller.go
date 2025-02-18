package controllers

import (
	"net/http"
	"warehouse-api/config"
	"warehouse-api/models"

	"github.com/gin-gonic/gin"
)

type ProductRequest struct {
	Name        string `json:"name" binding:"required"`
	Type        string `json:"type" binding:"required"`
	ShelfParent string `json:"shelfparent" binding:"required"`
	ShelfChild  string `json:"shelfchild" binding:"required"`
}

func CreateProduct(c *gin.Context) {
	var req ProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request format",
			"correct_json": ProductRequest{
				Name:        "Example Product",
				Type:        "Example Type",
				ShelfParent: "A",
				ShelfChild:  "A1",
			},
		})
		return
	}

	var shelf models.Shelf
	if err := config.DB.Where("shelf_parent = ? AND shelf_child = ?", req.ShelfParent, req.ShelfChild).First(&shelf).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Shelf not found",
			"message": "The provided shelf and shelf path do not exist. Please create them first.",
		})
		return
	}

	product := models.Product{Name: req.Name, Type: req.Type, ShelfID: shelf.ID}
	config.DB.Create(&product)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Product created successfully",
		"product": product,
	})
}

func GetProductLocation(c *gin.Context) {
	var product models.Product
	productID := c.Param("id")

	if err := config.DB.Preload("Shelf").First(&product, productID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Location": product.Shelf.ShelfChild,
	})
}
