package services

import (
	"currency-conversion/config"
	"currency-conversion/dto"
	"currency-conversion/repo"
	"currency-conversion/utils"
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
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
	logrus.Info("GetCurrencies called")
	client := utils.Default()
	client.PrivateToken = config.Cfg.APIKey

	res, err := client.FastGet(config.Cfg.URLs.AllCurrencies)
	if err != nil {
		logrus.Errorf("Failed to make request to API: %v", err)
		return nil, fmt.Errorf("error when creating a request: %s", err)
	}

	var data dto.Currencies
	if err := json.Unmarshal(res.Body(), &data); err != nil {
		logrus.Errorf("JSON parsing failed: %v", err)
		return nil, fmt.Errorf("error during JSON parsing: %s", err)
	}

	logrus.Debugf("Successfully fetched and parsed currencies data: %v", data)
	return &data, nil
}

func (s *ratesService) AddCurrencies() (*dto.Currencies, error) {
	logrus.Info("AddCurrencies called")
	data, err := s.GetCurrencies()
	if err != nil {
		logrus.Errorf("Failed to fetch currencies: %v", err)
		return nil, fmt.Errorf("error fetching currencies: %s", err)
	}

	if err := s.currencyRepo.AddCurrencies(data); err != nil {
		logrus.Errorf("Failed to add or update currency: %v", err)
		return nil, fmt.Errorf("error adding or updating currency: %s", err)
	}

	logrus.Info("Currencies data added or updated successfully")
	return data, nil
}

func (s *ratesService) GetExchangeRates() (*dto.ExchangeRates, error) {
	logrus.Info("GetExchangeRates called")
	client := utils.Default()
	client.PrivateToken = config.Cfg.APIKey

	res, err := client.FastGet(config.Cfg.URLs.LatestExchangeRates)
	if err != nil {
		logrus.Errorf("Failed to make request to API: %v", err)
		return nil, fmt.Errorf("error when creating a request: %s", err)
	}

	var data dto.ExchangeRates
	if err := json.Unmarshal(res.Body(), &data); err != nil {
		logrus.Errorf("JSON parsing failed: %v", err)
		return nil, fmt.Errorf("error during JSON parsing: %s", err)
	}

	logrus.Debugf("Successfully fetched and parsed exchange rates data: %v", data)
	return &data, nil
}

func (s *ratesService) AddRates() (*dto.ExchangeRates, error) {
	logrus.Info("AddRates called")
	data, err := s.GetExchangeRates()
	if err != nil {
		logrus.Errorf("Failed to fetch exchange rates: %v", err)
		return nil, fmt.Errorf("error fetching exchange rates: %s", err)
	}

	if err := s.exchangeRatesRepo.AddRates(data); err != nil {
		logrus.Errorf("Failed to add or update exchange rates: %v", err)
		return nil, fmt.Errorf("error adding or updating exchange rates: %s", err)
	}

	logrus.Info("Exchange rates data added or updated successfully")
	return data, nil
}

func (s *ratesService) UpdateRates() (string, error) {
	logrus.Info("UpdateRates called")
	if err := s.rateHistoriesRepo.UpdateRatesHistories(); err != nil {
		logrus.Errorf("Failed to update exchange rate histories: %v", err)
		return "", fmt.Errorf("error updating exchange rate histories: %s", err)
	}

	if _, err := s.AddRates(); err != nil {
		logrus.Errorf("Failed to insert latest exchange rates: %v", err)
		return "", fmt.Errorf("error inserting latest exchange rates: %s", err)
	}

	logrus.Info("Exchange rates updated successfully")
	return "Exchange rates updated successfully", nil
}
