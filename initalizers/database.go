package initalizers

import (
	"fmt"
	"log"
	"os"

	"go-crud-v1/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error

	errorENV := godotenv.Load()
	if errorENV != nil {
		log.Fatal("Failed to load env file")
	}

	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, dbUser, dbPass, dbName, dbPort)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	err = DB.AutoMigrate(&models.Post{})
	if err != nil {
		log.Fatal("Failed to migrate database")
	}
}
