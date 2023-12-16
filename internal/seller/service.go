package seller

import (
	"context"
	"errors"

	"github.com/iofabela/meli_bootcamp_go_w5-5/internal/domain"
)

// Errors
var (
	ErrNotFound = errors.New("seller not found")
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Seller, error)
	Get(ctx context.Context, id int) (domain.Seller, error)
	Exists(ctx context.Context, cid int) bool
	Save(ctx context.Context, s domain.Seller) (int, error)
	Update(ctx context.Context, s domain.Seller) error
	Delete(ctx context.Context, id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

// La funcion Extraer todos los sellers existentes
func (s *service) GetAll(ctx context.Context) ([]domain.Seller, error) {
	return s.repository.GetAll(ctx)
}

// La funcion permite Extrae un seller especifico segun su id
func (s *service) Get(ctx context.Context, id int) (domain.Seller, error) {
	return s.repository.Get(ctx, id)
}

// La funcion Valida si existe un seller con cierto cid
// Con el fin de que sean unicos
func (s *service) Exists(ctx context.Context, cid int) bool {
	return s.repository.Exists(ctx, cid)
}

// La funcion Guarda un nuevo seller
func (se *service) Save(ctx context.Context, s domain.Seller) (int, error) {
	return se.repository.Save(ctx, s)
}

// La funcion Actualiza un seller segun su id
// Se le pasan los datos a actualizar, valida
// que cumpla los requerimientos y actualiza
func (se *service) Update(ctx context.Context, s domain.Seller) error {
	return se.repository.Update(ctx, s)
}

// La funcion Elimina un seller segun id
// Se le pasa el id y elimina la informacion asociada al id
func (se *service) Delete(ctx context.Context, id int) error {
	return se.repository.Delete(ctx, id)
}
