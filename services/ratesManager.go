package services

import (
	"currency-conversion/dto"
	"currency-conversion/models"
	"currency-conversion/utils"
	"encoding/json"
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
)

// @TODO: сделать получение конкретной валюты + вставка в history и обновление текущих

func (MSQ *Mysql) GetCurrencies() (*dto.Currencies, error) {

	var data dto.Currencies

	utils.LoadEnv()
	client := utils.Default()
	client.PrivateToken = os.Getenv("API_KEY")

	res, err := client.FastGet(os.Getenv("URL_currencies"))
	if err != nil {
		return nil, fmt.Errorf("error when creating a request: %s", err)
	}

	if err := json.Unmarshal(res.Body(), &data); err != nil {
		return nil, fmt.Errorf("error during JSON parsing: %s", err)
	}

	return &data, nil
}

func (MSQ *Mysql) InsertCurrencies() (*dto.Currencies, error) {

	var CurrenciesData []models.Currency

	data, err := MSQ.GetCurrencies()
	if err != nil {
		return nil, fmt.Errorf("error inserting currencies")
	}

	for _, value := range data.Data {
		CurrenciesData = append(CurrenciesData, value)
	}

	MSQ.db.Save(&CurrenciesData)

	return data, nil
}

func (MSQ *Mysql) GetExchangeRates() (*dto.ExchangeRates, error) {

	var data dto.ExchangeRates

	utils.LoadEnv()
	client := utils.Default()
	client.PrivateToken = os.Getenv("API_KEY")

	res, err := client.FastGet(os.Getenv("URL_exchange_rates"))
	if err != nil {
		return nil, fmt.Errorf("error when creating a request: %s", err)
	}

	if err := json.Unmarshal(res.Body(), &data); err != nil {
		return nil, fmt.Errorf("error during JSON parsing: %s", err)
	}

	return &data, nil
}

func (MSQ *Mysql) InsertExchangeRates() (*dto.ExchangeRates, error) {

	var ExchangeRateData []models.ExchangeRates

	data, err := MSQ.GetExchangeRates()
	if err != nil {
		return nil, fmt.Errorf("error inserting exchange rates")
	}

	for _, value := range data.Data {
		ExchangeRateData = append(ExchangeRateData, value)
	}

	//MSQ.db.Model(&models.ExchangeRates{}).Updates(ExchangeRateData)
	MSQ.db.Save(&ExchangeRateData)

	return data, nil

}

//@TODO: sheduled updates

func (MSQ *Mysql) UpdateRates() (*dto.ExchangeRateHistory, error) {

	var count int64

	var data dto.ExchangeRateHistory

	if err := MSQ.db.Model(&models.ExchangeRateHistory{}).Count(&count).Error; err != nil {
		log.Fatal()
	}

	if count == 0 {
		MSQ.db.Exec("INSERT INTO exchange_rate_histories SELECT * FROM exchange_rates")
		fmt.Println("insert successful")
	} else {
		MSQ.db.Exec("REPLACE INTO exchange_rate_histories SELECT * FROM exchange_rates")
		fmt.Println("update successful")
	}

	if _, err := MSQ.InsertExchangeRates(); err != nil {
		return nil, fmt.Errorf("error updating exchange rates: %s", err)
	}

	// if err := MSQ.db.Select("*").Find(&data).Error; err != nil {
	// 	return nil, fmt.Errorf("error fetching exchange rate history: %s", err)
	// }

	// fmt.Println(data)

	return &data, nil
}

// @TODO: add event to queue (rabbitmq)
