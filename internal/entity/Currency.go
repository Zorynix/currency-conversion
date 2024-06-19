package entity

import (
	"time"

	"gorm.io/gorm"
)

type Currency struct {
	Code          string `gorm:"primaryKey" json:"code"`
	Name          string `json:"name"`
	SymbolNative  string `json:"symbol_native"`
	DecimalDigits int    `json:"decimal_digits"`
	Active        bool
	MainAreaId    int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
