package models

import (
	"time"

	"gorm.io/gorm"
)

type ExchangeRateHistory struct {
	Code             string `gorm:"primaryKey"`
	CurrencyId       int
	TargetCurrencyId int
	ExchangeRate     float32
	RateSourceId     int
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}
