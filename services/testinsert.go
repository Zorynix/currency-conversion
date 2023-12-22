package services

import (
	"currency-conversion/models"
	"time"
)

func (MSQ *Mysql) TestInsert() error {

	const layout = "2006-01-02 15:04:05"
	createdAtStr := "22-12-2023 15:49:30"
	updatedAtStr := "25-01-2024 12:30:30"

	createdAt, err := time.Parse(layout, createdAtStr)
	if err != nil {
		panic(err)
	}

	updatedAt, err := time.Parse(layout, updatedAtStr)
	if err != nil {
		panic(err)
	}

	MSQ.db.Save(models.CurrenciesExchangeRates{Id: 919, CurrencyId: 32, TargetCurencyId: 12, ExchangeRate: 0.3341324, RateSourceId: 3411, CreatedAt: createdAt, UpdatedAt: updatedAt})
	return nil
}
