package product_batch

import (
	"context"
	"testing"
	"time"

	"github.com/extmatperez/meli_bootcamp_go_w5-5/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestSave(t *testing.T) {
	mockDB, mock := mocks.CreateMockDBSave(t)
	ctx, close := context.WithTimeout(context.Background(), time.Second*5)
	defer close()
	mockedRepo := NewRepository(mockDB)
	res, err := mockedRepo.Save(ctx, mocks.MockProductBatch)
	assert.NoError(t, err)
	assert.Equal(t, mocks.MockProductBatch.Id, res)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSaveNonExistantProduct(t *testing.T) {
	mockDB, mock := mocks.CreateMockDBSaveNonExistantProduct(t)
	ctx, close := context.WithTimeout(context.Background(), time.Second*5)
	defer close()
	mockedRepo := NewRepository(mockDB)
	res, err := mockedRepo.Save(ctx, mocks.MockProductBatchNonExistantProduct)
	assert.Error(t, err)
	assert.Empty(t, res)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSaveNonExistantSection(t *testing.T) {
	mockDB, mock := mocks.CreateMockDBSaveNonExistantSection(t)
	ctx, close := context.WithTimeout(context.Background(), time.Second*5)
	defer close()
	mockedRepo := NewRepository(mockDB)
	res, err := mockedRepo.Save(ctx, mocks.MockProductBatchNonExistantSection)
	assert.Error(t, err)
	assert.Empty(t, res)
	assert.NoError(t, mock.ExpectationsWereMet())
}
