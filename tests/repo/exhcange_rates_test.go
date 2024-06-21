package repo_test

import (
	"context"
	"testing"
	"time"

	"currency-conversion/internal/dto"
	"currency-conversion/internal/entity"
	mysqlrepo "currency-conversion/internal/repo/mysql"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetExchangeRates(t *testing.T) {
	db, mock := SetupMockDB(t)

	exchangeRatesRepo := mysqlrepo.NewExchangeRatesRepo(db)
	now := time.Now()
	expectedRows := sqlmock.NewRows([]string{"code", "currency_id", "target_currency_id", "exchange_rate", "rate_source_id", "created_at", "updated_at", "deleted_at"}).
		AddRow("USD_EUR", 1, 2, 0.85, 1, now, now, nil).
		AddRow("EUR_GBP", 2, 3, 0.75, 1, now, now, nil)

	mock.ExpectQuery("^SELECT \\* FROM `exchange_rates` WHERE `exchange_rates`.`deleted_at` IS NULL").
		WillReturnRows(expectedRows)

	ctx := context.TODO()
	result, err := exchangeRatesRepo.GetExchangeRates(ctx)

	expectedExchangeRates := &dto.ExchangeRates{
		Data: map[string]entity.ExchangeRates{
			"USD_EUR": {Code: "USD_EUR", CurrencyId: 1, TargetCurrencyId: 2, ExchangeRate: 0.85, RateSourceId: 1, CreatedAt: now, UpdatedAt: now},
			"EUR_GBP": {Code: "EUR_GBP", CurrencyId: 2, TargetCurrencyId: 3, ExchangeRate: 0.75, RateSourceId: 1, CreatedAt: now, UpdatedAt: now},
		},
	}

	for code, expectedRate := range expectedExchangeRates.Data {
		actualRate, exists := result.Data[code]
		require.True(t, exists)
		assert.Equal(t, expectedRate.Code, actualRate.Code)
		assert.Equal(t, expectedRate.CurrencyId, actualRate.CurrencyId)
		assert.Equal(t, expectedRate.TargetCurrencyId, actualRate.TargetCurrencyId)
		assert.Equal(t, expectedRate.ExchangeRate, actualRate.ExchangeRate)
		assert.Equal(t, expectedRate.RateSourceId, actualRate.RateSourceId)
	}

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestAddRates(t *testing.T) {
	db, mock := SetupMockDB(t)

	exchangeRatesRepo := mysqlrepo.NewExchangeRatesRepo(db)

	ratesToAdd := &dto.ExchangeRates{
		Data: map[string]entity.ExchangeRates{
			"USD_EUR": {Code: "USD_EUR", CurrencyId: 1, TargetCurrencyId: 2, ExchangeRate: 0.85, RateSourceId: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			"EUR_GBP": {Code: "EUR_GBP", CurrencyId: 2, TargetCurrencyId: 3, ExchangeRate: 0.75, RateSourceId: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		},
	}

	mock.ExpectBegin()

	for _, rate := range ratesToAdd.Data {
		mock.ExpectExec("^INSERT INTO `exchange_rates`").
			WithArgs(rate.Code, rate.CurrencyId, rate.TargetCurrencyId, rate.ExchangeRate, rate.RateSourceId, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(1, 1))
	}

	mock.ExpectCommit()

	ctx := context.TODO()
	err := exchangeRatesRepo.AddRates(ctx, ratesToAdd)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
