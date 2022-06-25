package service

import (
	"github.com/AXlIS/gofermart/internal/storage"
	"github.com/AXlIS/gofermart/pkg/auth"
)

type UsersService struct {
	store        storage.Users
	tokenManager auth.TokenManager
}

func NewUsersService(store storage.Users, tokenManager auth.TokenManager) *UsersService {
	return &UsersService{store: store, tokenManager: tokenManager}
}
