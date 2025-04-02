package handlers

import (
	"Framework/config"
	"Framework/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAuthors(c *gin.Context) {
	var authors []models.Author
	// Zugriff auf die globale DB-Instanz aus config und Abfrage ausführen
	if err := config.DB.Find(&authors).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve authors"})
		return
	}
	c.JSON(http.StatusOK, authors)
}

func CreateAuthor(c *gin.Context) {
	var newAuthor models.Author
	// JSON-Daten in das struct newAuthor binden
	if err := c.ShouldBindJSON(&newAuthor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Neues Autor in die Datenbank einfügen
	if err := config.DB.Create(&newAuthor).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create author"})
		return
	}

	c.JSON(http.StatusCreated, newAuthor)
}
