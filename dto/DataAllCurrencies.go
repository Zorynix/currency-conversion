package dto

import "currency-conversion/models"

type DataAllCurrencies struct {
	Data map[string]models.Currency `json:"data"`
}
