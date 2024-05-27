package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/Healtcheck", healtcheck)

	app.Post("/api/todos", postTodo)

	log.Fatal(app.Listen(":4000"))

}

var AllTodo []Todo

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
	Body  string `json:"body"`
}

func postTodo(c *fiber.Ctx) error {

	todo := Todo{}

	if err := c.BodyParser(&todo); err != nil {
		return err
	}

	todo.ID = len(AllTodo) + 1

	AllTodo = append(AllTodo, todo)

	return c.JSON(AllTodo)

}

func healtcheck(c *fiber.Ctx) error {

	return c.SendString("Healthcheck ok!")

}
