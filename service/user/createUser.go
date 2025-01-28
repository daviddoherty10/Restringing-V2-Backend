package user

import (
	"Restringing-V2/entity"
	"Restringing-V2/internal/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRequestBody struct {
	Firstname string `json:"firstname" binding:"required"`
	Surname   string `json:"surname" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

func HandleUserCreation(db database.Service, c *gin.Context) {
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
	user.Password = requestBody.Password

	if err := db.CreateUser(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Line 25: " + err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
