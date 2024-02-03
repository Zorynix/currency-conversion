package models

import (
	"gorm.io/gorm"
)

type CurrenciesExchangeRates struct {
	gorm.Model
	CurrencyId       int
	TargetCurrencyId int
	ExchangeRate     float32 `gorm:"serializer:json" json:"value"`
	RateSourceId     int
}
