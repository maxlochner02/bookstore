// Datei: handlers/categories.go
package handlers

import (
	"Framework/config"
	"Framework/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetCategories gibt alle Kategorien zurück
func GetCategories(c *gin.Context) {
	var categories []models.Category
	// Abfrage der Kategorien aus der DB
	if err := config.DB.Find(&categories).Error; err != nil {
		// Detailierte Fehlermeldung für den Fehlerfall
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve categories: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, categories)
}

// CreateCategory erstellt eine neue Kategorie
func CreateCategory(c *gin.Context) {
	var newCategory models.Category
	// JSON-Daten binden
	if err := c.ShouldBindJSON(&newCategory); err != nil {
		// Rückgabe einer detaillierten Fehlermeldung im Fehlerfall
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format: " + err.Error()})
		return
	}

	// Validierung: Der Name der Kategorie muss angegeben werden
	if newCategory.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category name is required"})
		return
	}

	// Speichern der neuen Kategorie in der DB
	if err := config.DB.Create(&newCategory).Error; err != nil {
		// Detailierte Fehlermeldung im Fehlerfall
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newCategory)
}
