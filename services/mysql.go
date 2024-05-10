package services

import (
	"context"
	"currency-conversion/dto"
	"currency-conversion/models"
	"currency-conversion/utils"
	"os"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database interface {
	Ping(ctx context.Context) error

	GetCurrencies() (*dto.Currencies, error)
	InsertCurrencies() (*dto.Currencies, error)
	GetExchangeRates() (*dto.ExchangeRates, error)
	UpdateRates() (*dto.ExchangeRateHistory, error)
	InsertExchangeRates() (*dto.ExchangeRates, error)
}

type Mysql struct {
	DB *gorm.DB
}

func NewMySQL(ctx context.Context) (Database, error) {

	utils.LoadEnv()

	DSN := os.Getenv("DSN")

	conn, err := gorm.Open(mysql.New(mysql.Config{
		DSN: DSN}))

	if err != nil {
		log.Fatal().Interface("unable to create mysql connection pool: %v", err).Msg("")
	}

	err = conn.AutoMigrate(&models.ExchangeRateHistory{}, &models.ExchangeRates{}, &models.Currency{})
	if err != nil {
		log.Fatal().Interface("unable to automigrate: %v", err).Msg("")
	}

	return &Mysql{DB: conn}, nil
}

func (msq *Mysql) Ping(ctx context.Context) error {
	db, err := msq.DB.DB()
	if err != nil {
		log.Fatal().Interface("unable to create mysql connection pool: %v", err).Msg("")
	}

	return db.Ping()
}

func (msq *Mysql) Close() {

	db, err := msq.DB.DB()
	if err != nil {
		log.Fatal().Interface("unable to create mysql connection pool: %v", err).Msg("")
	}
	db.Close()
}
