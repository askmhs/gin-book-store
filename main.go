package main

import (
	"net/http"

	"github.com/askmhs/gin-book-store/config"
	"github.com/askmhs/gin-book-store/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	r := gin.Default()

	config.ConnectDatabase()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "Hello, world!",
		})
	})

	// Register routes
	routes.RegisterRoutes(r)

	err := r.Run(":" + config.AppConfig.AppPort)

	if err != nil {
		return
	}
}
