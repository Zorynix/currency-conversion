package dto

type AllData struct {
	DataAllCurrencies       DataAllCurrencies       `json:"currentDataAllCurrencies"`
	DataLatestExchangeRates DataLatestExchangeRates `json:"currentDataLatestExchangeRates"`
}
