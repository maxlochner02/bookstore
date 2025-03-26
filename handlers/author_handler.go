package handlers

import (
	"Framework/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

var authors = []models.Author{
	{ID: 1, Name: "Goethe"},
	{ID: 2, Name: "Schiller"},
}

func GetAuthors(c *gin.Context) {
	c.JSON(http.StatusOK, authors)
}

func CreateAuthor(c *gin.Context) {
	var newAuthor models.Author
	if err := c.ShouldBindJSON(&newAuthor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newAuthor.ID = len(authors) + 1
	authors = append(authors, newAuthor)
	c.JSON(http.StatusCreated, newAuthor)
}
