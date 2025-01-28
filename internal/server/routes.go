package server

import (
	"Restringing-V2/controllers"
	"Restringing-V2/middlewares"
	"Restringing-V2/service/user"
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

	v1Router := r.Group("/api/v1")
	{
		userRouter := v1Router.Group("/user")
		{
			userRouter.POST("/create-user", func(ctx *gin.Context) {
				user.HandleUserCreation(s.db, ctx)
			})

			userRouter.GET("/get-user-by-id/:id", func(ctx *gin.Context) {
				user.HandleGetUserById(s.db, ctx)
			})
		}
		orderRouter := v1Router.Group("order")
		{
			orderRouter.GET("/get-order", func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, gin.H{"order": "No orders"})
			})
		}
		auth := r.Group("/auth")
		{
			auth.POST("/login", func(ctx *gin.Context) {
				controllers.Login(ctx, s.db)
			})
		}
		protected := r.Group("/protected")
		protected.Use(middlewares.AuthMiddleware())
		{
			protected.GET("/data", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "You have access to protected data!"})
			})
		}
	}

	return r
}
