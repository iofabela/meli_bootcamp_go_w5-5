package mocks

import (
	"context"
	"fmt"

	"github.com/iofabela/meli_bootcamp_go_w5-5/internal/domain"
)

// MockListProducts ...
var MockListProducts []domain.Product = []domain.Product{
	{
		ID:             1,
		Description:    "test",
		ExpirationRate: 1,
		FreezingRate:   2,
		Height:         6.4,
		Length:         4.5,
		Netweight:      3.4,
		ProductCode:    "ssd",
		RecomFreezTemp: 1.3,
		Width:          1.2,
		ProductTypeID:  2,
		SellerID:       2,
	}, {
		ID:             2,
		Description:    "test #1",
		ExpirationRate: 2,
		FreezingRate:   3,
		Height:         3.2,
		Length:         3.2,
		Netweight:      2.1,
		ProductCode:    "raspberry",
		RecomFreezTemp: 1.2,
		Width:          2.8,
		ProductTypeID:  3,
		SellerID:       3,
	},
}

// MockUpdateProduct ...
var MockUpdateProduct domain.Product = domain.Product{
	ID:             2,
	Description:    "test #1",
	ExpirationRate: 1,
	FreezingRate:   5,
	Height:         6.2,
	Length:         4.2,
	Netweight:      8.1,
	ProductCode:    "arduino",
	RecomFreezTemp: 1.2,
	Width:          4.4,
	ProductTypeID:  3,
	SellerID:       3,
}

// MockUpdateConflictProduct ...
var MockUpdateConflictProduct domain.Product = domain.Product{
	Description:    "",
	ExpirationRate: 0,
	FreezingRate:   0,
	Height:         0,
	Length:         0,
	Netweight:      0,
	ProductCode:    "arduino",
	RecomFreezTemp: 0,
	Width:          0,
	ProductTypeID:  1,
	SellerID:       0,
}

// MockCreateProduct ...
var MockCreateProduct domain.Product = domain.Product{
	Description:    "Create #1",
	ExpirationRate: 6,
	FreezingRate:   3,
	Height:         3.2,
	Length:         9.2,
	Netweight:      10.1,
	ProductCode:    "screen",
	RecomFreezTemp: 1.2,
	Width:          8.3,
	ProductTypeID:  3,
	SellerID:       5,
}

// MockProductNonExistent ...
var MockProductNonExistent domain.Product = domain.Product{
	Description:    "",
	ExpirationRate: 0,
	FreezingRate:   0,
	Height:         0,
	Length:         0,
	Netweight:      0,
	ProductCode:    "",
	RecomFreezTemp: 0,
	Width:          0,
	ProductTypeID:  0,
	SellerID:       0,
}

// MockDataEmptyProduct ...
var MockDataEmptyProduct []domain.Product = []domain.Product{}

// CONSTANTS
const ( //	ERRORS	 = messages
	ProductNotFound  = "product with id: %d, not found"
	FailReading      = "cant read database"
	FailWriting      = "cant write database, error: %w"
	ProductCodeError = "a product code already exists with: %d"
)

// MockRepositoryProduct ...
type MockRepositoryProduct struct {
	Data []domain.Product
}

// Delete ...
func (r *MockRepositoryProduct) Delete(ctx context.Context, id int) error {
	for i, product := range r.Data {
		if product.ID == id {
			r.Data = append(r.Data[:i], r.Data[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf(ProductNotFound, id)
}

// Exists ...
func (r *MockRepositoryProduct) Exists(ctx context.Context, productCode string) bool {
	for _, product := range r.Data {
		if product.ProductCode == productCode {
			return true
		}
	}
	return false
}

// Get ...
func (r *MockRepositoryProduct) Get(ctx context.Context, id int) (domain.Product, error) {
	for _, product := range r.Data {
		if product.ID == id {
			return product, nil
		}
	}
	return domain.Product{}, fmt.Errorf(ProductNotFound, id)
}

// GetAll ...
func (r *MockRepositoryProduct) GetAll(ctx context.Context) ([]domain.Product, error) {
	return r.Data, nil
}

// Save ...
func (r *MockRepositoryProduct) Save(ctx context.Context, p domain.Product) (int, error) {
	id := len(r.Data) + 1
	p.ID = id
	check := r.Exists(ctx, p.ProductCode)
	if !check {
		r.Data = append(r.Data, p)
		return id, nil
	}
	return 0, fmt.Errorf(ProductCodeError, id)
}

// Update ...
func (r *MockRepositoryProduct) Update(ctx context.Context, p domain.Product) error {
	for _, product := range r.Data {
		if product.ID == p.ID {
			product = p
			return nil
		}
	}
	return fmt.Errorf(ProductNotFound, p.ID)
}
