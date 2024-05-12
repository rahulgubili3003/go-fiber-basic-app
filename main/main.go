package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	// Initialize
	app := fiber.New()

	// Accepts the request on this path and return the Json value as mentioned
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, we are using Go Fiber!!",
		})
	})

	// Listen on Port 3000
	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
