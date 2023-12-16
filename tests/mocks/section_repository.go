package mocks

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/extmatperez/meli_bootcamp_go_w5-5/internal/domain"
	"github.com/stretchr/testify/assert"
)

type MockSectionRepository struct {
	MockData []domain.Section
}

func (mk *MockSectionRepository) GetAll(ctx context.Context) ([]domain.Section, error) {
	return mk.MockData, nil
}

func (mk *MockSectionRepository) Get(ctx context.Context, id int) (domain.Section, error) {
	for _, section := range mk.MockData {
		if section.ID == id {
			return section, nil
		}
	}
	return domain.Section{}, fmt.Errorf("no se encontro seccion: %d", id)
}

func (mk *MockSectionRepository) Exists(ctx context.Context, cid int) bool {
	for _, section := range mk.MockData {
		if section.ID == cid {
			return true
		}
	}
	return false
}

func (mk *MockSectionRepository) Save(ctx context.Context, s domain.Section) (int, error) {
	newId := len(mk.MockData) + 1
	s.ID = newId
	mk.MockData = append(mk.MockData, s)
	return newId, nil
}

func (mk *MockSectionRepository) Update(ctx context.Context, s domain.Section) error {
	for _, section := range mk.MockData {
		if section.ID == s.ID {
			section = s
			return nil
		}
	}
	return fmt.Errorf("no se encontro seccion: %d", s.ID)
}

func (mk *MockSectionRepository) Delete(ctx context.Context, id int) error {
	for i, section := range mk.MockData {
		if section.ID == id {
			mk.MockData = append(mk.MockData[:i], mk.MockData[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("no se encontro seccion: %d", id)
}

func (mk *MockSectionRepository) ReportProductsAll(ctx context.Context) ([]domain.ProductReport, error) {
	return []domain.ProductReport{}, nil
}

func (mk *MockSectionRepository) ReportProductsGet(ctx context.Context, id int) (domain.ProductReport, error) {
	return domain.ProductReport{}, nil
}

type MockSectionErrorRepository struct {
	MockData []domain.Section
}

func (mk *MockSectionErrorRepository) GetAll(ctx context.Context) ([]domain.Section, error) {
	return []domain.Section{}, fmt.Errorf("no se pueden obtener las secciones")
}

func (mk *MockSectionErrorRepository) Get(ctx context.Context, id int) (domain.Section, error) {
	for _, section := range mk.MockData {
		if section.ID == id {
			return section, nil
		}
	}
	return domain.Section{}, fmt.Errorf("no se encontro seccion: %d", id)
}

func (mk *MockSectionErrorRepository) Exists(ctx context.Context, cid int) bool {
	return false
}

func (mk *MockSectionErrorRepository) Save(ctx context.Context, s domain.Section) (int, error) {
	return 0, fmt.Errorf("no se pudo crear seccion")
}

func (mk *MockSectionErrorRepository) Update(ctx context.Context, s domain.Section) error {
	return fmt.Errorf("no se actualizo seccion: %d", s.ID)
}

func (mk *MockSectionErrorRepository) Delete(ctx context.Context, id int) error {
	return fmt.Errorf("no se encontro elimino: %d", id)
}

func (mk *MockSectionErrorRepository) ReportProductsAll(ctx context.Context) ([]domain.ProductReport, error) {
	return []domain.ProductReport{}, fmt.Errorf("cant do report")
}

func (mk *MockSectionErrorRepository) ReportProductsGet(ctx context.Context, id int) (domain.ProductReport, error) {
	return domain.ProductReport{}, fmt.Errorf("NuevoError")
}

func CreateMockDBProductsGetAll(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()

	// añadir mocks de product a db
	query := "SELECT .* FROM sections s *"
	expectedProduct := mock.ExpectQuery(query)

	productMockRows := sqlmock.NewRows([]string{"id", "section_number", "product_count"})
	productMockRows.AddRow(1, 250, 100)
	productMockRows.AddRow(2, 120, 10)
	expectedProduct.WillReturnRows(productMockRows)

	assert.NoError(t, err)
	return db, mock
}

func CreateMockDBProductsGet(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()

	// añadir mocks de product a db
	query := "SELECT .* FROM sections s .* WHERE s.id = ? .*"
	expectedProduct := mock.ExpectPrepare(query).ExpectQuery().WithArgs(1)

	productMockRows := sqlmock.NewRows([]string{"id", "section_number", "product_count"})
	productMockRows.AddRow(1, 250, 100)
	expectedProduct.WillReturnRows(productMockRows)

	assert.NoError(t, err)
	return db, mock
}

var MockListaSections []domain.Section = []domain.Section{
	{
		ID:                 1,
		SectionNumber:      1,
		CurrentTemperature: 10,
		MinimumTemperature: 5,
		CurrentCapacity:    100,
		MinimumCapacity:    10,
		MaximumCapacity:    1000,
		WarehouseID:        10,
		ProductTypeID:      1,
	},
	MockBuscarSection,
}

var MockNuevaSection domain.Section = domain.Section{
	ID:                 0,
	SectionNumber:      10,
	CurrentTemperature: 40,
	MinimumTemperature: 30,
	CurrentCapacity:    20,
	MinimumCapacity:    0,
	MaximumCapacity:    100,
	WarehouseID:        5,
	ProductTypeID:      5,
}

var MockNuevaSectionConflicto domain.Section = domain.Section{
	ID:                 10,
	SectionNumber:      1,
	CurrentTemperature: 40,
	MinimumTemperature: 30,
	CurrentCapacity:    20,
	MinimumCapacity:    0,
	MaximumCapacity:    100,
	WarehouseID:        5,
	ProductTypeID:      5,
}

var MockBuscarSection domain.Section = domain.Section{
	ID:                 2,
	SectionNumber:      2,
	CurrentTemperature: 20,
	MinimumTemperature: 0,
	CurrentCapacity:    10,
	MinimumCapacity:    1,
	MaximumCapacity:    100,
	WarehouseID:        2,
	ProductTypeID:      2,
}

var MockActualizarSection domain.Section = domain.Section{
	ID:                 2,
	SectionNumber:      2,
	CurrentTemperature: 15,
	MinimumTemperature: 10,
	CurrentCapacity:    5,
	MinimumCapacity:    1,
	MaximumCapacity:    100,
	WarehouseID:        2,
	ProductTypeID:      2,
}

var MockGetAllProductReport []domain.ProductReport = []domain.ProductReport{
	{
		SectionId:     1,
		SectionNumber: 250,
		ProductCount:  100,
	},
	{
		SectionId:     2,
		SectionNumber: 120,
		ProductCount:  10,
	},
}
