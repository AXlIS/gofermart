package service

import (
	"errors"
	g "github.com/AXlIS/gofermart"
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

	tokens, err := u.tokenManager.NewTokenPair(user.ID)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}

func (u *UsersService) GetNewAccess(id string) (string, error) {
	accessToken, err := u.tokenManager.NewAccessToken(id)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (u *UsersService) GetBalance(userID string) (g.Balance, error) {
	balance, err := u.store.GetBalance(userID)
	if err != nil {
		return g.Balance{}, err
	}
	return balance, nil
}

func (u *UsersService) Debit(userID string, sum, order float32) error {
	return u.store.Debit(userID, sum, order)
}

func (u *UsersService) GetWithdrawalsInfo(userID string) ([]g.Withdrawal, error) {
	withdrawal, err := u.store.GetWithdrawalsInfo(userID)
	if err != nil {
		return nil, err
	}
	return withdrawal, nil
}
