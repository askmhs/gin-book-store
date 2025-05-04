package main

import (
	"net/http"

	"github.com/askmhs/gin-book-store/models"
	"github.com/askmhs/gin-book-store/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "Hello, world!",
		})
	})

	// Register routes
	routes.RegisterRoutes(r)

	err := r.Run()

	if err != nil {
		return
	}
}
