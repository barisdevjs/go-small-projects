package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RequestBody struct {
	IDs []string `json:"ids"`
}

func DeleteEmployees(c *fiber.Ctx) error {
	var reqBody RequestBody
	if err := c.BodyParser(&reqBody); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Failed to parse request body"})
	}

	// Get the IDs from the request body
	idArray := reqBody.IDs

	// Create a slice to hold the ObjectIDs
	var objectIDs []primitive.ObjectID

	// Convert each string ID to primitive.ObjectID and add it to the slice
	for _, id := range idArray {
		employeeID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid ID format"})
		}
		objectIDs = append(objectIDs, employeeID)
	}

	// Get the MongoDB collection
	collection := mg.Db.Collection("employees")

	// Create the filter with $in operator to match multiple employee IDs
	filter := bson.M{"_id": bson.M{"$in": objectIDs}}

	// Perform the DeleteMany operation
	result, err := collection.DeleteMany(c.Context(), filter)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete records", "details": err.Error()})
	}

	if result.DeletedCount < 1 {
		return c.Status(404).JSON(fiber.Map{"error": "No records deleted"})
	}

	// Print all the deleted IDs
	fmt.Println("Deleted IDS ==> ", objectIDs)
	fmt.Println("Deleted Count ==> ", result.DeletedCount)

	return c.Status(200).JSON(fiber.Map{"message": "Records deleted successfully"})
}
