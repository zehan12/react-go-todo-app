package api

import (
	"strconv"

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

func HandleGetTodo(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "internal server error"})
	}

	if len(todos) == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "Todo list is empty!"})
	}

	// work in incremental id approach
	if id > len(todos) {
		return c.Status(400).JSON(fiber.Map{"error": ""})
	}

	todo := &models.Todo{}

	for i := 0; i <= len(todos); i++ {
		if todos[i].Id == id {
			*todo = todos[i]
			break
		}
	}
	return c.Status(200).JSON(todo)

}
