package employee

import (
	"context"
	"testing"

	"github.com/iofabela/meli_bootcamp_go_w5-5/internal/domain"
	"github.com/iofabela/meli_bootcamp_go_w5-5/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func createService() Service {
	m := &mocks.MockEmployeeRepository{
		MockData: mocks.MockEmployees,
	}
	return NewService(m)
}

func TestCreateOkEmployee(t *testing.T) {

	s := createService()

	mocks.MockNewEmployee.ID = 4

	ctx := context.TODO()
	res, err := s.Save(ctx, mocks.MockNewEmployee)

	assert.Nil(t, err)
	assert.Equal(t, mocks.MockNewEmployee.ID, res)
	assert.NotEqual(t, 0, res)
}

func TestCreateWithConflictEmployee(t *testing.T) {

	s := createService()

	expectedErr := "El empleado ya existe"

	ctx := context.TODO()
	res, err := s.Save(ctx, mocks.MockEmployees[0])

	assert.NotNil(t, err)
	assert.EqualError(t, err, expectedErr)
	assert.Equal(t, 0, res)
}

func TestFindAllEmployee(t *testing.T) {

	s := createService()

	ctx := context.TODO()
	res, err := s.GetAll(ctx)

	//Sin error esperado
	assert.Nil(t, err)
	assert.Equal(t, mocks.MockEmployees, res)
}

func TestFindByIdNonExistentEmployee(t *testing.T) {

	eID := 4
	expectedErr := "El id no existe"

	s := createService()

	ctx := context.TODO()
	resp, err := s.Get(ctx, eID)

	//Con error esperado
	assert.NotNil(t, err)
	//Respuesta nula
	assert.Equal(t, domain.Employee{}, resp)
	assert.EqualError(t, err, expectedErr)
}

func TestFindByIdExistentEmployee(t *testing.T) {

	eID := 1

	s := createService()

	ctx := context.TODO()
	resp, err := s.Get(ctx, eID)

	assert.Nil(t, err)
	assert.Equal(t, eID, resp.ID)
	assert.NotEqual(t, 0, resp.ID)
}

func TestUpdateExistentEmployee(t *testing.T) {

	s := createService()

	ctx := context.TODO()
	err := s.Update(ctx, mocks.MockUpdateEmployee)
	data, _ := s.GetAll(ctx)

	assert.Nil(t, err)
	assert.Equal(t, data[1], mocks.MockUpdateEmployee)
}

func TestUpdateNonExistentEmployee(t *testing.T) {

	s := createService()

	ctx := context.TODO()
	err := s.Update(ctx, mocks.MockEmployeeNonExistent)

	assert.NotNil(t, err)
}

func TestDeleteNonExistentEmployee(t *testing.T) {

	s := createService()

	ctx := context.TODO()
	idDelete := 8
	err := s.Delete(ctx, idDelete)

	assert.NotNil(t, err)
}

func TestDeleteOkEmployee(t *testing.T) {

	s := createService()

	ctx := context.TODO()
	idDelete := 1
	err := s.Delete(ctx, idDelete)
	verificationConsult, _ := s.GetAll(ctx)

	assert.Nil(t, err)
	assert.NotEqual(t, len(mocks.MockEmployees), len(verificationConsult))
}
