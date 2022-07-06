package storage

import (
	"github.com/AXlIS/gofermart/internal/domain"
	"github.com/jmoiron/sqlx"
)

type Users interface {
	Create(username, passwordHash string) error
	Get(username string) (domain.User, error)
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
