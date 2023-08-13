package main

import (
	"fiber-crm/database"
	"fiber-crm/lead"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "modernc.org/sqlite"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetAllLead)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDb() {
	var err error
	database.DBConn, err = gorm.Open("sqlite", "leads.db")
	if err != nil {
		fmt.Println("Failed to connect to the database")
		panic(err)
	}
	fmt.Println("Connection opened to the database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database migration completed")
}

func main() {
	app := fiber.New()
	initDb()
	setupRoutes(app)
	app.Listen(":8081")
	defer database.DBConn.Close()
}
