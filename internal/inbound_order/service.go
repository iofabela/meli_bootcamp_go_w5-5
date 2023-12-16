package inboundorder

import (
	"context"

	"github.com/iofabela/meli_bootcamp_go_w5-5/internal/domain"
)

// Service interface for handling requests
type Service interface {
	Save(ctx context.Context, i domain.InboundOrder) (int, error)
	Exists(ctx context.Context, employeeID int) (bool, error)
}

type service struct {
	repository Repository
}

// NewService creates a new service
func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) Save(ctx context.Context, i domain.InboundOrder) (int, error) {
	return s.repository.Save(ctx, i)
}

func (s *service) Exists(ctx context.Context, employeeID int) (bool, error) {
	return s.repository.Exists(ctx, employeeID)
}
