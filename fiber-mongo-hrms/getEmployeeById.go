package main

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetEmployee(c *fiber.Ctx) error {
	employeeID, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.SendStatus(400)
	}

	query := bson.D{{Key: "_id", Value: employeeID}}
	result := mg.Db.Collection("employees").FindOne(c.Context(), &query)

	var employee Employee
	err = result.Decode(&employee)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.Status(200).JSON(employee)

}
