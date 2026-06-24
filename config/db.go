package config
//Read database credentials from .env, connect to PostgreSQL using GORM, and make the connection available throughout the whole project.

import (
	"fmt"
	"log"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm" //orm is the library and this is the postgreSql driver GORM
)

var DB *gorm.DB //which will creates a DB and will store in the database connection we can use like a config.DB.Create(),config.DB.Find(), config.DB.Where()

func ConnectDB() { //this function connectes the db
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	log.Print("Host:",host,"port",port)

	dsn := fmt.Sprintf( //dsn stands "data source name" means everything combined database address, host, port, username, password, database name
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}) //this measn use postgrsql driver and connect using this dsn
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connected successfully")
	DB = db
}