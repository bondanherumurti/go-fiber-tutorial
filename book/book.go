package book

import (
	"github.com/gofiber/fiber"
)

func GetBooks(c *fiber.Ctx) error {
	return c.SendString("All Books")
}

func GetBook(c *fiber.Ctx) error {
	return c.SendString("A Single Book")
}

func UpdateBook(c *fiber.Ctx) error {
	return c.SendString("Update a single book")
}

func UpdateBook(c *fiber.Ctx) error {
	return c.SendString("Delete a single book")
}
