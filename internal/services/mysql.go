package services

import (
	"context"
	"database/sql"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql struct {
	DB *gorm.DB
}

func NewMySQL(ctx context.Context, dsn string) (*Mysql, error) {
	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		logrus.Errorf("Failed to open SQL connection: %v", err)
		return nil, err
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		logrus.Errorf("Failed to initialize Gorm: %v", err)
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
