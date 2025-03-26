package main

import (
	"Framework/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/Authors", handlers.GetAuthors)
	r.POST("/Authors", handlers.CreateAuthor)

	r.GET("/Categories", handlers.GetCategories)
	r.POST("/Categories", handlers.CreateCategory)

	r.GET("/Books", handlers.GetBooks)
	r.POST("/Books", handlers.CreateBook)
	r.GET("/Books/:id", handlers.GetBookByID)
	r.PUT("/Books/:id", handlers.UpdateBook)
	r.DELETE("/Books/:id", handlers.DeleteBook)

	r.Run(":8080")
}
