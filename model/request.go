package model

import "database/sql"

type Request struct {
	Id           int            `db:"id"`
	Method       string         `db:"method"`
	Endpoint     string         `db:"endpoint"`
	Body         sql.NullString `db:"body"`
	LastResponse sql.NullString `db:"last_response"`
}

type RequestStore struct {
	DB       *sql.DB
	requests []Request
}

func NewRequestStore(db *sql.DB) *RequestStore {
	return &RequestStore{DB: db}
}

func (store *RequestStore) GetRequests() ([]Request, error) {

	rows, err := store.DB.Query("SELECT id, method, endpoint, body, last_response FROM request;")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	store.requests = nil

	for rows.Next() {
		var req Request
		if err := rows.Scan(&req.Id, &req.Method, &req.Endpoint, &req.Body, &req.LastResponse); err != nil {
			return nil, err
		}
		store.requests = append(store.requests, req)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return store.requests, nil
}

func (s *RequestStore) GetRequest(index int) Request {
	return s.requests[index]
}

func (store *RequestStore) InsertRequest(r Request) error {
	stmt, err := store.DB.Prepare("INSERT INTO request (method, endpoint) VALUES (?, ?);")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(r.Method, r.Endpoint)

	if err != nil {
		return err
	}

	return nil
}

func (store *RequestStore) DeleteRequest(index int) error {

	stmt, err := store.DB.Prepare("DELETE FROM request WHERE id = ?;")

	if err != nil {
		return err
	}

	defer stmt.Close()

	req := store.requests[index]

	_, err = stmt.Exec(req.Id)

	if err != nil {
		return err
	}

	return nil
}

func (store *RequestStore) UpdateRequest(r Request) error {

	stmt, err := store.DB.Prepare("UPDATE request SET method = ?, endpoint = ? WHERE id = ?;")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(r.Method, r.Endpoint, r.Id)

	if err != nil {
		return err
	}

	return nil
}
