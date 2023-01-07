package models

type Book struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Author string `json:"author"`
	Title  string `json:"title"`
}
