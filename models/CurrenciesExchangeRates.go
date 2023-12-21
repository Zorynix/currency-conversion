package models

type CurrenciesExchangeRates struct {
	Id              int
	CurrencyId      int
	TargetCurencyId int
	ExchangeRate    string `json:"value"`
	RateSourceId    string `json:""`
	CreatedAt       string `json:""`
	UpdatedAt       string `json:""`
}
