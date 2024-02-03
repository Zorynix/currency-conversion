package services

import (
	"context"
	"currency-conversion/models"
	"currency-conversion/utils"
	"os"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql struct {
	db *gorm.DB
}

func NewMySQL(ctx context.Context) (*Mysql, error) {

	utils.LoadEnv()

	DSN := os.Getenv("DSN")

	conn, err := gorm.Open(mysql.New(mysql.Config{
		DSN: DSN}))

	if err != nil {
		log.Fatal().Interface("unable to create mysql connection pool: %v", err).Msg("")
	}

	err = conn.AutoMigrate(&models.CurrencyExchangeRateHistory{}, &models.CurrenciesExchangeRates{}, &models.Currency{})
	if err != nil {
		log.Fatal().Interface("unable to automigrate: %v", err).Msg("")
	}

	return &Mysql{db: conn}, nil
}

func (msq *Mysql) Ping(ctx context.Context) error {
	db, err := msq.db.DB()
	if err != nil {
		log.Fatal().Interface("unable to create mysql connection pool: %v", err).Msg("")
	}

	return db.Ping()
}

func (msq *Mysql) Close() {

	db, err := msq.db.DB()
	if err != nil {
		log.Fatal().Interface("unable to create mysql connection pool: %v", err).Msg("")
	}
	db.Close()
}
