package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func Init(dbPath string) {
	db, _ = sql.Open("sqlite3", dbPath)
}

func GetDB() *sql.DB {
	return db
}

func Close() {
	db.Close()
}
