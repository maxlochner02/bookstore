// Datei: handlers/book_handler.go
package handlers

import (
	"Framework/config"
	"Framework/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetBooks(c *gin.Context) {
	category := c.Query("category")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize := 5
	var filteredBooks []models.Book

	query := config.DB.Limit(pageSize).Offset((page - 1) * pageSize)
	if category != "" {
		query = query.Where("category_id = ?", category)
	}

	if err := query.Find(&filteredBooks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve books"})
		return
	}

	c.JSON(http.StatusOK, filteredBooks)
}

func GetBookByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var book models.Book
	if err := config.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

func CreateBook(c *gin.Context) {
	var newBook models.Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	if newBook.Title == "" || newBook.Price < 1000 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title is required and price must be at least 1000"})
		return
	}

	if err := config.DB.Create(&newBook).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		return
	}

	c.JSON(http.StatusCreated, newBook)
}

func UpdateBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updatedBook models.Book
	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	if updatedBook.Title == "" || updatedBook.Price < 1000 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title is required and price must be at least 1000"})
		return
	}

	var book models.Book
	if err := config.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	book.Title = updatedBook.Title
	book.AuthorID = updatedBook.AuthorID
	book.CategoryID = updatedBook.CategoryID
	book.Price = updatedBook.Price

	if err := config.DB.Save(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book"})
		return
	}

	c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var book models.Book
	if err := config.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	if err := config.DB.Delete(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}
