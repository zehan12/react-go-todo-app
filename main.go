package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("App Started")

	const PORT = ":4000"

	app := fiber.New()

	log.Fatal(app.Listen(PORT))
}
