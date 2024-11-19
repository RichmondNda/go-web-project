package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func Hello(c *fiber.Ctx) error {
	// Render the "Hello, World!" view
	return c.Render("hello", fiber.Map{
		"Message": "Hello, World Of Regis!",
	})
}
