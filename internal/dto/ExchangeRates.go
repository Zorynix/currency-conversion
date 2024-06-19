package dto

import (
	"currency-conversion/internal/entity"

	"gorm.io/gorm"
)

type ExchangeRates struct {
	gorm.Model
	Data map[string]entity.ExchangeRates `gorm:"serializer:json" json:"data"`
}
