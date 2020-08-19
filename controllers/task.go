package controllers

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"

	"github.com/naoya0x00/todo-go/models"
)

type TaskHandler struct {
	TaskModel *models.TaskModel
}

func (h *TaskHandler) Add(cmd *flag.FlagSet) {
	cmd.Parse(os.Args[2:])

	for _, title := range cmd.Args() {
		h.TaskModel.Add(title)
	}
	fmt.Println("added")
}

func (h *TaskHandler) Delete(cmd *flag.FlagSet) {
	deleteAllCompleted := cmd.Bool("d", false, "delete all completed tasks")
	cmd.Parse(os.Args[2:])

	if *deleteAllCompleted {
		h.TaskModel.DeleteAllCompleted()
	} else {
		var ids []uint64
		for _, strID := range cmd.Args() {
			id, err := strconv.ParseUint(strID, 10, 64)
			if err != nil {
				fmt.Println("invalid value")
				return
			}
			ids = append(ids, id)
		}

		for _, id := range ids {
			h.TaskModel.Delete(id)
		}
	}

	fmt.Println("deleted")
}

func (h *TaskHandler) Show(cmd *flag.FlagSet) {
	showType := cmd.Int("s", -1, "0: show only completed tasks, 1: show only not completed tasks")
	cmd.Parse(os.Args[2:])

	var tasks []models.Task
	if *showType == 0 || *showType == 1 {
		tasks = h.TaskModel.FindByDone(*showType != 0)
	} else {
		tasks = h.TaskModel.All()
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Title", "Done"})
	for _, task := range tasks {
		strTask := []string{
			strconv.FormatUint(task.ID, 10),
			task.Title,
			strconv.FormatBool(task.Done),
		}
		table.Append(strTask)
	}
	table.Render()
}

func (h *TaskHandler) UpdateDone(cmd *flag.FlagSet) {
	cancel := cmd.Bool("c", false, "cancel update done")
	cmd.Parse(os.Args[2:])

	var ids []uint64
	for _, strID := range cmd.Args() {
		id, err := strconv.ParseUint(strID, 10, 64)
		if err != nil {
			fmt.Println("invalid value")
			return
		}
		ids = append(ids, id)
	}

	for _, id := range ids {
		h.TaskModel.UpdateDone(id, *cancel != true)
	}

	fmt.Println("updated")
}
