package storage

import "github.com/jmoiron/sqlx"

type OrdersStorage struct {
	db *sqlx.DB
}

func NewOrdersStorage(db *sqlx.DB) *OrdersStorage {
	return &OrdersStorage{
		db: db,
	}
}