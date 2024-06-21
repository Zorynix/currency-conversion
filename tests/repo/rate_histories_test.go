package repo_test

import (
	"context"
	"testing"

	mysqlrepo "currency-conversion/internal/repo/mysql"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestUpdateRatesHistories(t *testing.T) {
	db, mock := SetupMockDB(t)
	rateHistoriesRepo := mysqlrepo.NewRateHistoriesRepo(db)

	// таблица exchange_rates пуста
	mock.ExpectBegin()
	mock.ExpectQuery("^SELECT count\\(\\*\\) FROM `exchange_rates`").
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))
	mock.ExpectExec("INSERT INTO exchange_rate_histories SELECT \\* FROM exchange_rates").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	ctx := context.TODO()
	err := rateHistoriesRepo.UpdateRatesHistories(ctx)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())

	// таблица exchange_rates содержит записи
	mock.ExpectBegin()
	mock.ExpectQuery("^SELECT count\\(\\*\\) FROM `exchange_rates`").
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(2))
	mock.ExpectExec(`INSERT INTO exchange_rate_histories\(code, currency_id, target_currency_id, exchange_rate, rate_source_id, updated_at\) SELECT code, currency_id, target_currency_id, exchange_rate, rate_source_id, updated_at FROM exchange_rates ON DUPLICATE KEY UPDATE code = VALUES\(code\), exchange_rate = VALUES\(exchange_rate\), rate_source_id = VALUES\(rate_source_id\), updated_at = VALUES\(updated_at\)`).
		WillReturnResult(sqlmock.NewResult(1, 2))
	mock.ExpectCommit()

	err = rateHistoriesRepo.UpdateRatesHistories(ctx)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
