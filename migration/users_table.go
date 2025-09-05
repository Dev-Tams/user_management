package migration

import (
    "database/sql"
    "log"
)

var createTable = `
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT,
    email TEXT,
    password TEXT,
    role TEXT
);`

func InitDB(db *sql.DB) {
    _, err := db.Exec(createTable)
    if err != nil {
        log.Fatal("failed to create table:", err)
    }
}
