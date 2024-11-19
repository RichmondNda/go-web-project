package main

import (
    "hello-world-app/routes"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/template/html/v2" 
)

func main() {

    // Load views HTML templates
    engine := html.New("./views", ".html")

    app := fiber.New(fiber.Config{
        Views: engine,
    })

    // Initialize routes
    routes.SetupRoutes(app)

    // Start server
    app.Listen(":3000")
}