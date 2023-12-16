package section

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/iofabela/meli_bootcamp_go_w5-5/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func createMockRepository() Service {
	m := mocks.MockSectionRepository{
		MockData: mocks.MockListaSections,
	}
	return NewService(&m)
}

func createMockErrorRepository() Service {
	m := mocks.MockSectionErrorRepository{
		MockData: mocks.MockListaSections,
	}
	return NewService(&m)
}

func TestCreateOkService(t *testing.T) {
	s := createMockRepository()
	ctx := context.TODO()
	newId, err := s.Save(ctx, mocks.MockNuevaSection)
	assert.NoError(t, err)
	assert.NotEmpty(t, newId)
}

func TestCreateConflictService(t *testing.T) {
	expectedErrStr := "section_number ya existe"
	m := mocks.MockSectionRepository{
		MockData: mocks.MockListaSections,
	}
	ctx := context.TODO()
	s := NewService(&m)
	newId, err := s.Save(ctx, mocks.MockNuevaSectionConflicto)
	assert.EqualError(t, err, expectedErrStr)
	assert.Empty(t, newId)
}

func TestCreateErrorService(t *testing.T) {
	ctx := context.TODO()
	s := createMockErrorRepository()
	newId, err := s.Save(ctx, mocks.MockNuevaSectionConflicto)
	assert.Error(t, err)
	assert.Empty(t, newId)
}

func TestFindAllService(t *testing.T) {
	s := createMockRepository()
	ctx := context.TODO()
	_, err := s.GetAll(ctx)
	assert.NoError(t, err)
}

func TestFindByIdNonExistantService(t *testing.T) {
	expectedErrFormat := "no se encontro seccion: %d"
	searchId := 100
	expectedErrStr := fmt.Sprintf(expectedErrFormat, searchId)
	s := createMockRepository()
	ctx := context.TODO()
	_, err := s.Get(ctx, searchId)
	assert.EqualError(t, err, expectedErrStr)
}

func TestFindByIdService(t *testing.T) {
	searchId := 1
	s := createMockRepository()
	ctx := context.TODO()
	res, err := s.Get(ctx, searchId)
	assert.NoError(t, err)
	assert.NotEmpty(t, res.ID)
}

func TestUpdateExistantService(t *testing.T) {
	s := createMockRepository()
	ctx := context.TODO()
	updatedSection, err := s.Update(ctx, mocks.MockActualizarSection)
	assert.NoError(t, err)
	assert.NotEmpty(t, updatedSection.ID)
}

func TestUpdateNonExistantService(t *testing.T) {
	expectedErrFormat := "no se encontro seccion: %d"
	expectedErrStr := fmt.Sprintf(expectedErrFormat, mocks.MockNuevaSectionConflicto.ID)
	s := createMockRepository()
	ctx := context.TODO()
	updatedSection, err := s.Update(ctx, mocks.MockNuevaSectionConflicto)
	assert.Error(t, err)
	assert.Empty(t, updatedSection.ID)
	assert.EqualError(t, err, expectedErrStr)
}

func TestDeleteNonExistantService(t *testing.T) {
	expectedErrFormat := "no se encontro seccion: %d"
	deletedId := 10
	expectedErrStr := fmt.Sprintf(expectedErrFormat, deletedId)
	s := createMockRepository()
	ctx := context.TODO()
	err := s.Delete(ctx, deletedId)
	assert.EqualError(t, err, expectedErrStr)
}

func TestDeleteService(t *testing.T) {
	deletedId := 1
	s := createMockRepository()
	ctx := context.TODO()
	err := s.Delete(ctx, deletedId)
	deletedSection, errBusqueda := s.Get(ctx, deletedId)
	assert.NoError(t, err)
	assert.Error(t, errBusqueda)
	assert.Empty(t, deletedSection.ID)
}

func TestReportProductsAllRepository(t *testing.T) {
	mockDB, mock := mocks.CreateMockDBProductsGetAll(t)
	repo := NewRepository(mockDB)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	res, err := repo.ReportProductsAll(ctx)
	assert.NoError(t, err)
	assert.Equal(t, mocks.MockGetAllProductReport, res)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestReportProductsGetRepository(t *testing.T) {
	mockDB, mock := mocks.CreateMockDBProductsGet(t)
	repo := NewRepository(mockDB)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	res, err := repo.ReportProductsGet(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, mocks.MockGetAllProductReport[0], res)
	assert.NoError(t, mock.ExpectationsWereMet())
}
