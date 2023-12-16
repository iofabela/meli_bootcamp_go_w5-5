package seller

import (
	"context"
	"testing"

	"github.com/iofabela/meli_bootcamp_go_w5-5/internal/domain"
	"github.com/iofabela/meli_bootcamp_go_w5-5/tests/mocks"

	"github.com/stretchr/testify/assert"
)

// test create ok
func TestCreateOkSeller(t *testing.T) {
	//arrange
	mockRepository := &mocks.MockSellerRepo{
		MockSeller: mocks.MockListSellers,
	}
	service := NewService(mockRepository)

	mocks.MockNewSeller.ID = 3
	//act
	resp, err := service.Save(context.TODO(), mocks.MockNewSeller)
	//assert
	assert.Nil(t, err)
	assert.Equal(t, mocks.MockNewSeller.ID, resp)
	assert.Equal(t, mocks.MockNewSeller, mockRepository.MockSeller[2])
}

// test create with conflict
func TestCreateConflictSeller(t *testing.T) {
	//arrange
	mockRepository := &mocks.MockSellerRepo{
		MockSeller: mocks.MockListSellers,
	}
	service := NewService(mockRepository)
	//act
	_, err := service.Save(context.TODO(), mocks.MockNewSellerWithConflict)
	//assert
	assert.NotNil(t, err)
}

// test find all sellers
func TestFindAllSeller(t *testing.T) {
	//arrange
	mockRepository := &mocks.MockSellerRepo{
		MockSeller: mocks.MockListSellers,
	}
	//act
	service := NewService(mockRepository)

	resp, err := service.GetAll(context.TODO())
	//assert
	assert.Nil(t, err)
	assert.Equal(t, mockRepository.MockSeller, resp)
}

// test find by no existen seller
func TestFindByIdNonExistentSeller(t *testing.T) {
	//arrange
	mockRepository := &mocks.MockSellerRepo{
		MockSeller: mocks.MockListSellers,
	}
	//act
	service := NewService(mockRepository)

	resp, err := service.Get(context.TODO(), 3)

	//assert
	assert.NotNil(t, err)
	assert.Equal(t, domain.Seller{}, resp)
}

func TestFindByIdExistentSeller(t *testing.T) {
	//arrange
	mockRepository := &mocks.MockSellerRepo{
		MockSeller: mocks.MockListSellers,
	}
	service := NewService(mockRepository)

	resp, err := service.Get(context.TODO(), 2)
	//assert
	assert.Nil(t, err)
	assert.Equal(t, mocks.MockListSellers[1], resp)
}

func TestUpdateOkSeller(t *testing.T) {
	//act
	mockRepository := &mocks.MockSellerRepo{
		MockSeller: mocks.MockListSellers,
	}
	s := NewService(mockRepository)
	ctx := context.TODO()
	err := s.Update(ctx, mocks.MockUpdateSeller)
	db, _ := s.GetAll(ctx)
	//assert
	assert.Nil(t, err)
	assert.Equal(t, db[1], mocks.MockUpdateSeller)
}

func TestUpdateNotExistentSeller(t *testing.T) {
	//arrange
	mockRepository := &mocks.MockSellerRepo{
		MockSeller: mocks.MockListSellers,
	}
	s := NewService(mockRepository)
	//act
	ctx := context.TODO()
	err := s.Update(ctx, mocks.MockNewSellerWithConflict)
	//assert
	assert.NotNil(t, err)

}

func TestDeleteOkSeller(t *testing.T) {
	//act
	sId := 10
	mockRepository := &mocks.MockSellerRepo{
		MockSeller: mocks.MockListSellers,
	}
	s := NewService(mockRepository)
	ctx := context.TODO()
	err := s.Delete(ctx, sId)
	delet, errSearch := s.Get(ctx, sId)
	//assert
	assert.NotNil(t, err)
	assert.NotNil(t, errSearch)
	assert.Equal(t, 0, delet.ID)
}

func TestDeleteNonExistentSeller(t *testing.T) {
	sId := 10
	mockRepository := &mocks.MockSellerRepo{
		MockSeller: mocks.MockListSellers,
	}
	s := NewService(mockRepository)
	ctx := context.TODO()
	err := s.Delete(ctx, sId)
	assert.NotNil(t, err)
}

func TestExistsSellerCid(t *testing.T) {
	mockRepository := &mocks.MockSellerRepo{
		MockSeller: mocks.MockListSellers,
	}
	service := NewService(mockRepository)

	exists := service.Exists(context.TODO(), 2)

	//assert
	assert.True(t, exists)
}

func TestNonExistsSellerCid(t *testing.T) {
	mockRepository := &mocks.MockSellerRepo{
		MockSeller: mocks.MockListSellers,
	}
	service := NewService(mockRepository)

	exists := service.Exists(context.TODO(), 7)

	//assert
	assert.False(t, exists)
}
