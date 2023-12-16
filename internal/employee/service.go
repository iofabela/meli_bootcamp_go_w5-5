package employee

import (
	"context"
	"errors"

	"github.com/iofabela/meli_bootcamp_go_w5-5/internal/domain"
)

// Errors
var (
	ErrNotFound = errors.New("employee not found")
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Employee, error)
	Get(ctx context.Context, id int) (domain.Employee, error)
	Exists(ctx context.Context, cardNumberID string) bool
	Save(ctx context.Context, e domain.Employee) (int, error)
	Update(ctx context.Context, e domain.Employee) error
	Delete(ctx context.Context, id int) error
	GetInboundOrders(ctx context.Context, id int) ([]domain.EmployeeOrders, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) GetAll(ctx context.Context) ([]domain.Employee, error) {
	return s.repository.GetAll(ctx)
}

func (s *service) Get(ctx context.Context, id int) (domain.Employee, error) {
	return s.repository.Get(ctx, id)
}

func (s *service) Exists(ctx context.Context, cardNumberID string) bool {
	return s.repository.Exists(ctx, cardNumberID)
}

func (s *service) Save(ctx context.Context, e domain.Employee) (int, error) {
	return s.repository.Save(ctx, e)
}

func (s *service) Update(ctx context.Context, e domain.Employee) error {
	return s.repository.Update(ctx, e)
}

func (s *service) Delete(ctx context.Context, id int) error {
	return s.repository.Delete(ctx, id)
}

func (s *service) GetInboundOrders(ctx context.Context, id int) ([]domain.EmployeeOrders, error) {
	return s.repository.GetInboundOrders(ctx, id)
}
