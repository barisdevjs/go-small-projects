package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var mg MongoInstance
var USER_NAME string
var DB_PASS string

const dbName = "employee"

type Employee struct {
	ID     string  `json:"id,omitempty" bson:"_id,omitempty"` // bson just for mongoDb to tell it is id
	Name   string  `json:"name"`
	Salary float64 `json:"salary"`
	Age    float64 `json:"age"`
}

func Connect() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	USER_NAME = os.Getenv("USER_NAME")
	DB_PASS = os.Getenv("DB_PASS")

	escapedUser := url.QueryEscape(USER_NAME)
	escapedPass := url.QueryEscape(DB_PASS)
	mongoURI := "mongodb+srv://" + escapedUser + ":" + escapedPass + "@cluster0.smzb3os.mongodb.net/?retryWrites=true&w=majority"

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI))
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return err
	}

	db := client.Database(dbName)
	mg = MongoInstance{
		Client: client,
		Db:     db,
	}
	fmt.Println(mongoURI)
	return nil
}

// fiber ~~ like express.js
func main() {

	if err := Connect(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New()

	app.Get("/employee", GetEmployees)

	app.Post("/employee", func(c *fiber.Ctx) error {
		collection := mg.Db.Collection("employees")

		employee := new(Employee)

		if err := c.BodyParser(employee); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		employee.ID = ""

		insertionResult, err := collection.InsertOne(c.Context(), employee)

		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
		createdRecord := collection.FindOne(c.Context(), filter)

		createdEmployee := &Employee{}
		createdRecord.Decode(createdEmployee)

		return c.Status(201).JSON(createdEmployee)

	})

	// app.Post("/employee")
	// app.Put("/employee/:id")
	// app.Delete("/employee/:id")

	log.Fatal(app.Listen(":8082"))

}
