package carry

import (
	"context"

	"github.com/extmatperez/meli_bootcamp_go_w5-5/internal/domain"
)

type Service interface {
	Save(ctx context.Context, w domain.Carry) (int, error)
	CIDExists(ctx context.Context, cid string) (bool, error)
	LocalityExists(ctx context.Context, id int) (bool, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) Save(ctx context.Context, c domain.Carry) (int, error) {
	return s.repository.Save(ctx, c)
}

func (s *service) CIDExists(ctx context.Context, cid string) (bool, error) {
	return s.repository.CIDExists(ctx, cid)
}

func (s *service) LocalityExists(ctx context.Context, id int) (bool, error) {
	return s.repository.LocalityExists(ctx, id)
}
