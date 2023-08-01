package main

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteEmployees(c *fiber.Ctx) error {
	// Get the "ids" parameter from the URL query parameters
	ids := c.Query("ids")

	// Split the comma-separated IDs into individual IDs
	idArray := strings.Split(ids, ",")

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
