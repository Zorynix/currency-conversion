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

// @TODO: add receiving of a specific currency

func (MSQ *Mysql) GetCurrencies() (*dto.Currencies, error) {
	log.Debug().Msg("Starting GetCurrencies method")
	defer log.Debug().Msg("GetCurrencies method completed")

	var data dto.Currencies

	utils.LoadEnv()
	client := utils.Default()
	client.PrivateToken = os.Getenv("API_KEY")

	res, err := client.FastGet(os.Getenv("url_all_currencies"))
	if err != nil {
		log.Error().Err(err).Msg("Failed to make request to API")
		return nil, fmt.Errorf("error when creating a request: %s", err)
	}

	if err := json.Unmarshal(res.Body(), &data); err != nil {
		log.Error().Err(err).Msg("JSON parsing failed")
		return nil, fmt.Errorf("error during JSON parsing: %s", err)
	}

	log.Debug().Interface("CurrenciesData", data).Msg("Successfully fetched and parsed currencies data")
	return &data, nil
}

func (MSQ *Mysql) InsertCurrencies() (*dto.Currencies, error) {
	log.Debug().Msg("Starting InsertCurrencies method")
	defer log.Debug().Msg("InsertCurrencies method completed")

	var CurrenciesData []models.Currency

	data, err := MSQ.GetCurrencies()
	if err != nil {
		log.Error().Err(err).Msg("Failed to fetch currencies")
		return nil, fmt.Errorf("error fetching currencies: %s", err)
	}

	for _, value := range data.Data {
		CurrenciesData = append(CurrenciesData, value)
	}

	log.Debug().Interface("CurrenciesToInsert", CurrenciesData).Msg("Ready to insert currencies data")
	MSQ.DB.Save(&CurrenciesData)
	log.Info().Msg("Currencies data inserted into database successfully")

	return data, nil
}

func (MSQ *Mysql) GetExchangeRates() (*dto.ExchangeRates, error) {
	log.Debug().Msg("Starting GetExchangeRates method")
	defer log.Debug().Msg("GetExchangeRates method completed")

	log.Info().Msg("GetExchangeRates called")

	var data dto.ExchangeRates

	utils.LoadEnv()
	client := utils.Default()
	client.PrivateToken = os.Getenv("API_KEY")

	res, err := client.FastGet(os.Getenv("url_latest_exchange_rates"))
	if err != nil {
		log.Error().Err(err).Msg("Failed to make request to API")
		return nil, fmt.Errorf("error when creating a request: %s", err)
	}

	if err := json.Unmarshal(res.Body(), &data); err != nil {
		log.Error().Err(err).Msg("JSON parsing failed")
		return nil, fmt.Errorf("error during JSON parsing: %s", err)
	}

	log.Debug().Interface("ExchangeRatesData", data).Msg("Successfully fetched and parsed exchange rates data")
	return &data, nil
}

func (MSQ *Mysql) InsertExchangeRates() (*dto.ExchangeRates, error) {
	log.Debug().Msg("Starting InsertExchangeRates method")
	defer log.Debug().Msg("InsertExchangeRates method completed")

	log.Info().Msg("InsertExchangeRates called")

	var ExchangeRateData []models.ExchangeRates

	data, err := MSQ.GetExchangeRates()
	if err != nil {
		log.Error().Err(err).Msg("Failed to fetch exchange rates")
		return nil, fmt.Errorf("error fetching exchange rates: %s", err)
	}

	for _, value := range data.Data {
		ExchangeRateData = append(ExchangeRateData, value)
	}

	log.Debug().Interface("ExchangeRatesToInsert", ExchangeRateData).Msg("Ready to insert exchange rates data")
	MSQ.DB.Save(&ExchangeRateData)
	log.Info().Msg("Exchange rates data inserted into database successfully")

	return data, nil
}

//@TODO: sheduled updates

func (MSQ *Mysql) UpdateRates() (*dto.ExchangeRateHistory, error) {
	log.Debug().Msg("Starting UpdateRates method")
	defer log.Debug().Msg("UpdateRates method completed")

	log.Info().Msg("UpdateRates called")

	var count int64
	var data dto.ExchangeRateHistory

	if err := MSQ.DB.Model(&models.ExchangeRateHistory{}).Count(&count).Error; err != nil {
		log.Error().Err(err).Msg("Failed to count exchange rate histories")
		return nil, fmt.Errorf("error counting exchange rate histories: %s", err)
	}

	if count == 0 {
		log.Debug().Msg("No exchange rate histories found, inserting initial data from exchange rates")
		if err := MSQ.DB.Exec("INSERT INTO exchange_rate_histories SELECT * FROM exchange_rates").Error; err != nil {
			log.Error().Err(err).Msg("Failed to insert initial exchange rate histories")
			return nil, fmt.Errorf("error inserting initial exchange rate histories: %s", err)
		}
		log.Info().Msg("Initial exchange rate histories inserted successfully")
	} else {
		log.Debug().Msg("Exchange rate histories found, updating with latest exchange rates")
		if err := MSQ.DB.Exec("REPLACE INTO exchange_rate_histories SELECT * FROM exchange_rates").Error; err != nil {
			log.Error().Err(err).Msg("Failed to update exchange rate histories")
			return nil, fmt.Errorf("error updating exchange rate histories: %s", err)
		}
		log.Info().Msg("Exchange rate histories updated successfully")
	}

	if _, err := MSQ.InsertExchangeRates(); err != nil {
		log.Error().Err(err).Msg("Failed to insert latest exchange rates")
		return nil, fmt.Errorf("error inserting latest exchange rates: %s", err)
	}

	return &data, nil
}

// @TODO: add event to queue (rabbitmq)
