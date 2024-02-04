package model

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type AppModel struct {
    Request *RequestStore
    Host *HostStore
}

func NewAppModel(path string) *AppModel {
    db, err := initializeStore(path) 

	if err != nil {
		log.Panicf("Error initializing database: %v", err)
	}

    request := NewRequestStore(db)
    host := NewHostStore(db)

    return &AppModel{Request: request, Host: host}
}

func initializeStore(path string) (*sql.DB, error) {

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
