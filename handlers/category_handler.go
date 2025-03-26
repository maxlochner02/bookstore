package handlers

import (
	"Framework/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

var categories = []models.Category{
	{ID: 1, Name: "tragedy"},
	{ID: 2, Name: "drama"},
}

func GetCategories(c *gin.Context) {
	c.JSON(http.StatusOK, categories)
}

func CreateCategory(c *gin.Context) {
	var newCategory models.Category
	if err := c.ShouldBindJSON(&newCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newCategory.ID = len(books) + 1
	categories = append(categories, newCategory)
	c.JSON(http.StatusCreated, newCategory)
}
