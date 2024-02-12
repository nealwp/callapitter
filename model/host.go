package model

import "database/sql"

type Host struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}

type HostStore struct {
	DB *sql.DB
}

func NewHostStore(db *sql.DB) *HostStore {
	return &HostStore{DB: db}
}

func (store *HostStore) InsertHost(h Host) error {
	stmt, err := store.DB.Prepare("INSERT INTO host (name) VALUES (?);")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(h.Name)

	if err != nil {
		return err
	}

	return nil
}

func (store *HostStore) GetHosts() ([]Host, error) {
	var hosts []Host

	rows, err := store.DB.Query("SELECT id, name FROM host;")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var host Host
		if err := rows.Scan(&host.Id, &host.Name); err != nil {
			return nil, err
		}
		hosts = append(hosts, host)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return hosts, nil
}

func (s *HostStore) Delete(host Host) error {
    stmt, err := s.DB.Prepare("DELETE FROM host WHERE id = ?")
    if err != nil {
        return err
    }
    defer stmt.Close()

    _, err = stmt.Exec(host.Id)

    if err != nil {
        return err
    }

    return nil
}
