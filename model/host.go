package model

import "database/sql"

type HostStore struct {
    DB *sql.DB
}

func NewHostStore(db *sql.DB) *HostStore {
    return &HostStore{DB: db}
}
