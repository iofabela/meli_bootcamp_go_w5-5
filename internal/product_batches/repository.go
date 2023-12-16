package product_batch

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/extmatperez/meli_bootcamp_go_w5-5/internal/domain"
)

// Repository encapsulates the storage of a section.
type Repository interface {
	Save(ctx context.Context, s domain.ProductBatches) (int, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) sectionExists(ctx context.Context, section_id int) bool {
	query := "SELECT id FROM sections WHERE id=?;"
	row := r.db.QueryRow(query, section_id)
	err := row.Scan(&section_id)
	return err == nil
}

func (r *repository) productExists(ctx context.Context, product_id int) bool {
	query := "SELECT id FROM products WHERE id=?;"
	row := r.db.QueryRow(query, product_id)
	err := row.Scan(&product_id)
	return err == nil
}

func (r *repository) Save(ctx context.Context, pd domain.ProductBatches) (int, error) {
	if !r.productExists(ctx, pd.ProductId) {
		return 0, fmt.Errorf("product with id: %d doesnt exists", pd.ProductId)
	}
	if !r.sectionExists(ctx, pd.SectionId) {
		return 0, fmt.Errorf("section with id: %d doesnt exists", pd.SectionId)
	}
	query := "INSERT INTO product_batches " +
		"(" +
		"batch_number, " +
		"current_quantity, " +
		"current_temperature, " +
		"due_date, " +
		"initial_quantity, " +
		"manufacturing_date, " +
		"manufacturing_hour, " +
		"minimum_temperature, " +
		"product_id, " +
		"section_id" +
		")" +
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(
		&pd.BatchNumber,
		&pd.CurrentQuantity,
		&pd.CurrentTemperature,
		&pd.DueDate,
		&pd.InitialQuantity,
		&pd.ManufacturingDate,
		&pd.ManufacturingHour,
		&pd.MinumumTemperature,
		&pd.ProductId,
		&pd.SectionId,
	)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
