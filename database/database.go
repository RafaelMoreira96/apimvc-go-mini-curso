package database

import (
	"catalogolivros/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() *gorm.DB {
	database, err := gorm.Open(sqlite.Open("files/database.sqlite"), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro: ", err)
		return nil
	}
	db = database

	db.AutoMigrate(models.Book{})

	return db.Begin()
}

func GetDatabase() *gorm.DB {
	return db
}
