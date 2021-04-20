package main

import (
	"github.com/gofiber/fiber/v2"

	"net/http"
)

func main() {
    app := fiber.New()

    app.Get("/hello", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World 👋!")
    })

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusNotFound)
	})

    app.Listen(":3000")
}