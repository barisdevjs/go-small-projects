package main

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateEmployee(c *fiber.Ctx) error {
	idParam := c.Params("id")
	employeeID, err := primitive.ObjectIDFromHex(idParam)

	if err != nil {
		return c.SendStatus(400)
	}

	employee := new(Employee)

	if err := c.BodyParser(employee); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	query := bson.D{{Key: "_id", Value: employeeID}}
	update := bson.D{
		{Key: "$set",
			Value: bson.D{
				{Key: "name", Value: employee.Name},
				{Key: "age", Value: employee.Age},
				{Key: "salary", Value: employee.Salary},
			},
		},
	}

	err = mg.Db.Collection("employees").FindOneAndUpdate(c.Context(), query, update).Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.SendStatus(400)
		}
		return c.SendStatus(500)
	}

	employee.ID = idParam
	return c.Status(200).JSON(employee)

}
