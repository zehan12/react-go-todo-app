package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("App Started")

	const PORT = ":3000"

	app := fiber.New()

	app.Listen(PORT)
}
