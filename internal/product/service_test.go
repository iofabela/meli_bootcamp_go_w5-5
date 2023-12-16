package product

import (
	"context"
	"fmt"
	"testing"

	"github.com/iofabela/meli_bootcamp_go_w5-5/tests/mocks"
	"github.com/stretchr/testify/assert"
)

// Test to check the creation of a product
func TestCreateOk(t *testing.T) {
	// act
	data := mocks.MockListProducts
	repository := &mocks.MockRepositoryProduct{
		Data: data,
	}
	service := NewService(repository)
	// Test Execution
	ctx := context.TODO()
	_, err := service.Save(ctx, mocks.MockCreateProduct)
	// Validation
	assert.Nil(t, err)
}

// Conflict test to create
func TestCreateConflict(t *testing.T) {
	// act
	data := mocks.MockListProducts
	repository := &mocks.MockRepositoryProduct{
		Data: data,
	}
	service := NewService(repository)
	// Test Execution
	ctx := context.TODO()
	_, err := service.Save(ctx, mocks.MockListProducts[0])
	// Validation
	assert.NotNil(t, err)
}

// Test to find all data
func TestFindAll(t *testing.T) {
	// act
	data := mocks.MockListProducts
	repository := &mocks.MockRepositoryProduct{
		Data: data,
	}
	service := NewService(repository)
	// Test Execution
	ctx := context.TODO()
	_, err := service.GetAll(ctx)
	// Validation
	assert.Nil(t, err)
}

// Test to find a data with the ID
func TestFinByIdExistent(t *testing.T) {
	// act
	data := mocks.MockListProducts
	repository := &mocks.MockRepositoryProduct{
		Data: data,
	}
	service := NewService(repository)
	// Test Execution
	idSelected := 2
	ctx := context.TODO()
	res, err := service.Get(ctx, idSelected)
	// Validation
	assert.Nil(t, err)
	assert.Equal(t, idSelected, res.ID)
}

// Test to check a non-existing data with the ID
func TestFinByIdNonExistent(t *testing.T) {
	// act
	data := mocks.MockListProducts
	repository := &mocks.MockRepositoryProduct{
		Data: data,
	}
	service := NewService(repository)
	// Test Execution
	idSelected := 5
	ctx := context.TODO()
	_, err := service.Get(ctx, idSelected)

	expectedError := fmt.Errorf(mocks.ProductNotFound, idSelected)

	// Validation
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, expectedError.Error())
}

func TestUpdateExistent(t *testing.T) {
	// act
	data := mocks.MockListProducts
	repository := &mocks.MockRepositoryProduct{
		Data: data,
	}
	service := NewService(repository)
	// Test Execution
	ctx := context.TODO()
	err := service.Update(ctx, mocks.MockUpdateProduct)
	db, _ := service.GetAll(ctx)
	// Validation
	assert.Nil(t, err)
	assert.NotEqual(t, db[1], mocks.MockUpdateProduct)
}

func TestUpdateNonExistent(t *testing.T) {
	// act
	data := mocks.MockListProducts
	repository := &mocks.MockRepositoryProduct{
		Data: data,
	}
	service := NewService(repository)
	// Test Execution
	ctx := context.TODO()
	err := service.Update(ctx, mocks.MockProductNonExistent)
	assert.NotNil(t, err)
}

func TestDeleteNonExistent(t *testing.T) {
	// act
	data := mocks.MockListProducts
	repository := &mocks.MockRepositoryProduct{
		Data: data,
	}
	service := NewService(repository)
	// Test Execution
	idSelected := 5
	ctx := context.TODO()
	err := service.Delete(ctx, idSelected)
	// Validation
	assert.NotNil(t, err)
}

func TestDelete(t *testing.T) {
	// act
	data := mocks.MockListProducts
	repository := &mocks.MockRepositoryProduct{
		Data: data,
	}
	service := NewService(repository)
	// Test Execution
	idSelected := 1
	ctx := context.TODO()
	err := service.Delete(ctx, idSelected)
	dataAfterDeleted, _ := repository.GetAll(ctx)

	// Validation
	assert.Nil(t, err)
	assert.NotEqual(t, len(mocks.MockListProducts), len(dataAfterDeleted),
		"Deben contener distinto tamaño en la lista")
}

func TestExists(t *testing.T) {
	// act
	data := mocks.MockListProducts
	repository := &mocks.MockRepositoryProduct{
		Data: data,
	}
	service := NewService(repository)
	// Test Execution
	ctx := context.TODO()
	check := service.Exists(ctx, mocks.MockListProducts[1].ProductCode)

	// Validation
	assert.True(t, check)
}
