package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/todo/react-go-todo-app/models"
)

var (
	todos = []models.Todo{}
)

func HandleGetTodos(c *fiber.Ctx) error {
	return c.Status(200).JSON(todos)
}

func HandleCreateTodo(c *fiber.Ctx) error {
	todo := &models.Todo{}

	if err := c.BodyParser(todo); err != nil {
		return err
	}

	if todo.Body == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Todo body is required"})
	}

	todo.Id = len(todos) + 1
	todos = append(todos, *todo)

	return c.Status(201).JSON(todo)
}
