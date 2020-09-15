package main

import (
	"log"

	"github.com/gofiber/fiber"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/books")
}

func main() {
	app := fiber.New()

	app.Get("/", helloWorld)

	log.Fatal(app.Listen(":3000"))
}
