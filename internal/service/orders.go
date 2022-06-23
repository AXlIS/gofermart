package service

import "github.com/AXlIS/gofermart/internal/storage"

type OrdersService struct {
	store storage.Orders
}

func NewOrdersService(store storage.Orders) *OrdersService {
	return &OrdersService{store: store}
}
