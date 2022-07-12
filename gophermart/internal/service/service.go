package service

import (
	g "github.com/AXlIS/gofermart"
	"github.com/AXlIS/gofermart/internal/storage"
	"github.com/AXlIS/gofermart/pkg/auth"
	"github.com/AXlIS/gofermart/pkg/hash"
)

type Users interface {
	Register(username, password string) error
	Login(username, password string) (*auth.Tokens, error)
	GetNewAccess(id string) (string, error)
	GetBalance(userID string) (g.Balance, error)
	Debit(userID string, sum, order float32) error
	GetWithdrawalsInfo(userID string) ([]g.Withdrawal, error)
}

type Orders interface {
	Get(userID string) ([]g.Order, error)
	Load(userID string, number int) error
}

type Service struct {
	Users   Users
	Orders  Orders
}

func NewService(store *storage.Storage, manager auth.TokenManager, hasher hash.Hasher) *Service {
	return &Service{
		Users:   NewUsersService(store.Users, manager, hasher),
		Orders:  NewOrdersService(store.Orders),
	}
}
