package router

import (
	"flag"
	"fmt"
	"os"

	"github.com/naoya0x00/todo-go/controllers"
)

func Start(taskHandler *controllers.TaskHandler) {
	if len(os.Args) < 2 {
		fmt.Println("need subcommands")
		return
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
