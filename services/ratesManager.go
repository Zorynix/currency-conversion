package services

import (
	"currency-conversion/dto"
	"currency-conversion/models"
	"currency-conversion/utils"
	"encoding/json"
	"fmt"
	"os"
)

func (MSQ *Mysql) InsertAllCurrencies() (*dto.DataAllCurrencies, error) {

	var allCurrenciesData []models.Currency

	data, err := MSQ.AllCurrencies()
	if err != nil {
		return nil, fmt.Errorf("error when writing data")
	}

	for _, value := range data.Data {
		allCurrenciesData = append(allCurrenciesData, value)
	}

	MSQ.db.Save(&allCurrenciesData)

	return data, nil
}

func (MSQ *Mysql) InsertLatestExchangeRates() (*dto.DataLatestExchangeRates, error) {

	var LatestExchangeData []models.CurrenciesExchangeRates

	data, err := MSQ.LatestExchangeRates()
	if err != nil {
		return nil, fmt.Errorf("error when writing data")
	}

	for _, value := range data.Data {
		LatestExchangeData = append(LatestExchangeData, value)
	}

	MSQ.db.Save(&LatestExchangeData)

	return data, nil
}

func (MSQ *Mysql) AllCurrencies() (*dto.DataAllCurrencies, error) {

	var data dto.DataAllCurrencies

	utils.LoadEnv()
	client := utils.Default()
	client.PrivateToken = os.Getenv("API_KEY")

	res, err := client.FastGet(os.Getenv("URL_all_currencies"))
	if err != nil {
		return nil, fmt.Errorf("error when creating a request: %s", err)
	}

	if err := json.Unmarshal(res.Body(), &data); err != nil {
		return nil, fmt.Errorf("error during JSON parsing: %s", err)
	}

	return &data, nil
}

func (MSQ *Mysql) LatestExchangeRates() (*dto.DataLatestExchangeRates, error) {

	var data dto.DataLatestExchangeRates

	utils.LoadEnv()
	client := utils.Default()
	client.PrivateToken = os.Getenv("API_KEY")

	res, err := client.FastGet(os.Getenv("URL_latest_exchange_rates"))
	if err != nil {
		return nil, fmt.Errorf("error when creating a request: %s", err)
	}

	if err := json.Unmarshal(res.Body(), &data); err != nil {
		return nil, fmt.Errorf("error during JSON parsing: %s", err)
	}

	return &data, nil
}

// @TODO: sheduled updates
func (MSQ *Mysql) updateRates( /* currencyCode */ ) {
	// Currencyapi Service # hourly
	// ECB Service # daily

	// rates Service interface

	// currency_codes := [180]string{"USD", "CAD", "CNY"}

	// for _, currencyCode := range currency_codes {
	// 	go updateRate
	// }
}

func (MSQ *Mysql) updateRate( /* currencyCode */ ) {
	// get latest from API
	// update to DB
	// --- resave to history table

	// @TODO: add event to queue (rabbitmq)
}
