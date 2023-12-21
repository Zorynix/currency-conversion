package models

type CurrencyExchangeRateHistory struct {
	Id              int
	CurrencyId      int
	TargetCurencyId int
	ExchangeRate    string `json:"value"`
	RateSourceId    string `json:""`
	UpdateDate      string `json:"last_updated_at"`
}
