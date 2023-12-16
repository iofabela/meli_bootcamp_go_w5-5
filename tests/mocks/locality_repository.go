package mocks

import (
	"database/sql"
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/iofabela/meli_bootcamp_go_w5-5/internal/domain"
	"github.com/iofabela/meli_bootcamp_go_w5-5/utils/queries"
)

var (
	LocalityTest = domain.Locality{
		ID: 1,
	}

	SellerTest = domain.Seller{
		CID: 1,
	}
)

func LocalityCreateOKMockDB() (*sql.DB, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, err
	}

	// ID Not Exists
	mock.
		ExpectQuery(regexp.QuoteMeta(queries.SelectIdLocality)).
		WithArgs(LocalityTest.ID).
		WillReturnError(sql.ErrNoRows)

	// Save OK
	mock.
		ExpectPrepare(regexp.QuoteMeta(queries.InsertLocality))

	mock.
		ExpectExec(regexp.QuoteMeta(queries.InsertLocality)).
		WithArgs(
			LocalityTest.ID,
			LocalityTest.LocalityName,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.
		ExpectPrepare(regexp.QuoteMeta(queries.InsertProvince))

	mock.
		ExpectExec(regexp.QuoteMeta(queries.InsertProvince)).
		WithArgs(
			LocalityTest.ProvinceName,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.
		ExpectPrepare(regexp.QuoteMeta(queries.InsertCountry))

	mock.
		ExpectExec(regexp.QuoteMeta(queries.InsertCountry)).
		WithArgs(
			LocalityTest.CountryName,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	return db, nil
}

func LocalityIdWithflictMockDB() (*sql.DB, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, err
	}

	// ID Exists
	rows := sqlmock.NewRows([]string{"id"})
	rows.AddRow(LocalityTest.ID)

	mock.
		ExpectQuery(regexp.QuoteMeta(queries.SelectIdLocality)).
		WithArgs(LocalityTest.ID).
		WillReturnRows(rows)

	return db, nil
}
