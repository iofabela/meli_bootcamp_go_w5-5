package product

import (
	"context"
	"errors"

	"github.com/iofabela/meli_bootcamp_go_w5-5/internal/domain"
)

// Errors
var (
	ErrNotFound = errors.New("product not found")
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Product, error)
	Exists(ctx context.Context, productCode string) bool
	Get(ctx context.Context, id int) (domain.Product, error)
	Save(ctx context.Context, p domain.Product) (int, error)
	Update(ctx context.Context, p domain.Product) error
	Delete(ctx context.Context, id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

// Get all 'products'
func (s *service) GetAll(ctx context.Context) ([]domain.Product, error) {
	return s.repository.GetAll(ctx)
}

// Get a 'product' with id
func (s *service) Get(ctx context.Context, id int) (domain.Product, error) {
	return s.repository.Get(ctx, id)
}

// Check if a 'product' exist with 'productCode'
func (s *service) Exists(ctx context.Context, productCode string) bool {
	return s.repository.Exists(ctx, productCode)
}

// Save a 'product'
func (s *service) Save(ctx context.Context, p domain.Product) (int, error) {
	return s.repository.Save(ctx, p)
}

// Update a 'product'
func (s *service) Update(ctx context.Context, p domain.Product) error {
	return s.repository.Update(ctx, p)
}

// Delete a 'product'
func (s *service) Delete(ctx context.Context, id int) error {
	return s.repository.Delete(ctx, id)
}
