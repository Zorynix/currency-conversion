package repo

import (
	"currency-conversion/dto"
	"currency-conversion/models"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ExchangeRatesRepo interface {
	GetExchangeRates() (*dto.ExchangeRates, error)
	AddRates(data *dto.ExchangeRates) error
}

type exchangeRatesRepo struct {
	DB *gorm.DB
}

func NewExchangeRatesRepo(db *gorm.DB) ExchangeRatesRepo {
	return &exchangeRatesRepo{DB: db}
}

func (r *exchangeRatesRepo) GetExchangeRates() (*dto.ExchangeRates, error) {
	log.Info().Msg("GetExchangeRates called")
	var exchangeRates []models.ExchangeRates
	if err := r.DB.Find(&exchangeRates).Error; err != nil {
		log.Error().Err(err).Msg("Failed to fetch exchange rates from database")
		return nil, err
	}

	ratesMap := make(map[string]models.ExchangeRates)
	for _, rate := range exchangeRates {
		ratesMap[rate.Code] = rate
	}

	data := &dto.ExchangeRates{Data: ratesMap}
	log.Debug().Interface("ExchangeRatesData", data).Msg("Successfully fetched exchange rates data")
	return data, nil
}

func (r *exchangeRatesRepo) AddRates(data *dto.ExchangeRates) error {
	log.Info().Msg("AddRates called")
	tx := r.DB.Begin()
	if tx.Error != nil {
		log.Error().Err(tx.Error).Msg("Failed to begin transaction")
		return tx.Error
	}

	for _, rate := range data.Data {
		if err := tx.Clauses(clause.OnConflict{UpdateAll: true}).Create(&rate).Error; err != nil {
			tx.Rollback()
			log.Error().Err(err).Msg("Failed to add or update exchange rate")
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		log.Error().Err(err).Msg("Failed to commit transaction")
		return err
	}

	log.Info().Msg("Exchange rates data inserted or updated in database successfully")
	return nil
}
