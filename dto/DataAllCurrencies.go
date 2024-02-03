package dto

import (
	"currency-conversion/models"

	"gorm.io/gorm"
)

type DataAllCurrencies struct {
	gorm.Model
	Data map[string]models.Currency `gorm:"serializer:json" json:"data"`
}
