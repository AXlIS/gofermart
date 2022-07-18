package storage

import (
	"fmt"
	g "github.com/AXlIS/gofermart"
	"github.com/jmoiron/sqlx"
	"time"
)

type OrdersStorage struct {
	db *sqlx.DB
}

func NewOrdersStorage(db *sqlx.DB) *OrdersStorage {
	return &OrdersStorage{
		db: db,
	}
}

func (o *OrdersStorage) Get(userID string) ([]g.Order, error) {
	var orders []g.Order
	query := fmt.Sprintf("SELECT id, status, accrual, uploaded_at FROM %s WHERE user_id=$1", ordersTable)

	if err := o.db.Select(&orders, query, userID); err != nil {
		return nil, err
	}

	fmt.Println(orders)

	return orders, nil
}

func (o *OrdersStorage) Load(userID string, number int) error {
	query := fmt.Sprintf("INSERT INTO %s (id, uploaded_at, user_id) VALUES ($1, $2, $3)", ordersTable)

	if _, err := o.db.Exec(query, number, time.Now().Unix(), userID); err != nil {
		return err
	}

	return nil
}
