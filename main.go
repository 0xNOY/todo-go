package main

import (
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/naoya0x00/todo-go/controllers"
	"github.com/naoya0x00/todo-go/db"
	"github.com/naoya0x00/todo-go/models"
	"github.com/naoya0x00/todo-go/router"
)

func main() {
	err := db.Init("todo-go.sqlite3")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		err = db.Close()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		return
	}

	taskHandler := controllers.TaskHandler{
		TaskModel: &models.TaskModel{
			DB: db.GetDB(),
		},
	}
	taskHandler.TaskModel.Init()

	router.Route(&taskHandler)

	err = db.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
