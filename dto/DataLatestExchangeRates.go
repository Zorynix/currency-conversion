package dto

import (
	"currency-conversion/models"

	"gorm.io/gorm"
)

type DataLatestExchangeRates struct {
	gorm.Model
	Data map[string]models.CurrenciesExchangeRates `gorm:"serializer:json" json:"data"`
}
