package purchaseOrder

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/iofabela/meli_bootcamp_go_w5-5/internal/domain"
	"github.com/iofabela/meli_bootcamp_go_w5-5/utils/queries"
	"github.com/stretchr/testify/assert"
)

func TestCreateOk(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	purchaseOrder := domain.PurchaseOrders{
		ID:          1,
		OrderNumber: "ABC1234",
	}
	mock.
		ExpectPrepare(regexp.QuoteMeta(queries.PurchaseOrderInsertIntoPO))
	mock.
		ExpectExec(regexp.QuoteMeta(queries.PurchaseOrderInsertIntoPO)).
		WithArgs(purchaseOrder.OrderNumber, purchaseOrder.OrderDate, purchaseOrder.TrackingCode, purchaseOrder.BuyerId, purchaseOrder.OrderStatusId).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.
		ExpectPrepare(regexp.QuoteMeta(queries.PurchaseOrderInsertIntoOD))
	mock.
		ExpectExec(regexp.QuoteMeta(queries.PurchaseOrderInsertIntoOD)).
		WithArgs(purchaseOrder.ProductRecordId).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewRepository(db)

	ctx := context.TODO()
	p, err := repo.Save(ctx, purchaseOrder)
	assert.NoError(t, err)
	assert.NotZero(t, p)
	assert.Equal(t, purchaseOrder.ID, p)
}

func TestCreateConflict(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Se define producto con Order Number vacío

	purchaseOrder := domain.PurchaseOrders{
		ID:          1,
		OrderNumber: "",
	}

	mock.
		ExpectPrepare(regexp.QuoteMeta(queries.PurchaseOrderInsertIntoPO))
	mock.
		ExpectExec(regexp.QuoteMeta(queries.PurchaseOrderInsertIntoPO)).
		WithArgs(purchaseOrder.OrderNumber, purchaseOrder.OrderDate, purchaseOrder.TrackingCode, purchaseOrder.BuyerId, purchaseOrder.OrderStatusId).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.
		ExpectPrepare(regexp.QuoteMeta(queries.PurchaseOrderInsertIntoOD))
	mock.
		ExpectExec(regexp.QuoteMeta(queries.PurchaseOrderInsertIntoOD)).
		WithArgs(purchaseOrder.ProductRecordId).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewRepository(db)

	ctx := context.TODO()
	p, err := repo.Save(ctx, purchaseOrder)
	assert.Equal(t, 0, p)
	assert.ErrorContains(t, err, "error: order_number empty")
}

func TestExistNonConflict(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	purchaseOrder := domain.PurchaseOrders{}

	mock.
		ExpectPrepare(regexp.QuoteMeta(queries.PurchaseOrderSelectOrderNumber))
	mock.
		ExpectExec(regexp.QuoteMeta(queries.PurchaseOrderSelectOrderNumber)).
		WithArgs(purchaseOrder.OrderNumber).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewRepository(db)

	ctx := context.TODO()
	p, err := repo.Exists(ctx, "ABC1234")
	assert.NoError(t, err)
	assert.Equal(t, false, p)

}

func TestExistConflict(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	purchaseOrder := domain.PurchaseOrders{}

	mock.
		ExpectPrepare(regexp.QuoteMeta(queries.PurchaseOrderSelectOrderNumber)).WillReturnError(sql.ErrConnDone)
	mock.
		ExpectExec(regexp.QuoteMeta(queries.PurchaseOrderSelectOrderNumber)).
		WithArgs(purchaseOrder.OrderNumber).
		WillReturnError(sql.ErrNoRows)

	repo := NewRepository(db)

	ctx := context.TODO()
	p, err := repo.Exists(ctx, "ABC1234")
	assert.Error(t, err)
	assert.Equal(t, false, p)

}

// func TestExistNoRows(t *testing.T) {

// 	db, mock, err := sqlmock.New()
// 	assert.NoError(t, err)
// 	defer db.Close()

// 	purchaseOrder := domain.PurchaseOrders{}

// 	mock.
// 		ExpectPrepare(regexp.QuoteMeta(queries.PurchaseOrderSelectOrderNumber))
// 	mock.
// 		ExpectExec(regexp.QuoteMeta(queries.PurchaseOrderSelectOrderNumber)).
// 		WithArgs(purchaseOrder.OrderNumber).
// 		WillReturnError(sql.ErrNoRows)

// 	repo := NewRepository(db)

// 	ctx := context.TODO()
// 	p, err := repo.Exists(ctx, "ABC1234")
// 	assert.Nil(t, err)
// 	assert.Equal(t, false, p)

// }

func TestCreateConflictSave(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	purchaseOrder := domain.PurchaseOrders{
		ID:          1,
		OrderNumber: "ABC1234",
	}
	mock.
		ExpectPrepare(regexp.QuoteMeta(queries.PurchaseOrderInsertIntoPO)).WillReturnError(sql.ErrConnDone)

	repo := NewRepository(db)

	ctx := context.TODO()
	p, err := repo.Save(ctx, purchaseOrder)
	assert.Error(t, err)
	assert.Equal(t, 0, p)

}
