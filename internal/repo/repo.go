package repo

import (
	"context"
	"currency-conversion/internal/dto"
	"currency-conversion/internal/repo/mysql"

	"gorm.io/gorm"
)

type Currency interface {
	GetCurrencies(ctx context.Context) (*dto.Currencies, error)
	AddCurrencies(ctx context.Context, data *dto.Currencies) error
}

type ExchangeRates interface {
	GetExchangeRates(ctx context.Context) (*dto.ExchangeRates, error)
	AddRates(ctx context.Context, data *dto.ExchangeRates) error
}

type RateHistories interface {
	UpdateRatesHistories(ctx context.Context) error
}

type Repositories struct {
	Currency
	ExchangeRates
	RateHistories
}

// NewRepositories создает новые репозитории и возвращает указатель на структуру Repositories
func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Currency:      mysql.NewCurrencyRepo(db),
		ExchangeRates: mysql.NewExchangeRatesRepo(db),
		RateHistories: mysql.NewRateHistoriesRepo(db),
	}
}
