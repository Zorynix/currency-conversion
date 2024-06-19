package entity

import (
	"time"

	"gorm.io/gorm"
)

type ExchangeRates struct {
	Code             string `gorm:"primaryKey"`
	CurrencyId       int
	TargetCurrencyId int
	ExchangeRate     float32 `json:"value"`
	RateSourceId     int
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}
