package router

import (
	"log"

	"github.com/AlejandroWaiz/go-react-todoapp/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Router struct {
	db database.DatabaseImplementation
}

type RouterImplementation interface {
	ListenAndServe()
}

// Returns new Fiber router with given database as data repository.
func New(db database.DatabaseImplementation) RouterImplementation {

	return &Router{db: db}

}

func (r *Router) ListenAndServe() {

	app := fiber.New()

	app.Use(cors.New())

	app.Get("/Healtcheck", r.healtcheck)

	app.Post("/api/todos", r.postTodo)

	app.Get("/api/todos", r.getTodos)

	app.Patch("/api/todos/:id/done", r.setTodoAsDone)

	app.Delete("/api/todos/:id/delete", r.deleteTodo)

	log.Fatal(app.Listen(":4000"))

}
