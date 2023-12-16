package mocks

import (
	"context"

	"github.com/extmatperez/meli_bootcamp_go_w5-5/internal/domain"
)

type MockEmployeeService struct {
	MockRepository MockEmployeeRepository
}

func (r *MockEmployeeService) GetAll(ctx context.Context) ([]domain.Employee, error) {
	return r.MockRepository.GetAll(ctx)
}

func (r *MockEmployeeService) Get(ctx context.Context, id int) (domain.Employee, error) {
	return r.MockRepository.Get(ctx, id)
}

func (r *MockEmployeeService) Exists(ctx context.Context, CardNumberID string) bool {
	return r.MockRepository.Exists(ctx, CardNumberID)
}

func (r *MockEmployeeService) Save(ctx context.Context, e domain.Employee) (int, error) {
	return r.MockRepository.Save(ctx, e)
}

func (r *MockEmployeeService) Update(ctx context.Context, e domain.Employee) error {
	return r.MockRepository.Update(ctx, e)
}

func (r *MockEmployeeService) Delete(ctx context.Context, id int) error {
	return r.MockRepository.Delete(ctx, id)
}

func (r *MockEmployeeService) GetInboundOrders(ctx context.Context, id int) ([]domain.EmployeeOrders, error) {
	return nil, nil
}
