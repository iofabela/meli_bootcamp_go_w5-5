package seller

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/iofabela/meli_bootcamp_go_w5-5/internal/domain"
	"github.com/iofabela/meli_bootcamp_go_w5-5/utils/queries"
)

// Repository encapsulates the storage of a Seller.
type Repository interface {
	GetAll(ctx context.Context) ([]domain.Seller, error)
	Get(ctx context.Context, id int) (domain.Seller, error)
	Exists(ctx context.Context, cid int) bool
	Save(ctx context.Context, s domain.Seller) (int, error)
	Update(ctx context.Context, s domain.Seller) error
	Delete(ctx context.Context, id int) error
	CIDExist(ctx context.Context, cid int) bool
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Seller, error) {
	query := "SELECT * FROM sellers"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	var sellers []domain.Seller

	for rows.Next() {
		s := domain.Seller{}
		_ = rows.Scan(&s.ID, &s.CID, &s.CompanyName, &s.Address, &s.Telephone)
		sellers = append(sellers, s)
	}

	return sellers, nil
}

func (r *repository) Get(ctx context.Context, id int) (domain.Seller, error) {
	query := "SELECT * FROM sellers WHERE id=?;"
	row := r.db.QueryRow(query, id)
	s := domain.Seller{}
	err := row.Scan(&s.ID, &s.CID, &s.CompanyName, &s.Address, &s.Telephone)
	if err != nil {
		return domain.Seller{}, err
	}

	return s, nil
}

func (r *repository) Exists(ctx context.Context, cid int) bool {
	query := "SELECT cid FROM sellers WHERE cid=?;"
	row := r.db.QueryRow(query, cid)
	err := row.Scan(&cid)
	return err == nil
}

func (r *repository) Save(ctx context.Context, s domain.Seller) (int, error) {
	if r.CIDExist(ctx, s.CID) {
		return 0, fmt.Errorf("seller with cid %v already exists", s.CID)
	}

	stmt, err := r.db.PrepareContext(ctx, queries.InsertSeller)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, &s.CID, &s.CompanyName, &s.Address, &s.Telephone, &s.LocalityId)
	if err != nil {
		return 0, err
	}
	stmt.Close()
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *repository) Update(ctx context.Context, s domain.Seller) error {
	query := "UPDATE sellers SET cid=?, company_name=?, address=?, telephone=? WHERE id=?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(s.CID, s.CompanyName, s.Address, s.Telephone, s.ID)
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
	query := "DELETE FROM sellers WHERE id=?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(id)
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

func (r *repository) CIDExist(ctx context.Context, cid int) bool {
	row := r.db.QueryRowContext(ctx, queries.SelectCidSeller, cid)
	err := row.Scan(&cid)
	return err == nil
}
