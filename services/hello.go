package services

import (
	"currency-conversion/models"
)

func (MSQ *Mysql) TestInsert() error {

	MSQ.db.Save(models.CurrenciesExchangeRates{Id: 100, CurrencyId: 0, TargetCurencyId: 22, ExchangeRate: "15", RateSourceId: "xui", CreatedAt: "2024-12-12 23:24:25", UpdatedAt: "2025-12-13 22:23:24"})

	return nil
}
