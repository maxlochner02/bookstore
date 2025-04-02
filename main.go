// Datei: main.go
package main

import (
	"Framework/config"
	"Framework/handlers"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	dsn := "user=postgres dbname=postgres sslmode=disable password=yourpassword"
	config.InitDB(dsn)
	config.Migrate()

	r := gin.Default()

	// Routen für Bücher
	r.GET("/books", handlers.GetBooks)
	r.GET("/books/:id", handlers.GetBookByID)
	r.POST("/books", handlers.CreateBook)
	r.PUT("/books/:id", handlers.UpdateBook)
	r.DELETE("/books/:id", handlers.DeleteBook)

	// Routen für Autoren
	r.GET("/authors", handlers.GetAuthors)
	r.POST("/authors", handlers.CreateAuthor)

	// Routen für Kategorien
	r.GET("/categories", handlers.GetCategories)
	r.POST("/categories", handlers.CreateCategory)

	// Server starten
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server starten fehlgeschlagen:", err)
	}
}
