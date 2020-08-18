package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/naoya0x00/todo-go/db"
	"github.com/naoya0x00/todo-go/router"
)

func main() {
	db.Init("test.sqlite3")
	defer db.Close()
	sql := `
		CREATE TABLE IF NOT EXISTS tasks(
			title STRING,
			done INTEGER
		)
	`
	db.GetDB().Exec(sql)

	router.Router(db.GetDB())
}
