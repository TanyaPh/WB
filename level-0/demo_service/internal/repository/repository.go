package repository

import (
	"api/internal/entity"
	"api/internal/repository/postgres"
	pg "api/pkg/postgres"
)

type Order interface {
	Create(order entity.Order) error
	GetById(orderId string) (entity.Order, error)
	GetAll() ([]entity.Order, error)
}

type Repository struct {
	Orders Order
}

func NewRepository(db *pg.Postgres) *Repository {
	return &Repository{
		Orders: postgres.NewOrderPostgres(db),
	}
}
