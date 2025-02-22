package controllers

import (
	"Restringing-V2/entity"
	"Restringing-V2/internal/database"
	"Restringing-V2/utils"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserRequestBody struct {
	Firstname        string `json:"firstname" binding:"required"`
	Surname          string `json:"surname" binding:"required"`
	Email            string `json:"email" binding:"required"`
	Username         string `json:"username" binding:"required"`
	Password         string `json:"password" binding:"required"`
	HasAcceptedTerms bool   `json:"has_accepted_terms" binding:"required"`
}

type IdRequestBody struct {
	Id uint `json:"id" binding:"required"`
}

func Login(c *gin.Context, db database.Service) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		log.Println("Failed to Bind Json on Login: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user, err := db.GetUserByEmail(credentials.Email)
	if err != nil {
		log.Println("Failed to Get User Data on Login: " + err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		log.Println("Failed to Compare Hash and Password on Login: " + err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Line 44" + err.Error()})
		return
	}

	token, err := utils.GenerateToken(uint(user.ID), uint64(time.Now().Add(24*time.Hour).Unix()))
	if err != nil {
		log.Println("Failed to Generate a token on Login: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.SetCookie("auth_token", token, 3600, "/", "", false, true) // 1-hour expiry
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Login successful",
		"token":   token,
	})

}

func CreateAccount(c *gin.Context, db database.Service) {
	var requestBody UserRequestBody

	// Parse JSON request
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		log.Println("Failed to Bind Json on account Creation: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}

	// Create User struct
	var user entity.User
	user.FirstName = requestBody.Firstname
	user.Surname = requestBody.Surname
	user.Username = requestBody.Username
	user.Email = requestBody.Email
	user.HasAcceptedTerms = requestBody.HasAcceptedTerms
	user.EmailVerification = false

	// Hash password
	var err error
	user.Password, err = utils.HashPassword(requestBody.Password)
	if err != nil {
		log.Println("Failed to Hash Password on account Creation: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Problem with password encryption"})
		return
	}

	// Save user to DB
	if err := db.CreateUser(user); err != nil {
		log.Println("Failed to save user data on account Creation: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error: " + err.Error()})
		return
	}

	// ✅ Always return JSON on success
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "User created successfully",
	})
}

func RequestAccountDeletion(c *gin.Context, db database.Service) {
	token := c.GetHeader("Authorization")
	isValid, err := utils.ValidateToken(token)
	if err != nil || isValid.Valid == false {
		c.Status(http.StatusUnauthorized)
		c.Abort()
		return
	}

	var idFromJson IdRequestBody

	if err := c.ShouldBindJSON(&idFromJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not bind Json"})
		return
	}

	if err := db.DeleteUser(idFromJson.Id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to delete user from Database"})
	}

}

func Logout(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	log.Println(userID)

	token, err := utils.GenerateToken(userID.(uint), uint64(1*time.Second))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to expire token"})
		log.Println("Failed to expire token")
		return
	}

	c.SetCookie("auth_token", token, -1, "/", "", false, false)

	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}
