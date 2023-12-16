package locality

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/iofabela/meli_bootcamp_go_w5-5/internal/domain"
	"github.com/iofabela/meli_bootcamp_go_w5-5/utils/queries"
)

type Repository interface {
	SaveLocality(ctx context.Context, l domain.Locality) (int, error)
	IDExist(ctx context.Context, id int) bool
	SellerReport(ctx context.Context, id int) (domain.ReportSeller, error)
	GetAllSellerReports(ctx context.Context) ([]domain.ReportSeller, error)
	GetCarryReport(ctx context.Context, id int) (domain.LocalityCarries, error)
	GetAllCarryReports(ctx context.Context) ([]domain.LocalityCarries, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

// Save Locality
func (r *repository) SaveLocality(ctx context.Context, l domain.Locality) (int, error) {
	if r.IDExist(ctx, l.ID) {
		return 0, fmt.Errorf("locality with id %v already exists", l.ID)
	}

	stmt, err := r.db.PrepareContext(ctx, queries.InsertLocality)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, &l.ID, &l.LocalityName)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	stmt, err = r.db.Prepare(queries.InsertProvince)
	if err != nil {
		return 0, err
	}

	res, err = stmt.ExecContext(ctx, &l.ProvinceName)
	if err != nil {
		return 0, err
	}

	stmt, err = r.db.Prepare(queries.InsertCountry)
	if err != nil {
		return 0, err
	}

	res, err = stmt.ExecContext(ctx, &l.CountryName)
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

// Validate Id
func (r *repository) IDExist(ctx context.Context, id int) bool {
	row := r.db.QueryRowContext(ctx, queries.SelectIdLocality, id)
	err := row.Scan(&id)
	return err == nil
}

// Report Seller for Id
func (r *repository) SellerReport(ctx context.Context, id int) (domain.ReportSeller, error) {
	stmt, err := r.db.PrepareContext(ctx, queries.LocalityGetSellerReportQuery)
	if err != nil {
		return domain.ReportSeller{}, err
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, id)

	lc := domain.ReportSeller{}
	if err := row.Scan(&lc.LocalityID, &lc.LocalityName, &lc.SellersCount); err != nil {
		if err == sql.ErrNoRows {
			return domain.ReportSeller{}, errors.New("seller not found")
		}
		return domain.ReportSeller{}, err
	}
	return lc, nil
}

// Report Seller All
func (r *repository) GetAllSellerReports(ctx context.Context) ([]domain.ReportSeller, error) {
	stmt, err := r.db.PrepareContext(ctx, queries.LocalityGetAllSellerReportsQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	var lcs []domain.ReportSeller

	for rows.Next() {
		lc := domain.ReportSeller{}
		_ = rows.Scan(&lc.LocalityID, &lc.LocalityName, &lc.SellersCount)
		lcs = append(lcs, lc)
	}

	return lcs, nil
}

func (r *repository) GetCarryReport(ctx context.Context, id int) (domain.LocalityCarries, error) {
	// Preparo la query
	stmt, err := r.db.PrepareContext(ctx, queries.LocalityGetCarryReportQuery)
	if err != nil {
		return domain.LocalityCarries{}, err
	}
	defer stmt.Close()

	// Ejecuto la query
	row := stmt.QueryRowContext(ctx, id)

	lc := domain.LocalityCarries{}
	// Obtengo el resultado o retorno error
	if err := row.Scan(&lc.LocalityID, &lc.LocalityName, &lc.CarriesCount); err != nil {
		// Si no hay filas retorno un not found
		if err == sql.ErrNoRows {
			return domain.LocalityCarries{}, errors.New("locality not found")
		}
		// En otro caso retorno el error obtenido
		return domain.LocalityCarries{}, err
	}
	// En caso de exito retorno el resultado
	return lc, nil
}

func (r *repository) GetAllCarryReports(ctx context.Context) ([]domain.LocalityCarries, error) {
	// Preparo la query
	stmt, err := r.db.PrepareContext(ctx, queries.LocalityGetAllCarryReportsQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Ejecuto la query
	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	var lcs []domain.LocalityCarries

	// Escaneo todas las filas y las agrego al slice
	for rows.Next() {
		lc := domain.LocalityCarries{}
		_ = rows.Scan(&lc.LocalityID, &lc.LocalityName, &lc.CarriesCount)
		lcs = append(lcs, lc)
	}

	return lcs, nil
}
