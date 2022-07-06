package storage

import (
	"fmt"
	"github.com/AXlIS/gofermart/internal/domain"
	"github.com/jmoiron/sqlx"
)

type UsersStorage struct {
	db *sqlx.DB
}

func NewUsersStorage(db *sqlx.DB) *UsersStorage {
	return &UsersStorage{
		db: db,
	}
}

func (u *UsersStorage) Create(username, passwordHash string) error {
	query := fmt.Sprintf("INSERT INTO %s (username, password) VALUES ($1, $2)", usersTable)

	if _, err := u.db.Exec(query, username, passwordHash); err != nil {
		return err
	}

	return nil
}

func (u *UsersStorage) Get(username string) (domain.User, error) {
	var user domain.User
	query := fmt.Sprintf("SELECT password FROM %s WHERE username=$1", usersTable)
	err := u.db.Get(&user, query, username)
	return user, err
}
