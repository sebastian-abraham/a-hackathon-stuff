package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sebastian-abraham/some-stuff/database"
	"github.com/sebastian-abraham/some-stuff/models"
	"github.com/sebastian-abraham/some-stuff/routes"
)

func main() {
	// Connect to the database
	db, err := database.NewConnection()
	if err != nil {
		log.Fatal("Could not load database")
	}
	// Create an instance of Fiber
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Document Section")
	})

	// Migrate the User model
	models.MigrateUser(db)

	// Define the routes
	routes.UserRoutes(db, app)

	// Start the server
	app.Listen(":8080")
}
