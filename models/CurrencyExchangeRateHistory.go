package models

import "time"

type CurrencyExchangeRateHistory struct {
	Id              int
	CurrencyId      int
	TargetCurencyId int
	ExchangeRate    float32 `json:"value"`
	RateSourceId    int
	UpdateDate      time.Time `json:"last_updated_at"`
}
