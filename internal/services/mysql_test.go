package services_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"currency-conversion/internal/services"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func setupTestDatabase(t *testing.T) (string, func()) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "mysql:8.0",
		ExposedPorts: []string{"3306/tcp"},
		Env: map[string]string{
			"MYSQL_ROOT_PASSWORD": "root",
			"MYSQL_DATABASE":      "testdb",
		},
		WaitingFor: wait.ForLog("ready for connections").WithStartupTimeout(5 * time.Minute),
	}

	mysqlC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Fatal(err)
	}

	cleanup := func() {
		if err := mysqlC.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %s", err)
		}
	}

	ip, err := mysqlC.Host(ctx)
	if err != nil {
		t.Fatal(err)
	}

	port, err := mysqlC.MappedPort(ctx, "3306")
	if err != nil {
		t.Fatal(err)
	}

	dsn := fmt.Sprintf("root:root@tcp(%s:%s)/testdb?charset=utf8mb4&parseTime=True&loc=Local", ip, port.Port())
	return dsn, cleanup
}

func TestNewMySQL(t *testing.T) {
	dsn, cleanup := setupTestDatabase(t)
	defer cleanup()

	ctx := context.Background()
	mysqlService, err := services.NewMySQL(ctx, dsn)
	assert.NoError(t, err)
	assert.NotNil(t, mysqlService)
	assert.NotNil(t, mysqlService.GetDB())

	sqlDB, err := mysqlService.GetDB().DB()
	assert.NoError(t, err)
	err = sqlDB.Ping()
	assert.NoError(t, err)
}

func TestMySQL_GetDB(t *testing.T) {
	dsn, cleanup := setupTestDatabase(t)
	defer cleanup()

	ctx := context.Background()
	mysqlService, err := services.NewMySQL(ctx, dsn)
	assert.NoError(t, err)
	assert.NotNil(t, mysqlService)

	db := mysqlService.GetDB()
	assert.NotNil(t, db)

	var version string
	err = db.Raw("SELECT VERSION()").Scan(&version).Error
	assert.NoError(t, err)
	logrus.Infof("MySQL version: %s", version)
}
