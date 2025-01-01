package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Document Section")
	})

	app.Listen(":8080")
}
