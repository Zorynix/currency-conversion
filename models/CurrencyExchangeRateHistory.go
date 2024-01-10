package models

import "time"

type CurrencyExchangeRateHistory struct {
	Id              int
	CurrencyId      int
	TargetCurencyId int
	ExchangeRate    float32
	RateSourceId    int
	UpdateDate      time.Time
}
