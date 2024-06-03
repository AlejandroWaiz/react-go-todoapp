package router

import (
	"fmt"

	"github.com/AlejandroWaiz/go-react-todoapp/model"
	"github.com/gofiber/fiber/v2"
)

func (r *Router) postTodo(c *fiber.Ctx) error {

	todo := model.Todo{}

	if err := c.BodyParser(&todo); err != nil {
		return fiber.NewError(400, "[Router] Invalid Todo body")
	}

	AllTodo, err := r.db.PostTodo(todo)

	if err != nil {
		return fiber.NewError(400, err.Error())
	}

	return c.JSON(AllTodo)

}

func (r *Router) getTodos(c *fiber.Ctx) error {

	AllTodo, err := r.db.GetTodos()

	if err != nil {
		return fiber.NewError(400, err.Error())
	}

	return c.JSON(AllTodo)

}

func (r *Router) updateTodo(c *fiber.Ctx) error {

	todo := model.Todo{}

	fmt.Println(todo)

	if err := c.BodyParser(&todo); err != nil {
		return fiber.NewError(400, "[Router] Invalid Todo body")
	}

	AllTodo, err := r.db.UpdateTodo(todo)

	if err != nil {
		return fiber.NewError(400, err.Error())
	}

	return c.JSON(AllTodo)

}

func (r *Router) deleteTodo(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(401).SendString("Invalid id")
	}

	AllTodo, err := r.db.DeleteTodo(id)

	if err != nil {
		return fiber.NewError(400, err.Error())
	}

	return c.JSON(AllTodo)

}

func (r *Router) setTodoAsDone(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(401).SendString("Invalid id")
	}

	AllTodo, err := r.db.SetTodoAsDone(id)

	if err != nil {
		return fiber.NewError(400, err.Error())
	}

	return c.JSON(AllTodo)

}

func (r *Router) healtcheck(c *fiber.Ctx) error {

	return c.SendString("Healthcheck ok!")

}
