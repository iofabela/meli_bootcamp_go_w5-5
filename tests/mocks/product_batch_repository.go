package mocks

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/extmatperez/meli_bootcamp_go_w5-5/internal/domain"
	"github.com/stretchr/testify/assert"
)

func CreateMockDBSave(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()

	// añadir mocks de product a db
	expectedProduct := mock.ExpectQuery("SELECT id FROM products")

	productMockRows := sqlmock.NewRows([]string{"id"})
	productMockRows.AddRow(1)
	expectedProduct.WithArgs(1).WillReturnRows(productMockRows)

	// añadir mocks de section a db
	sectionRows := []string{
		"id",
	}
	sectionMockRows := sqlmock.NewRows(sectionRows)
	sectionMockRows.AddRow(1)
	mock.ExpectQuery("SELECT id FROM sections").WithArgs(MockProductBatch.SectionId).WillReturnRows(sectionMockRows)

	// añadir mock INSERT product_batches
	prep := mock.ExpectPrepare("^INSERT INTO product_batches*")
	productBatch := MockProductBatch
	prep.ExpectExec().WithArgs(
		productBatch.BatchNumber,
		productBatch.CurrentQuantity,
		productBatch.CurrentTemperature,
		productBatch.DueDate,
		productBatch.InitialQuantity,
		productBatch.ManufacturingDate,
		productBatch.ManufacturingHour,
		productBatch.MinumumTemperature,
		productBatch.ProductId,
		productBatch.SectionId,
	).WillReturnResult(sqlmock.NewResult(int64(productBatch.Id), 1))

	assert.NoError(t, err)
	return db, mock
}

func CreateMockDBSaveNonExistantProduct(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()

	// añadir mocks de product a db
	expectedProduct := mock.ExpectQuery("SELECT id FROM products")

	productMockRows := sqlmock.NewRows([]string{"id"})
	expectedProduct.WithArgs(2).WillReturnRows(productMockRows)

	assert.NoError(t, err)
	return db, mock
}

func CreateMockDBSaveNonExistantSection(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()

	// añadir mocks de product a db
	expectedProduct := mock.ExpectQuery("SELECT id FROM products")

	productMockRows := sqlmock.NewRows([]string{"id"})
	productMockRows.AddRow(1)
	expectedProduct.WithArgs(1).WillReturnRows(productMockRows)

	expectedSection := mock.ExpectQuery("SELECT id FROM sections")

	sectionMockRows := sqlmock.NewRows([]string{"id"})
	expectedSection.WillReturnRows(sectionMockRows)

	assert.NoError(t, err)
	return db, mock
}

var MockProductBatch = domain.ProductBatches{
	Id:                 1,
	BatchNumber:        10,
	CurrentQuantity:    10,
	CurrentTemperature: 20,
	DueDate:            "2022-08-21",
	InitialQuantity:    1,
	ManufacturingDate:  "2021-12-10",
	ManufacturingHour:  "12:00",
	MinumumTemperature: 10,
	ProductId:          1,
	SectionId:          1,
}

var MockProductBatchNonExistantProduct = domain.ProductBatches{
	Id:                 1,
	BatchNumber:        10,
	CurrentQuantity:    10,
	CurrentTemperature: 20,
	DueDate:            "2022-08-21",
	InitialQuantity:    1,
	ManufacturingDate:  "2021-12-10",
	ManufacturingHour:  "12:00",
	MinumumTemperature: 10,
	ProductId:          2,
	SectionId:          1,
}

var MockProductBatchNonExistantSection = domain.ProductBatches{
	Id:                 1,
	BatchNumber:        10,
	CurrentQuantity:    10,
	CurrentTemperature: 20,
	DueDate:            "2022-08-21",
	InitialQuantity:    1,
	ManufacturingDate:  "2021-12-10",
	ManufacturingHour:  "12:00",
	MinumumTemperature: 10,
	ProductId:          1,
	SectionId:          2,
}
