package service

import (
	"errors"
	"fmt"
	"github.com/AXlIS/gofermart/internal/storage"
	"github.com/AXlIS/gofermart/pkg/auth"
	"github.com/AXlIS/gofermart/pkg/hash"
)

type UsersService struct {
	store        storage.Users
	tokenManager auth.TokenManager
	hasher       hash.Hasher
}

func NewUsersService(store storage.Users, tokenManager auth.TokenManager, hasher hash.Hasher) *UsersService {
	return &UsersService{store: store, tokenManager: tokenManager, hasher: hasher}
}

func (u *UsersService) Register(username, password string) error {
	passwordHash := u.hasher.Hash(password)

	if err := u.store.Create(username, passwordHash); err != nil {
		return err
	}

	return nil
}

func (u *UsersService) Login(username, password string) (*auth.Tokens, error) {
	passwordHash := u.hasher.Hash(password)

	user, err := u.store.Get(username)
	if err != nil {
		return nil, err
	}

	if passwordHash != user.Password {
		return nil, errors.New("incorrect password")
	}

	fmt.Println(user.Id)

	tokens, err := u.tokenManager.NewTokenPair(user.Id)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}

func (u *UsersService) GetNewAccess(id string) (string, error) {
	accessToken, err := u.tokenManager.NewAccessToken(id)
	if err != nil {
		fmt.Println("33")
		return "", err
	}

	return accessToken, nil
}
