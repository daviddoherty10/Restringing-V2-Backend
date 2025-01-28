package controllers

import (
	"Restringing-V2/entity"
	"Restringing-V2/internal/database"
	"Restringing-V2/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserRequestBody struct {
	Firstname string `json:"firstname" binding:"required"`
	Surname   string `json:"surname" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

func Login(c *gin.Context, db database.Service) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user, err := db.GetUserByUsername(credentials.Username)
	log.Println(credentials.Password)
	log.Println(credentials.Username)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Check credentials
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Line 44" + err.Error()})
		return
	}

	/*if user.Password != credentials.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Hello"})
		log.Println(user.Password + "  " + credentials.Password)
		return
	}*/

	// Generate token
	token, err := utils.GenerateToken(uint(user.ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func CreateAccount(c *gin.Context, db database.Service) {
	var requestBody UserRequestBody

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Line 23: " + err.Error()})
		return
	}

	var user entity.User
	user.FirstName = requestBody.Firstname
	user.Surname = requestBody.Surname
	user.Username = requestBody.Username
	user.Email = requestBody.Email
	var err error
	user.Password, err = utils.HashPassword(requestBody.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Problem with password"})
	}

	if err := db.CreateUser(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Line 25: " + err.Error()})
		return
	}

	c.Status(http.StatusOK)

}
