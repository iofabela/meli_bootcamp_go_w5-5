package locality

import (
	"context"

	"github.com/iofabela/meli_bootcamp_go_w5-5/internal/domain"
)

type Service interface {
	SaveLocality(ctx context.Context, l domain.Locality) (int, error)
	IDExist(ctx context.Context, id int) bool
	SellerReport(ctx context.Context, id int) (domain.ReportSeller, error)
	GetAllSellerReports(ctx context.Context) ([]domain.ReportSeller, error)
	GetCarryReport(ctx context.Context, id int) (domain.LocalityCarries, error)
	GetAllCarryReports(ctx context.Context) ([]domain.LocalityCarries, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

// Save Locality
func (s *service) SaveLocality(ctx context.Context, l domain.Locality) (int, error) {
	return s.repository.SaveLocality(ctx, l)
}

// Validate Id
func (s *service) IDExist(ctx context.Context, id int) bool {
	return s.repository.IDExist(ctx, id)
}

// Report Seller For Id
func (s *service) SellerReport(ctx context.Context, id int) (domain.ReportSeller, error) {
	return s.repository.SellerReport(ctx, id)
}

// Report Seller All
func (s *service) GetAllSellerReports(ctx context.Context) ([]domain.ReportSeller, error) {
	return s.repository.GetAllSellerReports(ctx)
}

func (s *service) GetCarryReport(ctx context.Context, id int) (domain.LocalityCarries, error) {
	return s.repository.GetCarryReport(ctx, id)
}

func (s *service) GetAllCarryReports(ctx context.Context) ([]domain.LocalityCarries, error) {
	return s.repository.GetAllCarryReports(ctx)
}
