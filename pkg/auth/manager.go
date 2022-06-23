package auth

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type TokenManager interface {
	NewToken(userID string, tokenTTL time.Duration) (string, error)
	NewTokenPair(userID string, accessTokenTTL, refreshTokenTTL time.Duration) (*Tokens, error)
	Parse(token string) (string, error)
}

type Manager struct {
	signingKey string
}

func NewManager(signingKey string) (*Manager, error) {
	if signingKey == "" {
		return nil, errors.New("empty singing key")
	}

	return &Manager{signingKey: signingKey}, nil
}

func (m *Manager) NewToken(userID string, tokenTTL time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(tokenTTL).Unix(),
		Subject:   userID,
	})

	return token.SignedString([]byte(m.signingKey))
}

func (m *Manager) NewTokenPair(userID string, accessTokenTTL, refreshTokenTTL time.Duration) (*Tokens, error) {
	accessToken, err := m.NewToken(userID, accessTokenTTL)
	if err != nil {
		return nil, err
	}

	refreshToken, err := m.NewToken(userID, refreshTokenTTL)
	if err != nil {
		return nil, err
	}

	return &Tokens{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

func (m *Manager) Parse(parseToken string) (string, error) {
	token, err := jwt.Parse(parseToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(m.signingKey), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("error get user claims from token")
	}

	return claims["sub"].(string), nil
}
