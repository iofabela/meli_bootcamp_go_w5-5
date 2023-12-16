package warehouse

import (
	"context"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w5-5/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w5-5/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCreateOkWarehouse(t *testing.T) {
	mockRepository := &mocks.MockWarehouseRepository{
		MockData: mocks.MockDataWarehouse,
	}
	service := NewService(mockRepository)

	warehouseOk := domain.Warehouse{
		Address:            "Calle Falsa 123",
		Telephone:          "+54 9 11 5487-5421",
		WarehouseCode:      "CTX-458",
		MinimumCapacity:    5,
		MinimumTemperature: 7,
	}

	warehouseOk.ID = 3

	resp, err := service.Save(context.TODO(), warehouseOk)

	//Test sin error
	assert.Nil(t, err)

	//Test ID retornado igual a ID esperado
	assert.Equal(t, warehouseOk.ID, resp)

	//Test BBDD actualizada correctamente
	assert.Equal(t, warehouseOk, mockRepository.MockData[2])
}

func TestCreateConflictWarehouse(t *testing.T) {
	mockRepository := &mocks.MockWarehouseRepository{
		MockData: mocks.MockDataWarehouse,
	}
	service := NewService(mockRepository)

	warehouseOk := domain.Warehouse{
		Address:            "Calle Falsa 123",
		Telephone:          "+54 9 11 5487-5421",
		WarehouseCode:      "CTX-555",
		MinimumCapacity:    5,
		MinimumTemperature: 7,
	}

	_, err := service.Save(context.TODO(), warehouseOk)

	//Test error encontrado
	assert.NotNil(t, err)
}

func TestFindAllWarehouse(t *testing.T) {
	mockRepository := &mocks.MockWarehouseRepository{
		MockData: mocks.MockDataWarehouse,
	}
	service := NewService(mockRepository)

	resp, err := service.GetAll(context.TODO())

	//Test sin error
	assert.Nil(t, err)

	//Test lista retornada igual a lista esperada
	assert.Equal(t, mockRepository.MockData, resp)
}

func TestFindByIdNonExistentWarehouse(t *testing.T) {
	mockRepository := &mocks.MockWarehouseRepository{
		MockData: mocks.MockDataWarehouse,
	}
	service := NewService(mockRepository)

	resp, err := service.Get(context.TODO(), 3)

	//Test con error
	assert.NotNil(t, err)

	//Test respuesta nula
	assert.Equal(t, domain.Warehouse{}, resp)
}

func TestFindByIdExistentWarehouse(t *testing.T) {
	mockRepository := &mocks.MockWarehouseRepository{
		MockData: mocks.MockDataWarehouse,
	}
	service := NewService(mockRepository)

	resp, err := service.Get(context.TODO(), 2)

	//Test sin error
	assert.Nil(t, err)

	//Test respuesta correcta
	assert.Equal(t, mocks.MockDataWarehouse[1], resp)
}

func TestUpdateExistentWarehouse(t *testing.T) {
	mockRepository := &mocks.MockWarehouseRepository{
		MockData: mocks.MockDataWarehouse,
	}
	service := NewService(mockRepository)

	warehouseUpdated := mocks.MockDataWarehouse[0]
	warehouseUpdated.MinimumCapacity = 10
	warehouseUpdated.MinimumTemperature = 17

	err := service.Update(context.TODO(), warehouseUpdated)

	//Test sin error
	assert.Nil(t, err)

	//Test BBDD actualizada correctamente
	assert.Equal(t, warehouseUpdated, mockRepository.MockData[0])
}

func TestUpdateNonExistentWarehouse(t *testing.T) {
	mockRepository := &mocks.MockWarehouseRepository{
		MockData: mocks.MockDataWarehouse,
	}
	service := NewService(mockRepository)

	oldDATA := mockRepository.MockData
	warehouseUpdated := mockRepository.MockData[0]
	warehouseUpdated.MinimumCapacity = 10
	warehouseUpdated.MinimumTemperature = 17
	warehouseUpdated.ID = 3

	err := service.Update(context.TODO(), warehouseUpdated)

	//Test con error
	assert.NotNil(t, err)

	//Test BBDD no actualizada
	assert.Equal(t, oldDATA, mockRepository.MockData)
}

func TestDeleteNonExistentWarehouse(t *testing.T) {
	mockRepository := &mocks.MockWarehouseRepository{
		MockData: mocks.MockDataWarehouse,
	}
	service := NewService(mockRepository)

	oldDATA := mockRepository.MockData

	err := service.Delete(context.TODO(), 3)

	//Test con error
	assert.NotNil(t, err)

	//Test BBDD no actualizada
	assert.Equal(t, oldDATA, mockRepository.MockData)
}

func TestDeleteExistentWarehouse(t *testing.T) {
	var data []domain.Warehouse = []domain.Warehouse{}
	data = append(data, mocks.MockDataWarehouse...)

	mockRepository := &mocks.MockWarehouseRepository{
		MockData: data,
	}
	service := NewService(mockRepository)
	oldDATA := mockRepository.MockData

	err := service.Delete(context.TODO(), 1)

	//Test sin error
	assert.Nil(t, err)
	//Test BBDD actualizada correctamente
	assert.Equal(t, oldDATA[1:], mockRepository.MockData)
}

func TestExistsWarehouseCode(t *testing.T) {
	mockRepository := &mocks.MockWarehouseRepository{
		MockData: mocks.MockDataWarehouse,
	}
	service := NewService(mockRepository)

	exists := service.Exists(context.TODO(), "CTX-555")

	//Test de Existencia verdadero
	assert.True(t, exists)
}

func TestNonExistsWarehouseCode(t *testing.T) {
	mockRepository := &mocks.MockWarehouseRepository{
		MockData: mocks.MockDataWarehouse,
	}
	service := NewService(mockRepository)

	exists := service.Exists(context.TODO(), "CTX-777")

	//Test de Existencia falso
	assert.False(t, exists)
}
