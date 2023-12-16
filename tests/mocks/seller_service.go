package mocks

import (
	"context"

	"github.com/iofabela/meli_bootcamp_go_w5-5/internal/domain"
)

type MockServiceSeller struct {
	MockRepo MockSellerRepo
}

func (m *MockServiceSeller) GetAll(ctx context.Context) ([]domain.Seller, error) {
	return m.MockRepo.GetAll(ctx)
}

func (m *MockServiceSeller) Get(ctx context.Context, id int) (domain.Seller, error) {
	return m.MockRepo.Get(ctx, id)
}

func (m *MockServiceSeller) Exists(ctx context.Context, cid int) bool {
	return m.MockRepo.Exists(ctx, cid)
}

func (m *MockServiceSeller) Save(ctx context.Context, s domain.Seller) (int, error) {
	return m.MockRepo.Save(ctx, s)
}

func (m *MockServiceSeller) Update(ctx context.Context, s domain.Seller) error {
	return m.MockRepo.Update(ctx, s)
}

func (m *MockServiceSeller) Delete(ctx context.Context, id int) error {
	return m.MockRepo.Delete(ctx, id)
}

func (m *MockServiceSeller) CIDExist(ctx context.Context, cid int) bool {
	return m.MockRepo.CIDExist(ctx, cid)
}
