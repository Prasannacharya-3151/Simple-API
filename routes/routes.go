package routes

import (
	"simple-api/handlers"

	"github.com/labstack/echo/v4/middleware"
)

func main() {
	api := r.Group("/api")
	{
		api.POST("/signup", handlers.Signup)
		api.POST("/signin", handlers.Login)

		//protected routes grp
		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.GET("/profile", handlers.Profile)
		}
	}
}