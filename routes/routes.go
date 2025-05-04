package routes

import (
	"github.com/askmhs/gin-book-store/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	books := r.Group("books")
	{
		books.GET("/", controllers.FindBooks)
		books.GET("/:id", controllers.FindBook)
		books.POST("/", controllers.CreateBook)
		books.PATCH("/:id", controllers.UpdateBook)
		books.DELETE("/:id", controllers.DeleteBook)
	}

}
