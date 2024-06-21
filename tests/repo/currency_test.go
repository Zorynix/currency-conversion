package repo_test

import (
	"context"
	"testing"

	"currency-conversion/internal/dto"
	"currency-conversion/internal/entity"
	mysqlrepo "currency-conversion/internal/repo/mysql"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	gormmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	dialector := gormmysql.New(gormmysql.Config{
		DriverName:                "mysql",
		Conn:                      db,
		SkipInitializeWithVersion: true,
	})

	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	require.NoError(t, err)

	return gormDB, mock
}

func TestGetCurrencies(t *testing.T) {
	db, mock := SetupMockDB(t)

	currencyRepo := mysqlrepo.NewCurrencyRepo(db)

	expectedRows := sqlmock.NewRows([]string{"code", "name"}).
		AddRow("USD", "United States Dollar").
		AddRow("EUR", "Euro")

	mock.ExpectQuery("^SELECT \\* FROM `currencies`").
		WillReturnRows(expectedRows)

	ctx := context.TODO()
	result, err := currencyRepo.GetCurrencies(ctx)

	expectedCurrencies := &dto.Currencies{
		Data: map[string]entity.Currency{
			"USD": {Code: "USD", Name: "United States Dollar"},
			"EUR": {Code: "EUR", Name: "Euro"},
		},
	}

	assert.NoError(t, err)
	assert.Equal(t, expectedCurrencies, result)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestAddCurrencies(t *testing.T) {
	db, mock := SetupMockDB(t)

	currencyRepo := mysqlrepo.NewCurrencyRepo(db)

	currenciesToAdd := &dto.Currencies{
		Data: map[string]entity.Currency{
			"USD": {Code: "USD", Name: "United States Dollar", SymbolNative: "$", DecimalDigits: 2, Active: true, MainAreaId: 1},
			"EUR": {Code: "EUR", Name: "Euro", SymbolNative: "€", DecimalDigits: 2, Active: true, MainAreaId: 2},
		},
	}

	mock.ExpectBegin()

	for _, currency := range currenciesToAdd.Data {
		mock.ExpectExec("^INSERT INTO `currencies`").
			WithArgs(currency.Code, currency.Name, currency.SymbolNative, currency.DecimalDigits, currency.Active, currency.MainAreaId, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(1, 1))
	}

	mock.ExpectCommit()

	ctx := context.TODO()
	err := currencyRepo.AddCurrencies(ctx, currenciesToAdd)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
