package models

import "gorm.io/gorm"

type Currency struct {
	gorm.Model
	Code          string `gorm:"serializer:json" json:"code"`
	Name          string `gorm:"serializer:json" json:"name"`
	SymbolNative  string `gorm:"serializer:json" json:"symbol_native"`
	DecimalDigits int    `gorm:"serializer:json" json:"decimal_digits"`
	Active        bool
	MainAreaId    int
}
