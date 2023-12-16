package inboundorder

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/iofabela/meli_bootcamp_go_w5-5/internal/domain"
)

// Repository encapsulates a repository interface
type Repository interface {
	Exists(ctx context.Context, employeeID int) (bool, error)
	Save(ctx context.Context, i domain.InboundOrder) (int, error)
}

type repository struct {
	db *sql.DB
}

// NewRepository creates a new repository
func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Exists(ctx context.Context, employeeID int) (bool, error) {
	query := "SELECT id FROM employees WHERE id=?;"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, employeeID)
	err = row.Scan(&employeeID)
	return err == nil, nil
}

func (r *repository) Save(ctx context.Context, i domain.InboundOrder) (int, error) {
	exist, err := r.Exists(ctx, i.EmployeeID)
	if err != nil {
		return 0, err
	} else if !exist {
		return 0, fmt.Errorf("error. The employee with the id: %v, not exists", i.EmployeeID)
	}

	query := "INSERT INTO inbound_orders(order_date, order_number, employe_id, product_batch_id, wareHouse_id) VALUES (?,?,?,?,?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, &i.OrderDate, &i.OrderNumber, &i.EmployeeID, &i.ProductBatchID, &i.WarehouseID)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
