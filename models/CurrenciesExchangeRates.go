package models

import (
	"time"

	"gorm.io/gorm"
)

//структура структур

type CurrenciesExchangeRates struct {
	gorm.Model
	CurrencyId      int
	TargetCurencyId int
	ExchangeRate    float32 `gorm:"serializer:json" json:"value"`
	RateSourceId    int
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
