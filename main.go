package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/todo/react-go-todo-app/api"
)

func main() {
	fmt.Println("App Started")

	const PORT = ":3000"

	app := fiber.New()

	// group routes to v1
	apiv1 := app.Group("/api/v1")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"message": "server is running!"})
	})

	apiv1.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{"message": "Welcome to Todo Api!"})
	})

	// todo routes
	todoRoutes := apiv1.Group("/todo")
	todoRoutes.Get("/", api.HandleGetTodos)
	todoRoutes.Post("/", api.HandleCreateTodo)
	todoRoutes.Get("/:id", api.HandleGetTodo)

	log.Fatal(app.Listen(PORT))
}
