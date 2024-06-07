package api

import "github.com/gofiber/fiber/v2"

func HandleCreateTodo(c *fiber.Ctx) error {
	return c.JSON("todo post")
}
