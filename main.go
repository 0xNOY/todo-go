package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/naoya0x00/todo-go/db"
	"github.com/naoya0x00/todo-go/router"
)

func main() {
	db.Init("test.sqlite3")
	defer db.Close()

	router.Router(db.GetDB())
}
