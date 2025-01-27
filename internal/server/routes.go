package server

import (
	"Restringing-V2/entity"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type UserRequestBody struct {
	Firstname string `json:"firstname" binding:"required"`
	Surname   string `json:"surname" binding:"required"`
	Email     string `json:"email" binding:"required"`
}

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Add your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true, // Enable cookies/auth
	}))

	r.GET("/", s.HelloWorldHandler)

	r.GET("/health", s.healthHandler)

	r.POST("/create-user", s.handleUserCreation)

	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}

func (s *Server) handleUserCreation(c *gin.Context) {
	var requestBody UserRequestBody

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user entity.User
	user.FirstName = requestBody.Firstname
	user.Surname = requestBody.Surname
	user.Email = requestBody.Email

	if err := s.db.CreateUser(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
