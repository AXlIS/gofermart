package service

import "github.com/AXlIS/gofermart/internal/storage"

type UsersService struct {
	store storage.Users
}

func NewUsersService(store storage.Users) *UsersService {
	return &UsersService{store: store}
}
