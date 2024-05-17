package repo

import (
	"currency-conversion/models"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type RateHistoriesRepo interface {
	UpdateRatesHistories() error
}

type rateHistoriesRepo struct {
	DB *gorm.DB
}

func NewRateHistoriesRepo(db *gorm.DB) RateHistoriesRepo {
	return &rateHistoriesRepo{DB: db}
}

func (r *rateHistoriesRepo) UpdateRatesHistories() error {
	log.Info().Msg("UpdateRatesHistories called")
	tx := r.DB.Begin()
	if tx.Error != nil {
		log.Error().Err(tx.Error).Msg("Failed to begin transaction")
		return tx.Error
	}

	var count int64
	if err := tx.Model(&models.ExchangeRates{}).Count(&count).Error; err != nil {
		tx.Rollback()
		log.Error().Err(err).Msg("Failed to count exchange rates")
		return err
	}

	if count == 0 {
		if err := tx.Exec("INSERT INTO exchange_rate_histories SELECT * FROM exchange_rates").Error; err != nil {
			tx.Rollback()
			log.Error().Err(err).Msg("Failed to insert initial exchange rate histories")
			return err
		}
		log.Info().Msg("Initial exchange rate histories inserted successfully")
	} else {
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
			return err
		}
		log.Info().Msg("Exchange rate histories updated successfully")
	}

	if err := tx.Commit().Error; err != nil {
		log.Error().Err(err).Msg("Failed to commit transaction")
		return err
	}

	log.Info().Msg("Rates histories transaction committed successfully")
	return nil
}
