package main

import (
	"eCommerce/config"
	"eCommerce/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("env/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// Instance Server
	server := gin.Default()

	// Database
	database := config.Database{}
	db := database.InitDatabase()

	// Routes
	router := routes.Route{Server: server, Db: db}
	router.Register()

	// Startup Server
	server.Run()
}
