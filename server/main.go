package main

import (
	"github.com/AlejandroWaiz/go-react-todoapp/database"
	"github.com/AlejandroWaiz/go-react-todoapp/router"
)

func main() {

	db, err := database.New("local")

	if err != nil {
		panic(err)
	}

	router := router.New(db)

	router.ListenAndServe()

}
