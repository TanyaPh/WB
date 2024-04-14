package cache

import (
	"api/internal/entity"
	"api/internal/repository"
	"errors"
	"sync"
)

type Cache struct {
	data map[string]entity.Order
	mu   sync.RWMutex
}

func NewCache(repo *repository.Repository) (*Cache, error) {
	c := Cache{
		data: make(map[string]entity.Order),
	}

	orders, err := repo.Orders.GetAll()
	if err != nil {
		return nil, err
	}

	for _, v := range orders {
		c.data[v.ID] = v
	}

	return &c, nil
}

func (c *Cache) Create(order entity.Order) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[order.ID] = order

	return nil
}

func (c *Cache) GetById(orderId string) (entity.Order, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	val, ok := c.data[orderId]
	if !ok {
		return val, errors.New("No such id in the cache")
	}

	return val, nil
}
