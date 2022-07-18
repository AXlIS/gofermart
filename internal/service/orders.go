package service

import (
	g "github.com/AXlIS/gofermart"
	"github.com/AXlIS/gofermart/internal/storage"
)

type OrdersService struct {
	store storage.Orders
}

func NewOrdersService(store storage.Orders) *OrdersService {
	return &OrdersService{store: store}
}

func (s *OrdersService) Get(userID string) ([]g.Order, error) {
	orders, err := s.store.Get(userID)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (s *OrdersService) Load(userID string, number int) error {
	return s.store.Load(userID, number)
}
