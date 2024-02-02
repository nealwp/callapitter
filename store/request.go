package store

import "database/sql"

type Request struct {
    Id int `db:"id"`
    Method string `db:"method"`
    Endpoint string `db:"endpoint"`
    Body sql.NullString `db:"body"`
    LastResponse sql.NullString `db:"last_response"`
}

type RequestStore struct {
    DB *sql.DB
}

func NewRequestStore(db *sql.DB) *RequestStore {
    return &RequestStore{DB: db}
}

func (store *RequestStore) GetRequests() ([]Request, error) {
    var requests []Request

    rows, err := store.DB.Query(`SELECT id, method, endpoint, body, last_response FROM request;`)

    if err != nil {
        return nil, err
    }

    defer rows.Close()

    for rows.Next() {
        var req Request 
        if err := rows.Scan(&req.Id, &req.Method, &req.Endpoint, &req.Body, &req.LastResponse); err != nil {
            return nil, err
        }
        requests = append(requests, req)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return requests, nil
}
