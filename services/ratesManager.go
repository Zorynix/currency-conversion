package services

import (
	"currency-conversion/dto"
	"currency-conversion/utils"
	"encoding/json"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Panic().Msg("No .env file found")
	}
}

func (MSQ *Mysql) TestApiAllCurrencies() (*dto.DataAllCurrencies, error) {
	data, err := MSQ.AllCurrencies()
	if err != nil {
		return nil, fmt.Errorf("Error when writing data")
	}

	return data, nil
}

func (MSQ *Mysql) TestApiLatestExchangeRates() (*dto.DataLatestExchangeRates, error) {

	data, err := MSQ.LatestExchangeRates()
	if err != nil {
		return nil, fmt.Errorf("Error when writing data")
	}

	return data, nil
}

func (MSQ *Mysql) AllCurrencies() (*dto.DataAllCurrencies, error) {

	client := utils.Default()
	client.PrivateToken = os.Getenv("API_KEY")
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error when uploading a file .env")
		panic(err)
	}

	res, err := client.FastGet(os.Getenv("URL_all_currencies"))
	if err != nil {
		return nil, fmt.Errorf("Error when creating a request: %s", err)
	}

	var data dto.DataAllCurrencies
	if err := json.Unmarshal(res.Body(), &data); err != nil {
		return nil, fmt.Errorf("Error during JSON parsing: %s", err)
	}

	//MSQ.db.Save(models.CurrenciesExchangeRates{Id: 300523, CurrencyId: 1001, TargetCurencyId: 9999, ExchangeRate: 30001, RateSourceId: 99, CreatedAt: time.Now(), UpdatedAt: date})
	return &data, nil
}

func (MSQ *Mysql) LatestExchangeRates() (*dto.DataLatestExchangeRates, error) {
	client := utils.Default()
	client.PrivateToken = os.Getenv("API_KEY")

	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error when uploading a file .env")
		panic(err)
	}

	res, err := client.FastGet(os.Getenv("URL_latest_exchange_rates"))
	if err != nil {
		return nil, fmt.Errorf("Error when creating a request: %s", err)
	}

	var data dto.DataLatestExchangeRates
	if err := json.Unmarshal(res.Body(), &data); err != nil {
		return nil, fmt.Errorf("Error during JSON parsing: %s", err)
	}

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(res.Body()))
	return &data, nil
}

func (MSQ *Mysql) TestInsert() error {

	return nil
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
