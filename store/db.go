package store

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func InitializeStore(path string) (*sql.DB, error) {

    db, err := sql.Open("sqlite3", path)

    if err != nil {
        return nil, err
    }

    const requestTableSql = `CREATE TABLE IF NOT EXISTS request (id INTEGER NOT NULL PRIMARY KEY, method TEXT NOT NULL, endpoint TEXT NOT NULL, body TEXT, last_response TEXT);`

    if _, err := db.Exec(requestTableSql); err != nil {
        return nil, err
    }
    
    const hostTableSql = `CREATE TABLE IF NOT EXISTS host (id INTEGER NOT NULL PRIMARY KEY, name TEXT NOT NULL);`

    if _, err := db.Exec(hostTableSql); err != nil {
        return nil, err
    }

    return db, nil
}

func readSqlFile(path string) (string, error) {
    content, err := os.ReadFile(path)
    if err != nil {
        return "", err
    }

    return string(content), nil
}
