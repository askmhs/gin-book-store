package routes

import (
	"github.com/askmhs/gin-book-store/config"
	"github.com/askmhs/gin-book-store/controllers"
	"github.com/askmhs/gin-book-store/repositories"
	"github.com/askmhs/gin-book-store/services"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	bookRepo := repositories.NewBookRepository(config.DB)
	bookController := controllers.NewBookController(services.NewBookService(bookRepo))

	userRepo := repositories.NewUserRepository(config.DB)
	userController := controllers.NewUserController(services.NewUserService(userRepo, services.NewJwtService()))

	books := r.Group("books")
	{
		books.GET("/", bookController.FindBooks)
		books.GET("/:id", bookController.FindBook)
		books.POST("/", bookController.CreateBook)
		books.PATCH("/:id", bookController.UpdateBook)
		books.DELETE("/:id", bookController.DeleteBook)
	}

	users := r.Group("users")
	{
		users.POST("/register", userController.RegisterUser)
		users.POST("/login", userController.LoginUser)
	}

}
