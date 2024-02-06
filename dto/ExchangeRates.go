package dto

import (
	"currency-conversion/models"

	"gorm.io/gorm"
)

type ExchangeRates struct {
	gorm.Model
	Data map[string]models.ExchangeRates `gorm:"serializer:json" json:"data"`
}
