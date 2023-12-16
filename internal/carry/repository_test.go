package carry

import (
	"context"
	"fmt"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w5-5/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCarryCreateOk(t *testing.T) {
	db, err := mocks.CarryCreateOKMockDB()
	assert.NoError(t, err)
	repo := NewRepository(db)

	result, err := repo.Save(context.TODO(), mocks.CarryTest)

	assert.NoError(t, err)
	assert.Equal(t, 1, result)
}

func TestCarryCreateConflictCID(t *testing.T) {
	db, err := mocks.CarryCIDConflictMockDB()
	assert.NoError(t, err)
	repo := NewRepository(db)

	result, err := repo.Save(context.TODO(), mocks.CarryTest)

	assert.Error(t, err)
	assert.Zero(t, result)
	assert.EqualError(t, err, fmt.Sprintf("carry with cid %v already exists", mocks.CarryTest.CID))
}

func TestCarryCreateConflictLocalityID(t *testing.T) {
	db, err := mocks.CarryLocalityConflictMockDB()
	assert.NoError(t, err)
	repo := NewRepository(db)

	result, err := repo.Save(context.TODO(), mocks.CarryTest)

	assert.Error(t, err)
	assert.Zero(t, result)
	assert.EqualError(
		t,
		err,
		fmt.Sprintf("locality with id %v not exists", mocks.CarryTest.LocalityID),
	)
}
