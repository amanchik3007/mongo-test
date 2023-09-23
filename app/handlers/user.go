package handlers

import (
	"context"

	"mongo-test/app/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserHandler contains handlers related to user operations.
type UserHandler struct {
	UserCollection *mongo.Collection
}

// NewUserHandler creates a new UserHandler.
func NewUserHandler(client *mongo.Client, databaseName string, collectionName string) *UserHandler {
	collection := client.Database(databaseName).Collection(collectionName)
	return &UserHandler{
		UserCollection: collection,
	}
}

// InsertUser inserts a new user into the MongoDB collection.
func (uh *UserHandler) InsertUser(c *fiber.Ctx) error {
	resp := models.BaseResponse{
		Success: true,
		Message: "Record successfully inserted",
	}
	// Parse JSON request body into a User struct
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		resp.Success = false
		resp.MsgType = "E_MONGO_INSERT"
		resp.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(resp)
	}

	// Insert the user into the MongoDB collection
	_, err := uh.UserCollection.InsertOne(context.Background(), user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to insert user",
		})
	}

	// Return a success response
	return c.Status(fiber.StatusCreated).JSON(resp)
}
