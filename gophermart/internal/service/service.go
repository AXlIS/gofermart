package service

import (
	"github.com/AXlIS/gofermart/internal/storage"
	"github.com/AXlIS/gofermart/pkg/auth"
	"github.com/AXlIS/gofermart/pkg/hash"
)

type Users interface {
	Register(username, password string) error
	Login(username, password string) (*auth.Tokens, error)
	GetNewAccess(id string) (string, error)
}

type Orders interface {
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
