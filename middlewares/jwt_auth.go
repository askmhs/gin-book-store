package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/askmhs/gin-book-store/services"
	"github.com/gin-gonic/gin"
)

func JwtAuth(jwt *services.JwtService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header missing or malformed",
			})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		fmt.Println(tokenString)

		claims, err := jwt.ValidateToken(tokenString)
		if err != nil {
			fmt.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Token is invalid or expired",
			})
			return
		}

		c.Set("user", claims)
		c.Next()
	}
}
