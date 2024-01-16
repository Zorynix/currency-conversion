package dto

import "currency-conversion/models"

type DataLatestExchangeRates struct {
	Data map[string]models.CurrenciesExchangeRates `json:"data"`
}
