package carry

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/iofabela/meli_bootcamp_go_w5-5/internal/domain"
	"github.com/iofabela/meli_bootcamp_go_w5-5/utils/queries"
)

// Repository encapsulates the storage of a warehouse.
type Repository interface {
	Save(ctx context.Context, w domain.Carry) (int, error)
	CIDExists(ctx context.Context, cid string) (bool, error)
	LocalityExists(ctx context.Context, id int) (bool, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Save(ctx context.Context, c domain.Carry) (int, error) {
	// Verifico que no exista el CI
	cidExists, err := r.CIDExists(ctx, c.CID)
	if err != nil {
		return 0, err
	}
	if cidExists {
		return 0, fmt.Errorf("carry with cid %v already exists", c.CID)
	}

	//Verifico que exista el LocalityID
	localityExists, err := r.LocalityExists(ctx, c.LocalityID)
	if err != nil {
		return 0, err
	}
	if !localityExists {
		return 0, fmt.Errorf("locality with id %v not exists", c.LocalityID)
	}

	// Preparo el query
	stmt, err := r.db.PrepareContext(ctx, queries.CarrySaveQuery)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	// Ejecuto el query
	res, err := stmt.ExecContext(ctx, &c.CID, &c.CompanyName, &c.Address, &c.Telephone, &c.LocalityID)
	if err != nil {
		return 0, err
	}

	// Obtengo el id
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	// Retorno si todo salio bien
	return int(id), nil
}

func (r *repository) CIDExists(ctx context.Context, cid string) (bool, error) {
	stmt, err := r.db.PrepareContext(ctx, queries.CarryCIDExistsQuery)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, cid)
	err = row.Scan(&cid)
	if err == sql.ErrNoRows {
		return false, nil
	}
	return err == nil, err
}

func (r *repository) LocalityExists(ctx context.Context, id int) (bool, error) {
	stmt, err := r.db.PrepareContext(ctx, queries.CarryLocalityExistsQuery)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, id)
	err = row.Scan(&id)
	if err == sql.ErrNoRows {
		return false, nil
	}
	return err == nil, err
}
