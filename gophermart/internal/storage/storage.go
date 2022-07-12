package storage

import (
	g "github.com/AXlIS/gofermart"
	"github.com/AXlIS/gofermart/internal/domain"
	"github.com/jmoiron/sqlx"
)

type Users interface {
	Create(username, passwordHash string) error
	Get(username string) (domain.User, error)
	GetBalance(userID string) (g.Balance, error)
	Debit(userID string, sum, order float32) error
	GetWithdrawalsInfo(userID string) ([]g.Withdrawal, error)
}

type Orders interface {
	Get(userID string) ([]g.Order, error)
	Load(userID string, number int) error
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
