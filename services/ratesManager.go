package services

import (
	"currency-conversion/dto"
	"currency-conversion/models"
	"currency-conversion/utils"
	"encoding/json"
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm/clause"
)

// @TODO: add receiving of a specific currency

func (MSQ *Mysql) GetCurrencies() (*dto.Currencies, error) {
	log.Info().Msg("GetCurrencies called")
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

func (MSQ *Mysql) AddCurrencies() (*dto.Currencies, error) {
	log.Info().Msg("AddCurrencies called")
	log.Debug().Msg("Starting AddCurrencies method")
	defer log.Debug().Msg("AddCurrencies method completed")

	data, err := MSQ.GetCurrencies()
	if err != nil {
		log.Error().Err(err).Msg("Failed to fetch currencies")
		return nil, fmt.Errorf("error fetching currencies: %s", err)
	}

	tx := MSQ.DB.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	for _, value := range data.Data {
		if err := tx.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(&value).Error; err != nil {
			tx.Rollback()
			log.Error().Err(err).Msg("Failed to add or update currency")
			return nil, fmt.Errorf("error adding or updating currency: %s", err)
		}
	}
	tx.Commit()
	log.Info().Msg("Currencies data inserted or updated in database successfully")

	return data, nil
}

func (MSQ *Mysql) GetExchangeRates() (*dto.ExchangeRates, error) {
	log.Info().Msg("GetExchangeRates called")
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

func (MSQ *Mysql) AddRates() (*dto.ExchangeRates, error) {
	log.Info().Msg("AddRates called")
	log.Debug().Msg("Starting AddRates method")
	defer log.Debug().Msg("AddRates method completed")

	data, err := MSQ.GetExchangeRates()
	if err != nil {
		log.Error().Err(err).Msg("Failed to fetch exchange rates")
		return nil, fmt.Errorf("error fetching exchange rates: %s", err)
	}

	tx := MSQ.DB.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	var count int64
	if err := tx.Model(&models.ExchangeRates{}).Count(&count).Error; err != nil {
		tx.Rollback()
		log.Error().Err(err).Msg("Failed to count exchange rates")
		return nil, fmt.Errorf("error counting exchange rates: %s", err)
	}

	if count == 0 {
		for _, rate := range data.Data {
			if err := tx.Create(&rate).Error; err != nil {
				tx.Rollback()
				log.Error().Err(err).Msg("Failed to add exchange rates")
				return nil, fmt.Errorf("error adding exchange rates: %s", err)
			}
		}
	} else {
		for _, rate := range data.Data {
			if err := tx.Model(&models.ExchangeRates{}).Where("code = ?", rate.Code).Updates(rate).Error; err != nil {
				tx.Rollback()
				log.Error().Err(err).Msg("Failed to update exchange rates")
				return nil, fmt.Errorf("error updating exchange rates: %s", err)
			}
		}
	}

	tx.Commit()
	log.Info().Msg("Exchange rates data inserted/updated into database successfully")
	return data, nil
}

//@TODO: sheduled updates

func (MSQ *Mysql) UpdateRates() (*dto.ExchangeRates, error) {
	log.Info().Msg("UpdateRates called")
	log.Debug().Msg("Starting UpdateRates method")
	defer log.Debug().Msg("UpdateRates method completed")

	var count int64
	var data dto.ExchangeRates

	tx := MSQ.DB.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	if err := tx.Model(&models.ExchangeRates{}).Count(&count).Error; err != nil {
		tx.Rollback()
		log.Error().Err(err).Msg("Failed to count exchange rate histories")
		return nil, fmt.Errorf("error counting exchange rate histories: %s", err)
	}

	if count == 0 {
		log.Debug().Msg("No exchange rate histories found, inserting initial data from exchange rates")
		if err := tx.Exec("INSERT INTO exchange_rate_histories SELECT * FROM exchange_rates").Error; err != nil {
			tx.Rollback()
			log.Error().Err(err).Msg("Failed to insert initial exchange rate histories")
			return nil, fmt.Errorf("error inserting initial exchange rate histories: %s", err)
		}
		log.Info().Msg("Initial exchange rate histories inserted successfully")
	} else {
		log.Debug().Msg("Exchange rate histories found, updating with latest exchange rates")
		if err := tx.Exec(`INSERT INTO exchange_rate_histories(code, currency_id, target_currency_id, exchange_rate, rate_source_id, updated_at)
        SELECT code, currency_id, target_currency_id, exchange_rate, rate_source_id, updated_at
        FROM exchange_rates
        ON DUPLICATE KEY UPDATE
			code = VALUES(code),
            exchange_rate = VALUES(exchange_rate),
            rate_source_id = VALUES(rate_source_id),
            updated_at = VALUES(updated_at)
    `).Error; err != nil {
			tx.Rollback()
			log.Error().Err(err).Msg("Failed to update exchange rate histories")
			return nil, fmt.Errorf("error updating exchange rate histories: %s", err)
		}
		log.Info().Msg("Exchange rate histories updated successfully")
	}

	tx.Commit()

	if _, err := MSQ.AddRates(); err != nil {
		tx.Rollback()
		log.Error().Err(err).Msg("Failed to insert latest exchange rates")
		return nil, fmt.Errorf("error inserting latest exchange rates: %s", err)
	}

	return &data, nil
}

// @TODO: add event to queue (rabbitmq)
