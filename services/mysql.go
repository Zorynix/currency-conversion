package services

import (
	"context"
	"currency-conversion/utils"
	"database/sql"
	"os"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql struct {
	DB *gorm.DB
}

func NewMySQL(ctx context.Context) (*Mysql, error) {

	utils.LoadEnv()

	dsn := os.Getenv("DSN")
	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Error().Err(err).Msg("Failed to open SQL connection")
		return nil, err
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		log.Error().Err(err).Msg("Failed to initialize Gorm")
		return nil, err
	}

	return &Mysql{DB: gormDB}, nil
}

type Database interface {
	GetDB() *gorm.DB
}

func (m *Mysql) GetDB() *gorm.DB {
	return m.DB
}
