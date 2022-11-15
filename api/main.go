package main

import (
	"math/rand"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	app.Get("/lucky", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"luckyNumber": rand.Intn(10),
		})
	})

	app.Listen(":3001")
}
