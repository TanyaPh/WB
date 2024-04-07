package service

import (
	"api/internal/cache"
	"api/internal/entity"
	"api/internal/repository"
	"encoding/json"

	"github.com/nats-io/stan.go"
)

type OrderService struct {
	repo     repository.Order
	cache	*cache.Cache
}

func newOrderService(repo repository.Order, cache *cache.Cache) *OrderService {
	return &OrderService{
		repo: repo,
		cache: cache,
	}
}

func (s *OrderService) Create(msg *stan.Msg) error {
	var input entity.Order
	err := json.Unmarshal(msg.Data, &input)
	if err != nil {
		return err
	}

	s.repo.Create(input)
	if err = s.repo.Create(input); err != nil {
		return err
	}

	return s.cache.Create(input)
}

func (s *OrderService) GetById(orderId string) (entity.Order, error) {
	return s.cache.GetById(orderId)
}
