package main

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetEmployees(c *fiber.Ctx) error {
	query := bson.D{{}}
	cursor, err := mg.Db.Collection("employees").Find(c.Context(), query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	var employees []Employee = make([]Employee, 0)

	if err := cursor.All(c.Context(), &employees); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(employees)
}
