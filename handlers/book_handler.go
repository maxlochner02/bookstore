package handlers

import (
	"Framework/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var books = []models.Book{
	{ID: 1, Title: "Faust", AuthorID: 1, CategoryID: 1, Price: 5000},
	{ID: 2, Title: "Die Räuber", AuthorID: 2, CategoryID: 2, Price: 6000},
}

func GetBooks(c *gin.Context) {
	category := c.Query("category")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize := 5
	filteredBooks := []models.Book{}

	for _, book := range books {
		if category == "" || strconv.Itoa(book.CategoryID) == category {
			filteredBooks = append(filteredBooks, book)
		}
	}

	start := (page - 1) * pageSize
	end := start + pageSize
	if start > len(filteredBooks) {
		c.JSON(http.StatusOK, []models.Book{})
		return
	}
	if end > len(filteredBooks) {
		end = len(filteredBooks)
	}

	c.JSON(http.StatusOK, filteredBooks[start:end])
}

func GetBookByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, book := range books {
		if book.ID == id {
			c.JSON(http.StatusOK, book)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{})
}

func CreateBook(c *gin.Context) {
	var newBook models.Book
	c.ShouldBindJSON(&newBook)
	if newBook.Title == "" || newBook.Price < 1000 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title is required and price must be at least 1000"})
		return
	}
	newBook.ID = len(books) + 1
	books = append(books, newBook)
	c.JSON(http.StatusCreated, newBook)
}

func UpdateBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updatedBook models.Book
	c.ShouldBindJSON(&updatedBook)
	if updatedBook.Title == "" || updatedBook.Price < 1000 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title is required and price must be at least 1000"})
		return
	}
	for i, book := range books {
		if book.ID == id {
			updatedBook.ID = id
			books[i] = updatedBook
			c.JSON(http.StatusOK, updatedBook)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{})
}

func DeleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			c.JSON(http.StatusOK, gin.H{})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{})
}
