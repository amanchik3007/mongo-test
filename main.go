package main

import (
	"log"

	"mongo-test/app/config"
	"mongo-test/app/database"
	"mongo-test/app/handlers"
	"mongo-test/app/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	// Load configuration
	cfg := config.NewConfig()

	// Create a Fiber instance
	app := fiber.New()

	// Initialize MongoDB connection
	mongoClient, err := database.ConnectMongoDB(cfg.MongoDBURI)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Initialize UserHandler with MongoDB collection
	userHandler := handlers.NewUserHandler(mongoClient, cfg.DbName, cfg.MongoUserCollection)

	// Initialize application routes and pass the UserHandler
	routes.SetupAPIRoutes(app, userHandler)

	// Start the Fiber server
	err = app.Listen(":" + cfg.Port)
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
