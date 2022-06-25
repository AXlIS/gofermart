package service

import (
	"github.com/AXlIS/gofermart/internal/storage"
	"github.com/AXlIS/gofermart/pkg/auth"
)

type Users interface {
}

type Orders interface {
}

type Service struct {
	Users   Users
	Orders  Orders
}

func NewService(store *storage.Storage, manager auth.TokenManager) *Service {
	return &Service{
		Users:   NewUsersService(store.Users, manager),
		Orders:  NewOrdersService(store.Orders),
	}
}
