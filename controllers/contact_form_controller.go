package controllers

import (
	"Restringing-V2/entity"
	"Restringing-V2/internal/database"

	"github.com/gin-gonic/gin"

	"log"
	"net/http"
)

type ContactFormBody struct {
	Email   string `json:"email" binding:"required"`
	Message string `json:"message" binding:"required"`
}

func CreateContactFormResponse(c *gin.Context, db database.Service) {

	var requestBody ContactFormBody
	// Parse JSON request
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		log.Println("Failed to Bind Json on Contact Form Createation: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}
	var form entity.ContactForm

	form.Email = requestBody.Email
	form.Message = requestBody.Email

	if err := db.CreateContactFormResponse(form); err != nil {
		log.Println("Failed to save contact form data: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Response Successfully saved",
	})
}
