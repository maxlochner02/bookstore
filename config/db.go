package config

import (
	"Framework/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

var DB *gorm.DB

func InitDB(dsn string) {
	var err error
	DB, err = gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}
}

func Migrate() {
	DB.AutoMigrate(&models.Book{}, &models.Author{}, &models.Category{})
}
