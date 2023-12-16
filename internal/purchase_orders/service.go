package purchaseOrder

import (
	"context"

	"github.com/iofabela/meli_bootcamp_go_w5-5/internal/domain"
)

type Service interface {
	Exists(ctx context.Context, orderNumber string) (bool, error)
	Save(ctx context.Context, b domain.PurchaseOrders) (int, error)
}

type service struct {
	repository Repository
}

// NewService receive a Repository structure and return a service interface
func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

// Exists method verify if the order number is already exist and return a bool with an error.
func (s *service) Exists(ctx context.Context, orderNumber string) (bool, error) {
	return s.repository.Exists(ctx, orderNumber)
}

// Save method is used to save the new purchase order in the database and return the id with an error.
func (s *service) Save(ctx context.Context, po domain.PurchaseOrders) (int, error) {
	return s.repository.Save(ctx, po)
}
