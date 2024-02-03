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

	data, err := MSQ.AllCurrencies()
	if err != nil {
		return nil, fmt.Errorf("Error when writing data")
	}

	return data, nil
}

func (MSQ *Mysql) InsertLatestExchangeRates() (*dto.DataLatestExchangeRates, error) {

	data, err := MSQ.LatestExchangeRates()
	if err != nil {
		return nil, fmt.Errorf("Error when writing data")
	}

	return data, nil
}

func (MSQ *Mysql) AllCurrencies() (*dto.DataAllCurrencies, error) {

	var data dto.DataAllCurrencies
	var allCurrenciesData []models.Currency

	utils.LoadEnv()
	client := utils.Default()
	client.PrivateToken = os.Getenv("API_KEY")

	res, err := client.FastGet(os.Getenv("URL_all_currencies"))
	if err != nil {
		return nil, fmt.Errorf("Error when creating a request: %s", err)
	}

	if err := json.Unmarshal(res.Body(), &data); err != nil {
		return nil, fmt.Errorf("Error during JSON parsing: %s", err)
	}

	for _, value := range data.Data {
		allCurrenciesData = append(allCurrenciesData, value)
	}

	MSQ.db.Save(&allCurrenciesData)

	return &data, nil
}

func (MSQ *Mysql) LatestExchangeRates() (*dto.DataLatestExchangeRates, error) {

	var data dto.DataLatestExchangeRates
	var LatestExchangeData []models.CurrenciesExchangeRates

	utils.LoadEnv()
	client := utils.Default()
	client.PrivateToken = os.Getenv("API_KEY")

	res, err := client.FastGet(os.Getenv("URL_latest_exchange_rates"))
	if err != nil {
		return nil, fmt.Errorf("Error when creating a request: %s", err)
	}

	if err := json.Unmarshal(res.Body(), &data); err != nil {
		return nil, fmt.Errorf("Error during JSON parsing: %s", err)
	}

	for _, value := range data.Data {
		LatestExchangeData = append(LatestExchangeData, value)
	}

	MSQ.db.Save(&LatestExchangeData)

	return &data, nil
}

func updateRates() {
	// Currencyapi Service # hourly
	// ECB Service # daily

	// rates Service interface

	// currency_codes := [180]string{"USD", "CAD", "CNY"}

	// for _, currencyCode := range currency_codes {
	// 	go updateRate
	// }
}

func updateRate( /* currencyCode */ ) {
	// get latest from API
	// update to DB
	// --- resave to history table

	// @TODO: add event to queue (rabbitmq)
}
