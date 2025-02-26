package server

import (
	"Restringing-V2/controllers"
	"Restringing-V2/middlewares"
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
		AllowCredentials: true,
	}))
	r.Use(middlewares.LoggingMiddleware(s.db))

	v1Router := r.Group("/api/v1")
	{
		auth := v1Router.Group("/auth")
		{
			auth.POST("/create-user", func(ctx *gin.Context) {
				controllers.CreateAccount(ctx, s.db)
			})
			auth.POST("/login", func(ctx *gin.Context) {
				controllers.Login(ctx, s.db)
			})
		}

		protected := v1Router.Group("/protected")
		protected.Use(middlewares.AuthMiddleware())
		{
			userRouter := protected.Group("/user")
			{
				userRouter.GET("/data", func(ctx *gin.Context) {
					controllers.GetUserData(ctx, s.db)
				})
			}
			protected.DELETE("/delete-user", func(ctx *gin.Context) {
				controllers.RequestAccountDeletion(ctx, s.db)
			})
			protected.POST("/create-potential-stringer", func(ctx *gin.Context) {
				controllers.CreateApplication(ctx, s.db)
			})
			protected.POST("/logout", func(ctx *gin.Context) {
				controllers.Logout(ctx)
			})

		}
		contactForm := v1Router.Group("/contact-form")
		{
			contactForm.POST("/create", func(ctx *gin.Context) {
				controllers.CreateContactFormResponse(ctx, s.db)
			})
		}
	}

	return r
}
