package routes

import (
	"hello-world-app/controllers"
	"github.com/gofiber/fiber/v2"

)

func SetupRoutes(app *fiber.App) {
	app.Get("/", controllers.Hello)
}