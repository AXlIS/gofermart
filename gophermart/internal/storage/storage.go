package storage

import "github.com/jmoiron/sqlx"

type Users interface {
}

type Orders interface {
}

type Storage struct {
	Users  Users
	Orders Orders
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{
		Users:  NewUsersStorage(db),
		Orders: NewOrdersStorage(db),
	}
}
