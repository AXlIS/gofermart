package storage

import (
	"fmt"
	g "github.com/AXlIS/gofermart"
	"github.com/AXlIS/gofermart/internal/domain"
	"github.com/jmoiron/sqlx"
	"time"
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
	query := fmt.Sprintf("SELECT id, password FROM %s WHERE username=$1", usersTable)
	err := u.db.Get(&user, query, username)
	return user, err
}

func (u *UsersStorage) GetBalance(userID string) (g.Balance, error) {
	var (
		accrual float32
		amount  float32
	)
	getAccrualQuery := fmt.Sprintf(`SELECT SUM(accrual)
									      FROM %s
                                          WHERE user_id = $1;`, ordersTable)
	if err := u.db.Get(&accrual, getAccrualQuery, userID); err != nil {
		return g.Balance{}, err
	}

	getWithdrawnQuery := fmt.Sprintf(`SELECT SUM(amount)
									        FROM %s
                                            WHERE user_id = $1;`, withdrawalTable)
	if err := u.db.Get(&amount, getWithdrawnQuery, userID); err != nil {
		return g.Balance{}, err
	}

	return g.Balance{
		Current:   accrual - amount,
		Withdrawn: amount,
	}, nil
}

func (u *UsersStorage) Debit(userID string, sum, order float32) error {
	query := fmt.Sprintf("INSERT INTO %s (user_id, order_id, amount, processed_at) VALUES ($1, $2, $3, $4)", withdrawalTable)
	if _, err := u.db.Exec(query, userID, order, sum, time.Now().Unix()); err != nil {
		return err
	}
	return nil
}

func (u *UsersStorage) GetWithdrawalsInfo(userID string) ([]g.Withdrawal, error) {
	var withdrawals []g.Withdrawal

	query := fmt.Sprintf("SELECT order_id as order, amount as sum, processed_at FROM %s WHERE user_id=$1", withdrawalTable)

	if err := u.db.Select(&withdrawals, query, userID); err != nil {
		return nil, err
	}

	return withdrawals, nil
}
