package database

import (
	"github.com/AlejandroWaiz/go-react-todoapp/database/local"
	"github.com/AlejandroWaiz/go-react-todoapp/model"
)

type Database struct {
}

type DatabaseImplementation interface {
	PostTodo(model.Todo) ([]model.Todo, error)
	UpdateTodo(todo model.Todo) ([]model.Todo, error)
	GetTodos() ([]model.Todo, error)
	DeleteTodo(todo model.Todo) ([]model.Todo, error)
	SetTodoAsDone(todo model.Todo) ([]model.Todo, error)
}

func New(database string) (DatabaseImplementation, error) {

	//TODO: implementar el devolver el tipo de base segun el string que te pasen al invocar la funcion
	switch database {

	//TODO: crear logica para utilizar firestore
	//case "firestore":
	//	return Database{}, nil

	case "local":

		db, _ := local.NewDatabase()

		return &db, nil

	default:

		db, _ := local.NewDatabase()

		return &db, nil

	}

}
