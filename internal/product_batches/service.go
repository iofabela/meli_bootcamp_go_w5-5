package product_batch

import (
	"context"

	"github.com/iofabela/meli_bootcamp_go_w5-5/internal/domain"
)

type Service interface {
	Save(ctx context.Context, s domain.ProductBatches) (int, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

// Llama al metodo Save del repositorio
func (ser *service) Save(ctx context.Context, pb domain.ProductBatches) (int, error) {
	return ser.repository.Save(ctx, pb)
}
