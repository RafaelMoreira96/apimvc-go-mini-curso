package controllers

import (
	"catalogolivros/database"
	"catalogolivros/models"
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
