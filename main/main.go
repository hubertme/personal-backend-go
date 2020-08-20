package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"wang.hihubert.personal-backend/database"
	"wang.hihubert.personal-backend/gin"
	"wang.hihubert.personal-backend/services"
)

func main() {
	fmt.Println("Hello, world!")

	// Dotenv setup
	dotenvErr := godotenv.Load(".env")
	if dotenvErr != nil {
		fmt.Println(dotenvErr.Error())
		log.Fatal("PANIC!!!! Failed to load dotenv!!!!!")
	}

	// Setup Services
	services.MailgunSetup()
	database.InitMySQL()

	// Gin server
	err := gin.SetupGinServer()
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal("PANIC!!!! Failed to setup Gin server!!!!!")
	}
}
