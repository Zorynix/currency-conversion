package services

import (
	"context"
	"currency-conversion/internal/dto"
	"currency-conversion/internal/repo"
)

type RatesService interface {
	GetCurrencies(ctx context.Context) (*dto.Currencies, error)
	AddCurrencies(ctx context.Context) (*dto.Currencies, error)
	GetExchangeRates(ctx context.Context) (*dto.ExchangeRates, error)
	AddRates(ctx context.Context) (*dto.ExchangeRates, error)
	UpdateRates(ctx context.Context) (string, error)
}

type Services struct {
	Rates RatesService
}

type ServicesDependencies struct {
	DB                Database
	CurrencyRepo      repo.Currency
	ExchangeRatesRepo repo.ExchangeRates
	RateHistoriesRepo repo.RateHistories
}

func NewServices(deps ServicesDependencies) *Services {
	return &Services{
		Rates: NewRatesService(deps.DB, deps.CurrencyRepo, deps.ExchangeRatesRepo, deps.RateHistoriesRepo),
	}
}
