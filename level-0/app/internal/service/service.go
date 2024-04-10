package service

import (
	"api/internal/cache"
	"api/internal/entity"
	"api/internal/repository"
)

type Order interface {
	Create(msg []byte) error
	GetById(orderId string) (entity.Order, error)
}

type Service struct {
	Order Order
}

func NewService(repo *repository.Repository, cache *cache.Cache) *Service {
	return &Service{
		Order: newOrderService(repo.Orders, cache),
	}
}
