package storage

import "github.com/jmoiron/sqlx"

func NewPostgresDB(uri string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", uri)
	if err != nil {
		return nil, nil
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
