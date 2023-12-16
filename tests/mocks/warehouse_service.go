package mocks

import (
	"context"
	"errors"

	"github.com/iofabela/meli_bootcamp_go_w5-5/internal/domain"
)

type MockWarehouseService struct {
	MockRepository MockWarehouseRepository
}

func (s *MockWarehouseService) GetAll(ctx context.Context) ([]domain.Warehouse, error) {
	return s.MockRepository.GetAll(ctx)
}

func (s *MockWarehouseService) Get(ctx context.Context, id int) (domain.Warehouse, error) {
	return s.MockRepository.Get(ctx, id)
}

func (s *MockWarehouseService) Exists(ctx context.Context, warehouseCode string) bool {
	return s.MockRepository.Exists(ctx, warehouseCode)
}

func (s *MockWarehouseService) Save(ctx context.Context, w domain.Warehouse) (int, error) {
	return s.MockRepository.Save(ctx, w)
}

func (s *MockWarehouseService) Update(ctx context.Context, w domain.Warehouse) error {
	return s.MockRepository.Update(ctx, w)
}

func (s *MockWarehouseService) Delete(ctx context.Context, id int) error {
	return s.MockRepository.Delete(ctx, id)
}

type MockWarehouseServiceError struct {
	MockRepository MockWarehouseRepository
}

func (s *MockWarehouseServiceError) GetAll(ctx context.Context) ([]domain.Warehouse, error) {
	return []domain.Warehouse{}, errors.New("communication error with the database")
}

func (s *MockWarehouseServiceError) Get(ctx context.Context, id int) (domain.Warehouse, error) {
	return s.MockRepository.Get(ctx, id)
}

func (s *MockWarehouseServiceError) Exists(ctx context.Context, warehouseCode string) bool {
	return s.MockRepository.Exists(ctx, warehouseCode)
}

func (s *MockWarehouseServiceError) Save(ctx context.Context, w domain.Warehouse) (int, error) {
	return 0, errors.New("communication error with the database")
}

func (s *MockWarehouseServiceError) Update(ctx context.Context, w domain.Warehouse) error {
	return errors.New("communication error with the database")
}

func (s *MockWarehouseServiceError) Delete(ctx context.Context, id int) error {
	return errors.New("communication error with the database")
}
