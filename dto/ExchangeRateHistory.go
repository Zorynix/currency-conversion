package dto

import (
	"currency-conversion/models"

	"gorm.io/gorm"
)

type ExchangeRateHistory struct {
	gorm.Model
	Data map[string]models.ExchangeRateHistory `gorm:"serializer:json" json:"data"`
}
