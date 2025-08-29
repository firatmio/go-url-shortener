package main

import (
    "database/sql"
    "log"

    _ "modernc.org/sqlite"
)

var db *sql.DB

func InitDB() {
    var err error
    db, err = sql.Open("sqlite", "./urls.db")
    if err != nil {
        log.Fatal(err)
    }

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
