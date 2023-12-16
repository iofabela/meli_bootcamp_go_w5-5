package mocks

import (
	"context"
	"errors"

	"github.com/iofabela/meli_bootcamp_go_w5-5/internal/domain"
)

type MockWarehouseRepository struct {
	MockData []domain.Warehouse
}

func (s *MockWarehouseRepository) GetAll(ctx context.Context) ([]domain.Warehouse, error) {
	return s.MockData, nil
}

func (s *MockWarehouseRepository) Get(ctx context.Context, id int) (domain.Warehouse, error) {
	for _, testWH := range s.MockData {
		if testWH.ID == id {
			return testWH, nil
		}
	}
	return domain.Warehouse{}, errors.New("warehouse not found")
}

func (s *MockWarehouseRepository) Exists(ctx context.Context, warehouseCode string) bool {
	for _, testWH := range s.MockData {
		if testWH.WarehouseCode == warehouseCode {
			return true
		}
	}
	return false
}

func (s *MockWarehouseRepository) Save(ctx context.Context, w domain.Warehouse) (int, error) {
	if w.WarehouseCode == "" {
		return 0, errors.New("warehouseCode is required")
	}
	if s.Exists(ctx, w.WarehouseCode) {
		return 0, errors.New("warehouseCode must be unique")
	}
	w.ID = s.MockData[len(s.MockData)-1].ID + 1
	s.MockData = append(s.MockData, w)
	return w.ID, nil
}

func (s *MockWarehouseRepository) Update(ctx context.Context, w domain.Warehouse) error {
	for i, testWH := range s.MockData {
		if testWH.ID == w.ID {
			if w.WarehouseCode != testWH.WarehouseCode && s.Exists(ctx, w.WarehouseCode) {
				return errors.New("warehouseCode must be unique")
			}
			s.MockData[i] = w
			return nil
		}
	}
	return errors.New("warehouse not found")
}

func (s *MockWarehouseRepository) Delete(ctx context.Context, id int) error {
	for i, testWH := range s.MockData {
		if testWH.ID == id {
			s.MockData = append(s.MockData[:i], s.MockData[i+1:]...)
			return nil
		}
	}
	return errors.New("warehouse not found")
}

var MockDataWarehouse []domain.Warehouse = []domain.Warehouse{
	{
		ID:                 1,
		Address:            "Calle Falsa 123",
		Telephone:          "+54 9 11 5487-5421",
		WarehouseCode:      "CTX-555",
		MinimumCapacity:    5,
		MinimumTemperature: 7,
	},
	{
		ID:                 2,
		Address:            "Polo DOT",
		Telephone:          "0800-555-4455",
		WarehouseCode:      "CTX-576",
		MinimumCapacity:    10,
		MinimumTemperature: -5,
	},
}

var MockEmptyDataWarehouse []domain.Warehouse = []domain.Warehouse{}
