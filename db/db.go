package db

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitializeDB(path string) error {
    var err error
    db, err = sql.Open("sqlite3", path)
    if err != nil {
        return err
    }

    defer db.Close()

    const createTableSql = `
        CREATE TABLE IF NOT EXISTS requests (
            id INTEGER NOT NULL PRIMARY KEY,
            method TEXT NOT NULL,
            endpoint TEXT NOT NULL,
            body TEXT,
            last_response TEXT
        );`

    if _, err := db.Exec(createTableSql); err != nil {
        return err
    }

    return nil
}

func readSqlFile(path string) (string, error) {
    content, err := os.ReadFile(path)
    if err != nil {
        return "", err
    }

    return string(content), nil
}
