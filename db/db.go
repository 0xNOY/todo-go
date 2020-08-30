package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func Init(dbPath string) (err error) {
	db, err = sql.Open("sqlite3", dbPath)
	return
}

func GetDB() *sql.DB {
	return db
}

func Close() (err error) {
	err = db.Close()
	return
}
