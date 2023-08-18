package models

import "github.com/go-playground/validator/v10"

type Book struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
	Genre  string `json:"genre" binding:"required"`
}

func (b *Book) Validate() error {
	validate := validator.New()
	return validate.Struct(b)
}
