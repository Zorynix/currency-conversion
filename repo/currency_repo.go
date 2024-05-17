package repo

import (
	"currency-conversion/dto"
	"currency-conversion/models"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CurrencyRepo interface {
	GetCurrencies() (*dto.Currencies, error)
	AddCurrencies(data *dto.Currencies) error
}

type currencyRepo struct {
	DB *gorm.DB
}

func NewCurrencyRepo(db *gorm.DB) CurrencyRepo {
	return &currencyRepo{DB: db}
}

func (r *currencyRepo) GetCurrencies() (*dto.Currencies, error) {
	log.Info().Msg("GetCurrencies called")
	var currencies []models.Currency
	if err := r.DB.Find(&currencies).Error; err != nil {
		log.Error().Err(err).Msg("Failed to fetch currencies from database")
		return nil, err
	}

	currencyMap := make(map[string]models.Currency)
	for _, currency := range currencies {
		currencyMap[currency.Code] = currency
	}

	data := &dto.Currencies{Data: currencyMap}
	log.Debug().Interface("CurrenciesData", data).Msg("Successfully fetched currencies data")
	return data, nil
}

func (r *currencyRepo) AddCurrencies(data *dto.Currencies) error {
	log.Info().Msg("AddCurrencies called")
	tx := r.DB.Begin()
	if tx.Error != nil {
		log.Error().Err(tx.Error).Msg("Failed to begin transaction")
		return tx.Error
	}

	for _, value := range data.Data {
		if err := tx.Clauses(clause.OnConflict{UpdateAll: true}).Create(&value).Error; err != nil {
			tx.Rollback()
			log.Error().Err(err).Msg("Failed to add or update currency")
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		log.Error().Err(err).Msg("Failed to commit transaction")
		return err
	}

	log.Info().Msg("Currencies data inserted or updated in database successfully")
	return nil
}
