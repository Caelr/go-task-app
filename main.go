package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	fmt.Println("Hello")
	app := fiber.New()

	// todos array
	todos := []Todo{}

	// TODO: finish routing
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"mes": "Hello world"})
	})

	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := Todo{} // FIX: {id: 0, completed: false, body: ''}
		if err := c.BodyParser(todo); err != nil {
			return err
		}

		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Enter a task"})
		}
	})

	log.Fatal(app.Listen(":8080"))
}
