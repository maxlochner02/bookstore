package models

type Book struct {
	ID         int    `json:"id" gorm:"primary_key"`
	Title      string `json:"title"`
	AuthorID   int    `json:"author_id"`
	CategoryID int    `json:"category_id"`
	Price      int    `json:"price"`
}
