package controllers

import (
	"net/http"

	"github.com/askmhs/gin-book-store/models"
	"github.com/askmhs/gin-book-store/services"
	"github.com/gin-gonic/gin"
)

type RegisterUserInput struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	UserName  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type LoginUserInput struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserController struct {
	userService *services.UserService
}

func NewUserController(s *services.UserService) *UserController {
	return &UserController{userService: s}
}

func (uc *UserController) RegisterUser(c *gin.Context) {
	var input RegisterUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := models.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		UserName:  input.UserName,
		Password:  input.Password,
	}

	err := uc.userService.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Registration is successful",
	})
}

func (uc *UserController) LoginUser(c *gin.Context) {
	var input LoginUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := uc.userService.LoginUser(input.UserName, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
