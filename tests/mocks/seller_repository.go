package mocks

import (
	"context"
	"errors"
	"fmt"

	"github.com/iofabela/meli_bootcamp_go_w5-5/internal/domain"
)

var MockListSellers []domain.Seller = []domain.Seller{
	{
		ID:          1,
		CID:         2,
		CompanyName: "MELI",
		Address:     "Medellin",
		Telephone:   "123",
	},
	{
		ID:          2,
		CID:         3,
		CompanyName: "DIGITAL HOUSE",
		Address:     "Medellin",
		Telephone:   "1234",
	},
}

var MockNewSeller domain.Seller = domain.Seller{
	ID:          3,
	CID:         1,
	CompanyName: "MELI",
	Address:     "Medellin",
	Telephone:   "123",
}

var MockNewSellerWithConflict domain.Seller = domain.Seller{
	ID:          4,
	CID:         2,
	CompanyName: "MELI",
	Address:     "Bogota",
	Telephone:   "123",
}
var MockSeahrSeller domain.Seller = domain.Seller{
	ID:          2,
	CID:         1,
	CompanyName: "MELI",
	Address:     "Medellin",
	Telephone:   "123",
}

var MockUpdateSeller domain.Seller = domain.Seller{
	ID:          2,
	CID:         3,
	CompanyName: "MELI",
	Address:     "Argentina",
	Telephone:   "123",
}

var MockMissingCid domain.Seller = domain.Seller{
	CompanyName: "MELI",
	Address:     "Argentina",
	Telephone:   "123",
}

var MockEmptyDataSeller []domain.Seller = []domain.Seller{}

const (
	SellerNotFound = "seller with id: %d, not found"
)

type MockSellerRepo struct {
	MockSeller []domain.Seller
}

// GetAll
func (d *MockSellerRepo) GetAll(ctx context.Context) ([]domain.Seller, error) {
	return d.MockSeller, nil
}

// Get
func (d *MockSellerRepo) Get(ctx context.Context, id int) (domain.Seller, error) {
	for _, seller := range d.MockSeller {
		if seller.ID == id {
			return seller, nil
		}
	}
	return domain.Seller{}, fmt.Errorf(SellerNotFound, id)
}

func (d *MockSellerRepo) Exists(ctx context.Context, cid int) bool {
	for _, seller := range d.MockSeller {
		if seller.CID == cid {
			return true
		}
	}
	return false
}

// Save
func (d *MockSellerRepo) Save(ctx context.Context, s domain.Seller) (int, error) {
	nId := len(d.MockSeller) + 1
	s.ID = nId
	v := d.Exists(ctx, s.CID)
	if !v {
		d.MockSeller = append(d.MockSeller, s)
		return nId, nil
	}
	return 0, fmt.Errorf("el Seller con cid %d ya existe", s.CID)
}

// Update
func (d *MockSellerRepo) Update(ctx context.Context, s domain.Seller) error {
	for i, seller := range d.MockSeller {
		if seller.ID == s.ID {
			if s.CID != seller.CID && d.Exists(ctx, s.CID) {
				return errors.New("cid must be unique")
			}
			d.MockSeller[i] = s
			return nil
		}
	}
	return errors.New(SellerNotFound)
}

// Delete
func (d *MockSellerRepo) Delete(ctx context.Context, id int) error {
	for i, seller := range d.MockSeller {
		if seller.ID == id {
			if i == len(d.MockSeller)-1 {
				d.MockSeller = d.MockSeller[:i]
			} else {
				d.MockSeller = append(d.MockSeller[:i], d.MockSeller...)
			}
			return nil
		}
	}
	return errors.New(SellerNotFound)
}

func (d *MockSellerRepo) CIDExist(ctx context.Context, cid int) bool {
	return false
}
