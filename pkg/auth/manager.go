package auth

import (
	"errors"
	"fmt"
	"github.com/AXlIS/gofermart/internal/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"time"
)

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type TokenManager interface {
	NewToken(userID string, tokenTTL time.Duration) (string, error)
	NewAccessToken(userID string) (string, error)
	NewTokenPair(userID string) (*Tokens, error)
	Parse(token string) (string, error)
}

type Manager struct {
	JWT config.JWTConfig
}

type Payload struct {
	ID        uuid.UUID `json:"id"`
	UserID    string    `json:"user_id"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

var ErrExpiredToken = errors.New("token has expired")

func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}

func NewPayload(userID string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		UserID:    userID,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

func NewManager(JWT config.JWTConfig) (*Manager, error) {
	if JWT.SigningKey == "" {
		return nil, errors.New("empty singing key")
	}

	return &Manager{
		JWT: JWT,
	}, nil
}

func (m *Manager) NewToken(userID string, tokenTTL time.Duration) (string, error) {

	payload, err := NewPayload(userID, tokenTTL)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	return token.SignedString([]byte(m.JWT.SigningKey))
}

func (m *Manager) NewAccessToken(userID string) (string, error) {
	return m.NewToken(userID, m.JWT.AccessTokenTTL)
}

func (m *Manager) NewTokenPair(userID string) (*Tokens, error) {
	accessToken, err := m.NewToken(userID, m.JWT.AccessTokenTTL)
	if err != nil {
		return nil, err
	}

	refreshToken, err := m.NewToken(userID, m.JWT.RefreshTokenTTL)
	if err != nil {
		return nil, err
	}

	return &Tokens{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

func (m *Manager) Parse(parseToken string) (string, error) {
	token, err := jwt.ParseWithClaims(parseToken, &Payload{}, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(m.JWT.SigningKey), nil
	})
	if err != nil {
		return "", err
	}

	payload, ok := token.Claims.(*Payload)
	if !ok {
		return "", errors.New("error get user claims from token")
	}

	return payload.UserID, nil
}
