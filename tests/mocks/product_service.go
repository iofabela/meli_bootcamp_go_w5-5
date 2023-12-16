package mocks

import (
	"context"

	"github.com/extmatperez/meli_bootcamp_go_w5-5/internal/domain"
)

// ServiceTestProduct ...
type ServiceTestProduct interface {
	GetAll(ctx context.Context) ([]domain.Product, error)
	Exists(ctx context.Context, productCode string) bool
	Get(ctx context.Context, id int) (domain.Product, error)
	Save(ctx context.Context, p domain.Product) (int, error)
	Update(ctx context.Context, p domain.Product) error
	Delete(ctx context.Context, id int) error
}

// MockServiceProduct ...
type MockServiceProduct struct {
	MockProductRepository MockRepositoryProduct
}

// GetAll ... | Get all 'products'
func (s *MockServiceProduct) GetAll(ctx context.Context) ([]domain.Product, error) {
	return s.MockProductRepository.GetAll(ctx)
}

// Get ... | Get a 'product' with id
func (s *MockServiceProduct) Get(ctx context.Context, id int) (domain.Product, error) {
	return s.MockProductRepository.Get(ctx, id)
}

// Exists ... | Check if a 'product' exist with 'productCode'
func (s *MockServiceProduct) Exists(ctx context.Context, productCode string) bool {
	return s.MockProductRepository.Exists(ctx, productCode)
}

// Save ... | Save a 'product'
func (s *MockServiceProduct) Save(ctx context.Context, p domain.Product) (int, error) {
	return s.MockProductRepository.Save(ctx, p)
}

// Update ... | Update a 'product'
func (s *MockServiceProduct) Update(ctx context.Context, p domain.Product) error {
	return s.MockProductRepository.Update(ctx, p)
}

// Delete ... | Delete a 'product'
func (s *MockServiceProduct) Delete(ctx context.Context, id int) error {
	return s.MockProductRepository.Delete(ctx, id)
}
