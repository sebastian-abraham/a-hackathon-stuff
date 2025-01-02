package controller

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/sebastian-abraham/some-stuff/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// Create a new instance of the User model
		user := new(models.User)

		// Parse the body contents
		err := c.BodyParser(user)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Error parsing request body",
				"error":   err.Error(),
			})
		}

		// Validate required fields
		if user.Email == "" || user.Password == "" {
			return c.Status(422).JSON(fiber.Map{
				"message": "Email and password are required",
			})
		}

		// Hash the password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Error hashing password",
				"error":   err.Error(),
			})
		}
		user.Password = string(hashedPassword)

		// Insert the user into the database
		err = db.Create(user).Error
		if err != nil {

			// Check if the error is a duplicate entry
			if strings.Contains(err.Error(), "(SQLSTATE 23505)") {
				return c.Status(409).JSON(fiber.Map{
					"message": "Email already taken",
				})
			}

			return c.Status(501).JSON(fiber.Map{
				"message": "Error inserting into database",
				"error":   err.Error(),
			})
		}

		return c.Status(201).JSON(fiber.Map{
			"message": "User created successfully",
		})
	}
}
