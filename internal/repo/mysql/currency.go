package mysql

import (
	"context"
	"currency-conversion/internal/dto"
	"currency-conversion/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type currencyRepo struct {
	DB *gorm.DB
}

func NewCurrencyRepo(db *gorm.DB) *currencyRepo {
	return &currencyRepo{DB: db}
}

func (r *currencyRepo) GetCurrencies(ctx context.Context) (*dto.Currencies, error) {
	logrus.Info("GetCurrencies called")
	var currencies []entity.Currency
	if err := r.DB.Find(&currencies).Error; err != nil {
		logrus.Errorf("Failed to fetch currencies from database: %v", err)
		return nil, err
	}

	currencyMap := make(map[string]entity.Currency)
	for _, currency := range currencies {
		currencyMap[currency.Code] = currency
	}

	data := &dto.Currencies{Data: currencyMap}
	logrus.WithField("CurrenciesData", data).Debug("Successfully fetched currencies data")
	return data, nil
}

func (r *currencyRepo) AddCurrencies(ctx context.Context, data *dto.Currencies) error {
	logrus.Info("AddCurrencies called")
	tx := r.DB.Begin()
	if tx.Error != nil {
		logrus.Errorf("Failed to begin transaction: %v", tx.Error)
		return tx.Error
	}

	for _, value := range data.Data {
		if err := tx.Clauses(clause.OnConflict{UpdateAll: true}).Create(&value).Error; err != nil {
			tx.Rollback()
			logrus.Errorf("Failed to add or update currency: %v", err)
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		logrus.Errorf("Failed to commit transaction: %v", err)
		return err
	}

	logrus.Info("Currencies data inserted or updated in database successfully")
	return nil
}
