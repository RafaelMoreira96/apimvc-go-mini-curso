package controllers

import (
	"catalogolivros/database"
	"catalogolivros/models"
	"encoding/csv"
	"io"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	db := database.GetDatabase()
	var book models.Book

	error := c.ShouldBindJSON(&book)
	if error != nil {
		c.JSON(400, gin.H{
			"error": error.Error(),
		})
		return
	}

	if validationError := book.Validate(); validationError != nil {
		c.JSON(400, gin.H{
			"error": validationError.Error(),
		})
		return
	}

	error = db.Create(&book).Error
	if error != nil {
		c.JSON(400, gin.H{
			"error": error.Error(),
		})
		return
	}

	c.JSON(200, book)
}

func ReadBook(c *gin.Context) {
	id := c.Param("id")

	newid, error := strconv.Atoi(id)
	if error != nil {
		c.JSON(400, gin.H{
			"error": error.Error(),
		})
		return
	}

	db := database.GetDatabase()
	var book models.Book

	error = db.First(&book, newid).Error
	if error != nil {
		c.JSON(400, gin.H{
			"error": error.Error(),
		})
		return
	}

	c.JSON(200, book)
}

func ReadAllBooks(c *gin.Context) {
	db := database.GetDatabase()
	var books []models.Book

	error := db.Find(&books).Error
	if error != nil {
		c.JSON(400, gin.H{
			"error": error.Error(),
		})
		return
	}

	c.JSON(200, books)
}

func UpdateBook(c *gin.Context) {
	db := database.GetDatabase()
	var book models.Book

	error := c.ShouldBindJSON(&book)
	if error != nil {
		c.JSON(400, gin.H{
			"error": error.Error(),
		})
		return
	}

	if validationError := book.Validate(); validationError != nil {
		c.JSON(400, gin.H{
			"error": validationError.Error(),
		})
		return
	}

	error = db.Save(&book).Error
	if error != nil {
		c.JSON(400, gin.H{
			"error": error.Error(),
		})
		return
	}
	c.JSON(200, book)
}

func DeleteBook(c *gin.Context) {
	id := c.Param(("id"))

	newid, error := strconv.Atoi(id)
	if error != nil {
		c.JSON(400, gin.H{
			"error": error.Error(),
		})
		return
	}

	db := database.GetDatabase()

	error = db.First(&models.Book{}, newid).Error
	if error != nil {
		c.JSON(400, gin.H{
			"error": error.Error(),
		})
		return
	}

	error = db.Delete(&models.Book{}, newid).Error
	if error != nil {
		c.JSON(400, gin.H{
			"error": error.Error(),
		})
		return
	}
	c.Status(204)
}

func CreateBooksByCSV(c *gin.Context) {
	file, err := os.Open("files/books.csv")
	if err != nil {
		c.JSON(500, gin.H{"error": "Erro ao abrir o arquivo CSV"})
		return
	}
	defer file.Close()

	db := database.GetDatabase()
	tx := db.Begin()

	reader := csv.NewReader(file)
	createdCount := 0
	errorCount := 0

	for {
		row, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			c.JSON(500, gin.H{"error": "Erro ao ler o arquivo CSV"})
			tx.Rollback()
			return
		}

		book := models.Book{
			Title:  row[1],
			Author: row[2],
			Genre:  row[3],
		}

		if err := tx.Create(&book).Error; err != nil {
			errorCount++
			continue
		}

		createdCount++
	}

	if errorCount > 0 {
		tx.Rollback()
		c.JSON(500, gin.H{"error": "Erro ao salvar alguns livros no banco de dados"})
		return
	}

	tx.Commit()

	c.JSON(200, gin.H{"message": "Importação concluída", "created": createdCount})
}
