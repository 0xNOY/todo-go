package router

import (
	"github.com/naoya0x00/todo-go/controllers"
)

func Route(taskHandler *controllers.TaskHandler) {
	router := Router{}

	router.AddCommand("ad", "<taskTitle: string> | add task", taskHandler.Add)
	router.AddCommand("rm", "<taskID: int> | delete task", taskHandler.Delete)
	router.AddCommand("st", "| show tasks", taskHandler.Show)
	router.AddCommand("ck", "<taskID: int> | complete task", taskHandler.UpdateDone)

	router.Start()
}
