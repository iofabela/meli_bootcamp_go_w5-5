package warehouse

import (
	"context"
	"database/sql"

	"github.com/extmatperez/meli_bootcamp_go_w5-5/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w5-5/utils/queries"
)

// Repository encapsulates the storage of a warehouse.
type Repository interface {
	GetAll(ctx context.Context) ([]domain.Warehouse, error)
	Get(ctx context.Context, id int) (domain.Warehouse, error)
	Exists(ctx context.Context, warehouseCode string) bool
	Save(ctx context.Context, w domain.Warehouse) (int, error)
	Update(ctx context.Context, w domain.Warehouse) error
	Delete(ctx context.Context, id int) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Warehouse, error) {
	rows, err := r.db.QueryContext(ctx, queries.WarehouseGetAllQuery)
	if err != nil {
		return nil, err
	}

	var warehouses []domain.Warehouse

	for rows.Next() {
		w := domain.Warehouse{}
		_ = rows.Scan(&w.ID, &w.Address, &w.Telephone, &w.WarehouseCode, &w.MinimumCapacity, &w.MinimumTemperature)
		warehouses = append(warehouses, w)
	}

	return warehouses, nil
}

func (r *repository) Get(ctx context.Context, id int) (domain.Warehouse, error) {
	row := r.db.QueryRowContext(ctx, queries.WarehouseGetQuery, id)
	w := domain.Warehouse{}
	err := row.Scan(&w.ID, &w.Address, &w.Telephone, &w.WarehouseCode, &w.MinimumCapacity, &w.MinimumTemperature)
	if err != nil {
		return domain.Warehouse{}, err
	}

	return w, nil
}

func (r *repository) Exists(ctx context.Context, warehouseCode string) bool {
	row := r.db.QueryRowContext(ctx, queries.WarehouseExistsQuery, warehouseCode)
	err := row.Scan(&warehouseCode)
	return err == nil
}

func (r *repository) Save(ctx context.Context, w domain.Warehouse) (int, error) {
	stmt, err := r.db.PrepareContext(ctx, queries.WarehouseSaveQuery)
	if err != nil {
		return 0, err
	}

	res, err := stmt.ExecContext(ctx, &w.Address, &w.Telephone, &w.WarehouseCode, &w.MinimumCapacity, &w.MinimumTemperature)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *repository) Update(ctx context.Context, w domain.Warehouse) error {
	stmt, err := r.db.PrepareContext(ctx, queries.WarehouseUpdateQuery)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, &w.Address, &w.Telephone, &w.WarehouseCode, &w.MinimumCapacity, &w.MinimumTemperature, &w.ID)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	stmt, err := r.db.PrepareContext(ctx, queries.WarehouseDeleteQuery)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affect < 1 {
		return ErrNotFound
	}

	return nil
}
