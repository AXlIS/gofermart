package service

import "github.com/AXlIS/gofermart/internal/storage"

type Users interface {
}

type Orders interface {
}

type Service struct {
	Users   Users
	Orders  Orders
}

func NewService(store *storage.Storage) *Service {
	return &Service{
		Users:   NewUsersService(store.Users),
		Orders:  NewOrdersService(store.Orders),
	}
}
