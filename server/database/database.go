package database

import (
	"github.com/AlejandroWaiz/go-react-todoapp/database/local"
	"github.com/AlejandroWaiz/go-react-todoapp/model"
)

type Database struct {
}

// TODO: Cambiar las funciones Delete y SetTodoAsDone para que trabajen solo con el id ya que no es necesario el cuerpo completo
type DatabaseImplementation interface {
	PostTodo(model.Todo) ([]model.Todo, error)
	UpdateTodo(todo model.Todo) ([]model.Todo, error)
	GetTodos() ([]model.Todo, error)
	DeleteTodo(todo model.Todo) ([]model.Todo, error)
	SetTodoAsDone(todo model.Todo) ([]model.Todo, error)
}

func New(database string) (DatabaseImplementation, error) {

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
