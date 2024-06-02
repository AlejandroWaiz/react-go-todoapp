package local

import (
	"errors"
	"log"

	"github.com/AlejandroWaiz/go-react-todoapp/model"
)

var AllTodo []model.Todo

type LocalDatabase struct {
}

func NewDatabase() (LocalDatabase, error) {

	log.Println("Running database on local!!!")

	return LocalDatabase{}, nil

}

func (l *LocalDatabase) GetTodos() ([]model.Todo, error) {

	return AllTodo, nil

}

func (l *LocalDatabase) PostTodo(todo model.Todo) ([]model.Todo, error) {

	todo.ID = len(AllTodo) + 1

	AllTodo = append(AllTodo, todo)

	return AllTodo, nil

}

func (l *LocalDatabase) UpdateTodo(todo model.Todo) ([]model.Todo, error) {

	var founded bool

	for i := range AllTodo {

		if AllTodo[i].ID == todo.ID {

			AllTodo[i] = todo

			founded = true

			break

		}

	}

	if founded {
		return AllTodo, nil
	} else {
		return AllTodo, errors.New("Todo not founded")
	}

}

func (l *LocalDatabase) DeleteTodo(id int) ([]model.Todo, error) {

	var founded bool

	var deleteElement = func(s []model.Todo, i int) []model.Todo {
		s[i] = s[len(s)-1]
		return s[:len(s)-1]
	}

	for i, t := range AllTodo {

		if t.ID == id {

			founded = true

			AllTodo = deleteElement(AllTodo, i)

		}

	}

	if founded {
		return AllTodo, nil
	} else {
		return AllTodo, errors.New("Todo not founded")
	}

}

func (l *LocalDatabase) SetTodoAsDone(id int) ([]model.Todo, error) {

	var founded bool

	for i, t := range AllTodo {

		if t.ID == id {

			AllTodo[i].Done = true

			founded = true

			break
		}

	}

	if founded {
		return AllTodo, nil
	} else {
		return AllTodo, errors.New("Todo not founded")
	}

}
