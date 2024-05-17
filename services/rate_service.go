package services

import (
	"currency-conversion/dto"
	"currency-conversion/repo"
	"currency-conversion/utils"
	"encoding/json"
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
)

type RatesService interface {
	GetCurrencies() (*dto.Currencies, error)
	AddCurrencies() (*dto.Currencies, error)
	GetExchangeRates() (*dto.ExchangeRates, error)
	AddRates() (*dto.ExchangeRates, error)
	UpdateRates() (string, error)
}

type ratesService struct {
	db                Database
	currencyRepo      repo.CurrencyRepo
	exchangeRatesRepo repo.ExchangeRatesRepo
	rateHistoriesRepo repo.RateHistoriesRepo
}

func NewRatesService(db Database, currencyRepo repo.CurrencyRepo, exchangeRatesRepo repo.ExchangeRatesRepo, rateHistoriesRepo repo.RateHistoriesRepo) RatesService {
	return &ratesService{
		db:                db,
		currencyRepo:      currencyRepo,
		exchangeRatesRepo: exchangeRatesRepo,
		rateHistoriesRepo: rateHistoriesRepo,
	}
}

func (s *ratesService) GetCurrencies() (*dto.Currencies, error) {
	log.Info().Msg("GetCurrencies called")
	utils.LoadEnv()
	client := utils.Default()
	client.PrivateToken = os.Getenv("API_KEY")

	res, err := client.FastGet(os.Getenv("url_all_currencies"))
	if err != nil {
		log.Error().Err(err).Msg("Failed to make request to API")
		return nil, fmt.Errorf("error when creating a request: %s", err)
	}

	var data dto.Currencies
	if err := json.Unmarshal(res.Body(), &data); err != nil {
		log.Error().Err(err).Msg("JSON parsing failed")
		return nil, fmt.Errorf("error during JSON parsing: %s", err)
	}

	log.Debug().Interface("CurrenciesData", data).Msg("Successfully fetched and parsed currencies data")
	return &data, nil
}

func (s *ratesService) AddCurrencies() (*dto.Currencies, error) {
	log.Info().Msg("AddCurrencies called")
	data, err := s.GetCurrencies()
	if err != nil {
		log.Error().Err(err).Msg("Failed to fetch currencies")
		return nil, fmt.Errorf("error fetching currencies: %s", err)
	}

	if err := s.currencyRepo.AddCurrencies(data); err != nil {
		log.Error().Err(err).Msg("Failed to add or update currency")
		return nil, fmt.Errorf("error adding or updating currency: %s", err)
	}

	log.Info().Msg("Currencies data added or updated successfully")
	return data, nil
}

func (s *ratesService) GetExchangeRates() (*dto.ExchangeRates, error) {
	log.Info().Msg("GetExchangeRates called")
	utils.LoadEnv()
	client := utils.Default()
	client.PrivateToken = os.Getenv("API_KEY")

	res, err := client.FastGet(os.Getenv("url_latest_exchange_rates"))
	if err != nil {
		log.Error().Err(err).Msg("Failed to make request to API")
		return nil, fmt.Errorf("error when creating a request: %s", err)
	}

	var data dto.ExchangeRates
	if err := json.Unmarshal(res.Body(), &data); err != nil {
		log.Error().Err(err).Msg("JSON parsing failed")
		return nil, fmt.Errorf("error during JSON parsing: %s", err)
	}

	log.Debug().Interface("ExchangeRatesData", data).Msg("Successfully fetched and parsed exchange rates data")
	return &data, nil
}

func (s *ratesService) AddRates() (*dto.ExchangeRates, error) {
	log.Info().Msg("AddRates called")
	data, err := s.GetExchangeRates()
	if err != nil {
		log.Error().Err(err).Msg("Failed to fetch exchange rates")
		return nil, fmt.Errorf("error fetching exchange rates: %s", err)
	}

	if err := s.exchangeRatesRepo.AddRates(data); err != nil {
		log.Error().Err(err).Msg("Failed to add or update exchange rates")
		return nil, fmt.Errorf("error adding or updating exchange rates: %s", err)
	}

	log.Info().Msg("Exchange rates data added or updated successfully")
	return data, nil
}

func (s *ratesService) UpdateRates() (string, error) {
	log.Info().Msg("UpdateRates called")
	if err := s.rateHistoriesRepo.UpdateRatesHistories(); err != nil {
		log.Error().Err(err).Msg("Failed to update exchange rate histories")
		return "", fmt.Errorf("error updating exchange rate histories: %s", err)
	}

	if _, err := s.AddRates(); err != nil {
		log.Error().Err(err).Msg("Failed to insert latest exchange rates")
		return "", fmt.Errorf("error inserting latest exchange rates: %s", err)
	}

	log.Info().Msg("Exchange rates updated successfully")
	return "Exchange rates updated successfully", nil
}
