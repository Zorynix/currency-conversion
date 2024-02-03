package models

import (
	"time"

	"gorm.io/gorm"
)

type CurrencyExchangeRateHistory struct {
	gorm.Model
	CurrencyId      int
	TargetCurencyId int
	ExchangeRate    float32
	RateSourceId    int
	UpdateDate      time.Time
}
