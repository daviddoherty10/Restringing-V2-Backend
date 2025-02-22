package utils

import (
	//"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte("your-secret-key") // Replace with a secure key

// GenerateToken creates a new JWT
func GenerateToken(userID uint, exp uint64) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     exp, // Expires in 24 hours
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ValidateToken checks if the JWT is valid
func ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
}
