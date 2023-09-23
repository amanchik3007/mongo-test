package main

import (
	"log"
	"os"

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

	// Create a Fiber instance
	app := fiber.New()

	// Initialize MongoDB connection
	mongoClient, err := database.ConnectMongoDB(os.Getenv("MONGODB_URI"))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Initialize UserHandler with MongoDB collection
	userHandler := handlers.NewUserHandler(mongoClient, "daily", "user")

	// Initialize application routes and pass the UserHandler
	routes.SetupAPIRoutes(app, userHandler)

	// Start the Fiber server
	err = app.Listen(":" + os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
