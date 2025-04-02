package models

type Author struct {
	ID   int    `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}
