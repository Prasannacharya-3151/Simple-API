package routes

import (
	"simple-api/handlers"
	"simple-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/signup", handlers.Signup)
		api.POST("/login", handlers.Login)

		//protected router group
		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.GET("/profile", handlers.Profile)
		}
	}
}