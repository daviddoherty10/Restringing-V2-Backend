package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"Restringing-V2/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Try to get token from cookies
		tokenString, err := c.Cookie("auth_token")
		if err != nil {
			// If no cookie, check the Authorization header as a fallback
			authHeader := c.GetHeader("Authorization")
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - missing token"})
				c.Abort()
				return
			}
			tokenString = strings.TrimPrefix(authHeader, "Bearer ")
		}

		// Debugging: Print the token string to check if it's retrieved correctly
		fmt.Println("Token received:", tokenString)

		// Validate the token
		token, err := utils.ValidateToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Next()
	}
}
