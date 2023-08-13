package lead

import (
	"fiber-crm/database"

	"github.com/gofiber/fiber/v2"
	_ "modernc.org/sqlite"
)

type Lead struct {
	ID      uint   `gorm:"primaryKey; not null"`
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

func GetLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.Find(&lead, id)
	return c.JSON(lead)
}

func GetAllLead(c *fiber.Ctx) error {
	db := database.DBConn
	var leads []Lead
	db.Find(&leads)
	return c.JSON(leads)
}

func NewLead(c *fiber.Ctx) error {
	db := database.DBConn
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		c.Status(503).SendString("Err creating the new lead")
	}
	db.Create(&lead)
	return c.JSON(lead)
}

func DeleteLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.First(&lead, id)
	if lead.Name == "" {
		c.Status(500).SendString("No Lead Found with ID")
	}
	db.Delete(&lead)
	c.SendString("Lead successfully Deleted")
	return c.JSON(&lead)
}
