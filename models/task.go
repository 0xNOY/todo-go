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

func (m *TaskModel) Init() {
	sql := `
		CREATE TABLE IF NOT EXISTS tasks(
			title STRING,
			done INTEGER
		)
	`
	m.DB.Exec(sql)
}

func (m *TaskModel) Add(title string) {
	sql := "INSERT INTO tasks (title, done) VALUES (?, ?)"
	m.DB.Exec(sql, title, false)
}

func (m *TaskModel) All() (tasks []Task) {
	sql := "SELECT rowid, * FROM tasks"
	rows, _ := m.DB.Query(sql)
	for rows.Next() {
		var task Task
		rows.Scan(&task.ID, &task.Title, &task.Done)
		tasks = append(tasks, task)
	}
	return
}

func (m *TaskModel) FindByDone(done bool) (tasks []Task) {
	sql := "SELECT rowid, * FROM tasks WHERE done = ?"
	rows, _ := m.DB.Query(sql, done)
	for rows.Next() {
		var task Task
		rows.Scan(&task.ID, &task.Title, &task.Done)
		tasks = append(tasks, task)
	}
	return
}

func (m *TaskModel) Delete(id uint64) {
	sql := "DELETE FROM tasks WHERE rowid = ?"
	m.DB.Exec(sql, id)
}

func (m *TaskModel) DeleteAllCompleted() {
	sql := "DELETE FROM tasks WHERE done = ?"
	m.DB.Exec(sql, true)
}

func (m *TaskModel) UpdateDone(id uint64, done bool) {
	sql := "UPDATE tasks SET done = ? WHERE rowid = ?"
	m.DB.Exec(sql, done, id)
}
