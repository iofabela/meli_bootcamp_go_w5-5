package locality

import (
	"context"
	"fmt"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w5-5/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestLocalityCreateOk(t *testing.T) {
	db, err := mocks.LocalityCreateOKMockDB()
	assert.NoError(t, err)
	repo := NewRepository(db)

	result, err := repo.SaveLocality(context.TODO(), mocks.LocalityTest)

	assert.NoError(t, err)
	assert.Equal(t, 1, result)
}

func TestLocalityCreateWithConflict(t *testing.T) {
	db, err := mocks.LocalityIdWithflictMockDB()
	assert.NoError(t, err)
	repo := NewRepository(db)

	result, err := repo.SaveLocality(context.TODO(), mocks.LocalityTest)

	assert.Error(t, err)
	assert.Zero(t, result)
	assert.EqualError(t, err, fmt.Sprintf("locality with id %v already exists", mocks.LocalityTest.ID))
}
