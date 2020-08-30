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
	err := cmd.Parse(os.Args[2:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	if len(cmd.Args()) <= 0 {
		fmt.Fprintln(os.Stderr, "argument is empty")
		return
	}

	for _, title := range cmd.Args() {
		err = h.TaskModel.Add(title)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
	}

	fmt.Println("added")
}

func (h *TaskHandler) Delete(cmd *flag.FlagSet) {
	deleteAllCompleted := cmd.Bool("d", false, "delete all completed tasks")
	err := cmd.Parse(os.Args[2:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	if *deleteAllCompleted {
		err = h.TaskModel.DeleteAllCompleted()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
	} else {
		if len(cmd.Args()) <= 0 {
			fmt.Fprintln(os.Stderr, "argument is empty")
			return
		}

		var ids []uint64
		for _, strID := range cmd.Args() {
			id, err := strconv.ParseUint(strID, 10, 64)
			if err != nil {
				fmt.Fprintln(os.Stderr, "invalid value")
				return
			}
			ids = append(ids, id)
		}

		for _, id := range ids {
			err = h.TaskModel.Delete(id)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
		}
	}

	fmt.Println("deleted")
}

func (h *TaskHandler) Show(cmd *flag.FlagSet) {
	showType := cmd.Int("s", -1, "0: show only completed tasks, 1: show only not completed tasks")
	err := cmd.Parse(os.Args[2:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	var tasks []models.Task
	if *showType == 0 || *showType == 1 {
		tasks, err = h.TaskModel.FindByDone(*showType != 0)
	} else {
		tasks, err = h.TaskModel.All()
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
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
	err := cmd.Parse(os.Args[2:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	if len(cmd.Args()) <= 0 {
		fmt.Fprintln(os.Stderr, "argument is empty")
		return
	}

	var ids []uint64
	for _, strID := range cmd.Args() {
		id, err := strconv.ParseUint(strID, 10, 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, "invalid value")
			return
		}
		ids = append(ids, id)
	}

	for _, id := range ids {
		err = h.TaskModel.UpdateDone(id, *cancel != true)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
	}

	fmt.Println("updated")
}
