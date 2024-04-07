package service

import (
	"api/internal/cache"
	"api/internal/entity"
	"api/internal/repository"

	"github.com/nats-io/stan.go"
)

type Order interface {
	Create(msg *stan.Msg) error
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
