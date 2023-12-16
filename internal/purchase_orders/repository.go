package purchaseOrder

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/iofabela/meli_bootcamp_go_w5-5/internal/domain"
	"github.com/iofabela/meli_bootcamp_go_w5-5/utils/queries"
)

type Repository interface {
	Exists(ctx context.Context, orderNumber string) (bool, error)
	Save(ctx context.Context, b domain.PurchaseOrders) (int, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Exists(ctx context.Context, orderNumber string) (bool, error) {
	query := queries.PurchaseOrderSelectOrderNumber
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return false, err
	}

	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, orderNumber)
	err = row.Scan(&orderNumber)
	if err == sql.ErrNoRows {
		return false, nil
	}

	return err == nil, nil
}

func (r *repository) Save(ctx context.Context, po domain.PurchaseOrders) (int, error) {

	if po.OrderNumber == "" {
		return 0, fmt.Errorf("error: order_number empty")
	}

	query := queries.PurchaseOrderInsertIntoPO
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, &po.OrderNumber, &po.OrderDate, &po.TrackingCode, &po.BuyerId, &po.OrderStatusId)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	query = queries.PurchaseOrderInsertIntoOD
	stmt, err = r.db.Prepare(query)
	if err != nil {
		return 0, err
	}

	res, err = stmt.ExecContext(ctx, &po.ProductRecordId)
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
