package models

import (
	"gorm.io/gorm"
)

type CurrencyExchangeRateHistory struct {
	gorm.Model
	CurrencyId       int
	TargetCurrencyId int
	ExchangeRate     float32
	RateSourceId     int
}
