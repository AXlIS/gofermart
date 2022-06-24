package storage

import "github.com/jmoiron/sqlx"

type UsersStorage struct {
	db *sqlx.DB
}

func NewUsersStorage(db *sqlx.DB) *UsersStorage {
	return &UsersStorage{
		db: db,
	}
}
