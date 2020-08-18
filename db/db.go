package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func Init(dbPath string) {
	db, _ = sql.Open("sqlite3", dbPath)

	sql := `
		CREATE TABLE IF NOT EXISTS tasks(
			title STRING,
			done INTEGER
		)
	`
	db.Exec(sql)
}

func GetDB() *sql.DB {
	return db
}

func Close() {
	db.Close()
}
