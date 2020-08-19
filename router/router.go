package router

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

	"github.com/naoya0x00/todo-go/controllers"
	"github.com/naoya0x00/todo-go/models"
)

func Start(db *sql.DB) {
	if len(os.Args) < 2 {
		fmt.Println("need subcommands")
		return
	}

	taskModel := models.TaskModel{
		DB: db,
	}
	taskHandler := controllers.TaskHandler{
		TaskModel: &taskModel,
	}

	switch os.Args[1] {
	case "add":
		cmd := flag.NewFlagSet("add", flag.ExitOnError)
		taskHandler.Add(cmd)

	case "del":
		cmd := flag.NewFlagSet("del", flag.ExitOnError)
		taskHandler.Delete(cmd)

	case "list":
		cmd := flag.NewFlagSet("list", flag.ExitOnError)
		taskHandler.Show(cmd)

	case "check":
		cmd := flag.NewFlagSet("check", flag.ExitOnError)
		taskHandler.UpdateDone(cmd)
	}
}
