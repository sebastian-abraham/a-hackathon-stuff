package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sebastian-abraham/some-stuff/database"
	"github.com/sebastian-abraham/some-stuff/routes"
)

func main() {

	db, err := database.NewConnection()
	if err != nil {
		log.Fatal("Could not load database")
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Document Section")
	})

	routes.UserRoutes(db, app)

	app.Listen(":8080")
}
