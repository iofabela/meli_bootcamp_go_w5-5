package mocks

import (
	"database/sql"
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/extmatperez/meli_bootcamp_go_w5-5/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w5-5/utils/queries"
)

var (
	CarryTest = domain.Carry{
		CID:        "ABC",
		LocalityID: 3,
	}
)

func CarryCreateOKMockDB() (*sql.DB, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, err
	}

	// CID Not Exists
	mock.
		ExpectPrepare(regexp.QuoteMeta(queries.CarryCIDExistsQuery))
	mock.
		ExpectQuery(regexp.QuoteMeta(queries.CarryCIDExistsQuery)).
		WithArgs(CarryTest.CID).
		WillReturnError(sql.ErrNoRows)

	// Locality Exists
	rows := sqlmock.NewRows([]string{"id"})
	rows.AddRow(CarryTest.LocalityID)

	mock.
		ExpectPrepare(regexp.QuoteMeta(queries.CarryLocalityExistsQuery))
	mock.
		ExpectQuery(regexp.QuoteMeta(queries.CarryLocalityExistsQuery)).
		WithArgs(CarryTest.LocalityID).
		WillReturnRows(rows)

	// Save OK
	mock.
		ExpectPrepare(regexp.QuoteMeta(queries.CarrySaveQuery))
	mock.
		ExpectExec(regexp.QuoteMeta(queries.CarrySaveQuery)).
		WithArgs(
			CarryTest.CID,
			CarryTest.CompanyName,
			CarryTest.Address,
			CarryTest.Telephone,
			CarryTest.LocalityID,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	return db, nil
}

func CarryCIDConflictMockDB() (*sql.DB, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, err
	}

	// CID Exists
	rows := sqlmock.NewRows([]string{"cid"})
	rows.AddRow(CarryTest.CID)

	mock.
		ExpectPrepare(regexp.QuoteMeta(queries.CarryCIDExistsQuery))
	mock.
		ExpectQuery(regexp.QuoteMeta(queries.CarryCIDExistsQuery)).
		WithArgs(CarryTest.CID).
		WillReturnRows(rows)

	return db, nil
}

func CarryLocalityConflictMockDB() (*sql.DB, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, err
	}

	// CID Not Exists
	mock.
		ExpectPrepare(regexp.QuoteMeta(queries.CarryCIDExistsQuery))
	mock.
		ExpectQuery(regexp.QuoteMeta(queries.CarryCIDExistsQuery)).
		WithArgs(CarryTest.CID).
		WillReturnError(sql.ErrNoRows)

	// Locality Not Exists
	mock.
		ExpectPrepare(regexp.QuoteMeta(queries.CarryLocalityExistsQuery))
	mock.
		ExpectQuery(regexp.QuoteMeta(queries.CarryLocalityExistsQuery)).
		WithArgs(CarryTest.LocalityID).
		WillReturnError(sql.ErrNoRows)

	return db, nil
}
