package models

import (
	"database/sql"
)

type Task struct {
	ID    uint64
	Title string
	Done  bool
}

type TaskModel struct {
	DB *sql.DB
}

func (m *TaskModel) Init() (err error) {
	sql := `
		CREATE TABLE IF NOT EXISTS tasks(
			title STRING,
			done INTEGER
		)
	`
	_, err = m.DB.Exec(sql)
	return
}

func (m *TaskModel) Add(title string) (err error) {
	sql := "INSERT INTO tasks (title, done) VALUES (?, ?)"
	_, err = m.DB.Exec(sql, title, false)
	return
}

func (m *TaskModel) All() (tasks []Task, err error) {
	sql := "SELECT rowid, * FROM tasks"
	rows, err := m.DB.Query(sql)
	if err != nil {
		return
	}

	for rows.Next() {
		var task Task
		rows.Scan(&task.ID, &task.Title, &task.Done)
		tasks = append(tasks, task)
	}
	return
}

func (m *TaskModel) FindByDone(done bool) (tasks []Task, err error) {
	sql := "SELECT rowid, * FROM tasks WHERE done = ?"
	rows, err := m.DB.Query(sql, done)
	if err != nil {
		return
	}

	for rows.Next() {
		var task Task
		rows.Scan(&task.ID, &task.Title, &task.Done)
		tasks = append(tasks, task)
	}
	return
}

func (m *TaskModel) Delete(id uint64) (err error) {
	sql := "DELETE FROM tasks WHERE rowid = ?"
	_, err = m.DB.Exec(sql, id)
	return
}

func (m *TaskModel) DeleteAllCompleted() (err error) {
	sql := "DELETE FROM tasks WHERE done = ?"
	_, err = m.DB.Exec(sql, true)
	return
}

func (m *TaskModel) UpdateDone(id uint64, done bool) (err error) {
	sql := "UPDATE tasks SET done = ? WHERE rowid = ?"
	_, err = m.DB.Exec(sql, done, id)
	return
}
