package dto

import "gorm.io/gorm"

type AllData struct {
	gorm.Model
	DataAllCurrencies       DataAllCurrencies       `gorm:"serializer:json" json:"DataAllCurrencies"`
	DataLatestExchangeRates DataLatestExchangeRates `gorm:"serializer:json" json:"DataLatestExchangeRates"`
}
