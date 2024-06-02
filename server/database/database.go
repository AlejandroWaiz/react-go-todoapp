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
	DeleteTodo(id int) ([]model.Todo, error)
	SetTodoAsDone(id int) ([]model.Todo, error)
}

// Return new database.
// Write 'local' to use local database. Local is the default database.
// Firestore database coming soon.
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
