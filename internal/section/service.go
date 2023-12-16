package section

import (
	"context"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w5-5/internal/domain"
)

// Errors
var (
	ErrNotFound = errors.New("section not found")
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Section, error)
	Get(ctx context.Context, id int) (domain.Section, error)
	Save(ctx context.Context, s domain.Section) (int, error)
	Update(ctx context.Context, s domain.Section) (domain.Section, error)
	Delete(ctx context.Context, id int) error
	ReportProductsGetAll(ctx context.Context) ([]domain.ProductReport, error)
	ReportProductsGet(ctx context.Context, id int) (domain.ProductReport, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

// Trae todas las secciones registradas por el repositorio usando el metodo GetAll del repositorio
func (ser *service) GetAll(ctx context.Context) ([]domain.Section, error) {
	return ser.repository.GetAll(ctx)

}

// Busca una seccion especifica por el id dado usando el metodo Get del repositorio
func (ser *service) Get(ctx context.Context, id int) (domain.Section, error) {
	return ser.repository.Get(ctx, id)
}

// Crea una nueva seccion y llama al metodo Save, del repositorio
func (ser *service) Save(ctx context.Context, s domain.Section) (int, error) {
	if err := ser.repository.Exists(ctx, s.SectionNumber); err {
		return 0, errors.New("section_number ya existe")
	}
	section, err := ser.repository.Save(ctx, s)
	if err != nil {
		return 0, err
	}
	return section, nil

}

// Dado un id de una seccion, la encuentra y modifica sus datos con los datos dados, posteriormente, guarda los cambios usando
// el metodo Update del repositorio
func (ser *service) Update(ctx context.Context, s domain.Section) (domain.Section, error) {
	err := ser.repository.Update(ctx, s)
	if err != nil {
		return domain.Section{}, err
	}
	return s, nil
}

// Dado un id de una seccion, la encuentra lo elimina usando el metodo Delete del repositorio
func (ser *service) Delete(ctx context.Context, id int) error {
	return ser.repository.Delete(ctx, id)
}

func (ser *service) ReportProductsGetAll(ctx context.Context) ([]domain.ProductReport, error) {
	return ser.repository.ReportProductsAll(ctx)
}

func (ser *service) ReportProductsGet(ctx context.Context, id int) (domain.ProductReport, error) {
	return ser.repository.ReportProductsGet(ctx, id)
}
