package controllers

import (
	"net/http"

	"github.com/askmhs/gin-book-store/config"
	"github.com/askmhs/gin-book-store/models"
	"github.com/askmhs/gin-book-store/services"

	"github.com/gin-gonic/gin"
)

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

type BookController struct {
	bookService *services.BookService
}

func NewBookController(bookService *services.BookService) *BookController {
	return &BookController{bookService: bookService}
}

func (bc *BookController) FindBooks(c *gin.Context) {
	filters := map[string]any{
		"title": c.Query("title"),
	}

	books, err := bc.bookService.GetBooks(filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}

func (bc *BookController) FindBook(c *gin.Context) {
	var book models.Book

	if err := config.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func (bc *BookController) CreateBook(c *gin.Context) {
	var input CreateBookInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	book := models.Book{Title: input.Title, Author: input.Author}
	config.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

func (bc *BookController) UpdateBook(c *gin.Context) {
	var book models.Book

	if err := config.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found!",
		})
		return
	}

	var input UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&book).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func (bc *BookController) DeleteBook(c *gin.Context) {
	var book models.Book

	if err := config.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found!",
		})
		return
	}

	config.DB.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"success": true})
}
