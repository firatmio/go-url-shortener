package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("sqlite3", "./urls.db")
	if err != nil {
		log.Fatal(err)
	}

	// Create table if not exists
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS urls (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            short TEXT NOT NULL UNIQUE,
            original TEXT NOT NULL
        );
    `)
	if err != nil {
		log.Fatal(err)
	}
}
