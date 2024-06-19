package dto

import (
	"currency-conversion/internal/entity"

	"gorm.io/gorm"
)

type Currencies struct {
	gorm.Model
	Data map[string]entity.Currency `gorm:"serializer:json" json:"data"`
}
