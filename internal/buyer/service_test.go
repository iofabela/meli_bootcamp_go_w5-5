package buyer

import (
	"context"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w5-5/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w5-5/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCreateOkBuyer(t *testing.T) {
	//arrange
	newBuyer := domain.Buyer{
		CardNumberID: "LZM12345",
		FirstName:    "PEDRO",
		LastName:     "FAUNDEZ",
	}

	//act
	mockRepo := &mocks.MockBuyerRepository{
		DataMock: mocks.MockDataBuyers,
	}
	service := NewService(mockRepo)

	result, err := service.Save(context.TODO(), newBuyer)

	//assert
	//No devuelva error
	assert.Nil(t, err)
	//El resultado es el mismo al esperado
	id := mockRepo.DataMock[2].ID
	assert.Equal(t, id, result)

	newBuyer.ID = id
	assert.Equal(t, newBuyer, mockRepo.DataMock[2])

}

func TestCreateConflictBuyer(t *testing.T) {
	//arrange
	newBuyer := domain.Buyer{
		CardNumberID: "ABC1234",
		FirstName:    "PEDRO",
		LastName:     "FAUNDEZ",
	}

	//act
	mockRepo := &mocks.MockBuyerRepository{
		DataMock: mocks.MockDataBuyers,
	}
	service := NewService(mockRepo)

	_, err := service.Save(context.TODO(), newBuyer)

	//assert
	//Devuelva error esperado
	assert.ErrorContains(t, err, "error: buyer with this card_number_id already exist")

}

func TestFindAllBuyer(t *testing.T) {
	//arrange
	dataBase := mocks.MockDataBuyers

	//act
	service := NewService(&mocks.MockBuyerRepository{
		DataMock: dataBase,
	})
	result, err := service.GetAll(context.TODO())

	//assert
	//No devuelva error
	assert.Nil(t, err)
	//El resultado es el mismo al esperado
	assert.Equal(t, dataBase, result)

}

func TestFindByIdNonExistentBuyer(t *testing.T) {
	//arrange
	dataBase := mocks.MockDataBuyers

	//act
	service := NewService(&mocks.MockBuyerRepository{
		DataMock: dataBase,
	})
	result, err := service.Get(context.TODO(), 4)

	//assert
	//Devuelva error esperado
	assert.ErrorContains(t, err, "error: buyer with id:4 not found")

	//Devuelve nulo
	assert.Equal(t, domain.Buyer{}, result)

}

func TestFindByIdExistentBuyer(t *testing.T) {
	//arrange
	dataBase := mocks.MockDataBuyers

	//act
	service := NewService(&mocks.MockBuyerRepository{
		DataMock: dataBase,
	})
	result, err := service.Get(context.TODO(), 2)

	//assert
	//No devuelva error
	assert.Nil(t, err)
	//El resultado es el mismo al esperado
	assert.Equal(t, dataBase[1], result)

}

func TestUpdateBuyer(t *testing.T) {
	//arrange

	dataBase := mocks.MockDataBuyers
	mockRepo := &mocks.MockBuyerRepository{
		DataMock: dataBase,
	}

	buyerToUpdate := domain.Buyer{
		ID:           1,
		CardNumberID: "ABC1234",
		FirstName:    "ALFREDO",
		LastName:     "LOPEZ",
	}

	//act
	service := NewService(mockRepo)
	err := service.Update(context.TODO(), buyerToUpdate)

	//assert
	//No devuelva error
	assert.Nil(t, err)
	//El resultado es el mismo al esperado
	assert.Equal(t, buyerToUpdate, mockRepo.DataMock[0])

}

func TestUpdateNonExistentBuyer(t *testing.T) {
	//arrange

	dataBase := mocks.MockDataBuyers
	mockRepo := &mocks.MockBuyerRepository{
		DataMock: dataBase,
	}

	buyerToUpdate := domain.Buyer{
		ID:           3,
		CardNumberID: "ABC1234",
		FirstName:    "ALFREDO",
		LastName:     "LOPEZ",
	}

	//act

	service := NewService(mockRepo)
	err := service.Update(context.TODO(), buyerToUpdate)

	//assert
	//Devuelve Nil cuando no existe el buyer
	assert.Nil(t, err)
}

func TestDeleteNonExistentBuyer(t *testing.T) {
	//arrange

	dataBase := mocks.MockDataBuyers
	mockRepo := &mocks.MockBuyerRepository{
		DataMock: dataBase,
	}

	//act

	service := NewService(mockRepo)
	err := service.Delete(context.TODO(), 3)

	//assert
	//Devuelve Nil cuando no existe el buyer
	assert.Nil(t, err)
}

func TestDeleteOkBuyer(t *testing.T) {
	//arrange

	dataBase := mocks.MockDataBuyers
	mockRepo := &mocks.MockBuyerRepository{
		DataMock: dataBase,
	}

	//act

	service := NewService(mockRepo)
	err := service.Delete(context.TODO(), 1)

	//assert
	//No devuelva error
	assert.Nil(t, err)

	//Eliminación exitosa
	assert.True(t, len(mockRepo.DataMock) == 1)
}

func TestExistsBuyer(t *testing.T) {

	//arrange

	dataBase := mocks.MockDataBuyers
	mockRepo := &mocks.MockBuyerRepository{
		DataMock: dataBase,
	}

	//act

	service := NewService(mockRepo)
	err := service.Exists(context.TODO(), "XYZ1234")

	//assert
	//Devuelve que existe
	assert.True(t, err)

}
