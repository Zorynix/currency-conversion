package repo

import (
	"currency-conversion/models"

	"github.com/sirupsen/logrus"
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
	logrus.Info("UpdateRatesHistories called")
	tx := r.DB.Begin()
	if tx.Error != nil {
		logrus.Errorf("Failed to begin transaction: %v", tx.Error)
		return tx.Error
	}

	var count int64
	if err := tx.Model(&models.ExchangeRates{}).Count(&count).Error; err != nil {
		tx.Rollback()
		logrus.Errorf("Failed to count exchange rates: %v", err)
		return err
	}

	if count == 0 {
		if err := tx.Exec("INSERT INTO exchange_rate_histories SELECT * FROM exchange_rates").Error; err != nil {
			tx.Rollback()
			logrus.Errorf("Failed to insert initial exchange rate histories: %v", err)
			return err
		}
		logrus.Info("Initial exchange rate histories inserted successfully")
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
			logrus.Errorf("Failed to update exchange rate histories: %v", err)
			return err
		}
		logrus.Info("Exchange rate histories updated successfully")
	}

	if err := tx.Commit().Error; err != nil {
		logrus.Errorf("Failed to commit transaction: %v", err)
		return err
	}

	logrus.Info("Rates histories transaction committed successfully")
	return nil
}
