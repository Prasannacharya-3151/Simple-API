package main

import (
	"log"
	"os"
	"simple-api/config"
	"simple-api/models"
	"simple-api/routes"
    "github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
)

func main() {
	//load .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system env vars")
	}

	//connect db
	config.ConnectDB()

	//auto migrate
	config.DB.AutoMigrate(&models.User{})

	r := gin.Default() //r which is a Gin server

	routes.SetupRoutes(r)

	port := os.Getenv("PORT")
	if port == ""{
		port = "8080"
	}

	log.Printf("Server running on port %s", port)
	r.Run(":" +port)
}