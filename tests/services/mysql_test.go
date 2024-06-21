package services

import (
	"context"
	"os"
	"testing"

	"currency-conversion/config"
	mysql "currency-conversion/internal/services"

	gormmysql "gorm.io/driver/mysql"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const MockConfigContent = `{
	"dsn": "root:1234@tcp(localhost:3306)/test_db"
}`

func setupMockConfig(t *testing.T) string {
	file, err := os.CreateTemp("", "config*.json")
	require.NoError(t, err)
	defer file.Close()

	_, err = file.WriteString(MockConfigContent)
	require.NoError(t, err)

	return file.Name()
}

func TestNewMySQL(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		configPath := setupMockConfig(t)
		defer os.Remove(configPath)

		err := config.LoadConfig(configPath)
		require.NoError(t, err)

		db, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery("SELECT VERSION()").WillReturnRows(sqlmock.NewRows([]string{"VERSION()"}).AddRow("5.7.34"))
		mock.ExpectPing()

		dsn := config.Cfg.DSN
		mockDB, err := gorm.Open(gormmysql.New(gormmysql.Config{
			Conn: db,
		}), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
		require.NoError(t, err)

		ctx := context.Background()
		mysqlService, err := mysql.NewMySQL(ctx, dsn)
		require.NoError(t, err)
		assert.NotNil(t, mysqlService)
		assert.NotNil(t, mysqlService.DB)
		assert.Equal(t, mockDB.Dialector.Name(), mysqlService.DB.Dialector.Name())
	})

	t.Run("sql open error", func(t *testing.T) {
		configPath := setupMockConfig(t)
		defer os.Remove(configPath)

		config.Cfg.DSN = "invalid_dsn"

		ctx := context.Background()
		mysqlService, err := mysql.NewMySQL(ctx, config.Cfg.DSN)
		assert.Error(t, err)
		assert.Nil(t, mysqlService)
	})

	t.Run("gorm open error", func(t *testing.T) {
		configPath := setupMockConfig(t)
		defer os.Remove(configPath)

		err := config.LoadConfig(configPath)
		require.NoError(t, err)

		db, _, err := sqlmock.New()
		require.NoError(t, err)
		defer db.Close()

		config.Cfg.DSN = "invalid_dsn"

		ctx := context.Background()
		mysqlService, err := mysql.NewMySQL(ctx, config.Cfg.DSN)
		assert.Error(t, err)
		assert.Nil(t, mysqlService)
	})
}
