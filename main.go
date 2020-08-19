package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/naoya0x00/todo-go/db"
	"github.com/naoya0x00/todo-go/router"
)

func main() {
	db.Init("todo-go.sqlite3")
	defer db.Close()

	router.Start(db.GetDB())
}
