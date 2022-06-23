package storage

import "github.com/jmoiron/sqlx"

type Users interface {
}

type Orders interface {
}

type Storage struct {
	Users  Users
	Orders Orders
	db     *sqlx.DB
}

func NewStorage() *Storage {
	return &Storage{}
}
