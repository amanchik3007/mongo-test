package routes

import (
	"mongo-test/app/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupAPIRoutes(app *fiber.App, userHandler *handlers.UserHandler) {
	api := app.Group("/api")

	// POST /api/users - Create a new user
	api.Post("/addUsers", userHandler.InsertUser)

	// Add other API routes as needed
}
