package buyer

import (
	"context"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w5-5/internal/domain"
)

// Errors
var (
	ErrNotFound = errors.New("buyer not found")
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Buyer, error)
	Get(ctx context.Context, id int) (domain.Buyer, error)
	Exists(ctx context.Context, cardNumberID string) bool
	Save(ctx context.Context, b domain.Buyer) (int, error)
	Update(ctx context.Context, b domain.Buyer) error
	Delete(ctx context.Context, id int) error
	GetPurchaseOrders(ctx context.Context, id int) ([]domain.BuyerOrders, error)
}

type service struct {
	repository Repository
}

//NewService receive a Repository structure and return a service interface
func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

//GetAll receive the context, generate a instance of repository.GetAll and return a list of buyer to repository and error
func (s *service) GetAll(ctx context.Context) ([]domain.Buyer, error) {
	return s.repository.GetAll(ctx)
}

//Get receive the context and the buyer id from the handler, generate a instance of repository.Get and return the buyer and error
func (s *service) Get(ctx context.Context, id int) (domain.Buyer, error) {
	return s.repository.Get(ctx, id)
}

//Exists receive the context and the buyer cardNumberID from the handler, generate a instance of repository.Exists and return a boolean
func (s *service) Exists(ctx context.Context, cardNumberID string) bool {
	return s.repository.Exists(ctx, cardNumberID)
}

//Save receive the context and the buyer to save, generate a instance of repository.Save and return the id and error
func (s *service) Save(ctx context.Context, b domain.Buyer) (int, error) {

	return s.repository.Save(ctx, b)
}

//Update receive the context and the buyer to update, generate a instance of repository.Update and return error
func (s *service) Update(ctx context.Context, b domain.Buyer) error {
	return s.repository.Update(ctx, b)
}

//Delete receive the context and the buyer id to delete, generate a instance of repository.Delete and return error
func (s *service) Delete(ctx context.Context, id int) error {
	return s.repository.Delete(ctx, id)
}

//GetPurchaseOrders is used to obain all buyers and the number of its purchases. The function return a slice of buyer with an error.
func (s *service) GetPurchaseOrders(ctx context.Context, id int) ([]domain.BuyerOrders, error) {
	return s.repository.GetPurchaseOrders(ctx, id)
}
