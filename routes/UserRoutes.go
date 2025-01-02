package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sebastian-abraham/some-stuff/controller"
	"gorm.io/gorm"
)

func UserRoutes(db *gorm.DB, app *fiber.App) {
	api := app.Group("/api")
	api.Post("/user", controller.CreateUser(db))
}
