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
			userRouter.GET("/get-user-by-id/:id", func(ctx *gin.Context) {
				user.HandleGetUserById(s.db, ctx)
			})
		}

		orderRouter := v1Router.Group("/order") // ✅ Added missing '/'
		{
			orderRouter.GET("/get-order", func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, gin.H{"order": "No orders"})
			})
		}

		// ✅ Corrected: auth is now inside v1Router
		auth := v1Router.Group("/auth")
		{
			auth.POST("/create-user", func(ctx *gin.Context) { // ✅ Moved from userRouter
				controllers.CreateAccount(ctx, s.db)
			})
			auth.POST("/login", func(ctx *gin.Context) {
				controllers.Login(ctx, s.db)
			})
			auth.DELETE("/delete", func(ctx *gin.Context) {
				controllers.RequestAccountDeletion(ctx, s.db)
			})
		}

		protected := v1Router.Group("/protected")
		protected.Use(middlewares.AuthMiddleware())
		{
			protected.GET("/data", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"message": "You have access to protected data!"})
			})
		}
	}

	return r
}
