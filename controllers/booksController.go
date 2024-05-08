package controllers

import (
	"github.com/abdurahmonov-azizbek/go-crud/initializers"
	"github.com/abdurahmonov-azizbek/go-crud/models"
	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	var body struct {
		Title       string
		Description string
		Author      string
	}

	c.Bind(&body)

	book := models.Book{Title: body.Title, Description: body.Description, Author: body.Author}
	result := initializers.DB.Create(&book)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"book": book,
	})
}

func GetBooks(c *gin.Context) {
	var books []models.Book
	initializers.DB.Find(&books)

	c.JSON(200, gin.H{
		"books": books,
	})
}

func GetBookById(c *gin.Context) {
	id := c.Param("id")
	var book models.Book

	initializers.DB.First(&book, id)
	c.JSON(200, gin.H{
		"book": book,
	})
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Title       string
		Description string
		Author      string
	}

	c.Bind(&body)

	var foundedBook models.Book
	initializers.DB.Find(&foundedBook, id)

	var result = initializers.DB.Model(&foundedBook).Updates(models.Book{
		Title:       body.Title,
		Description: body.Description,
		Author:      body.Author,
	})

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"book": foundedBook,
	})
}

func DeleteBookById(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Book{}, id)

	c.Status(200)
}
