package mysql_test

import (
	"context"
	"testing"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"currency-conversion/internal/dto"
	"currency-conversion/internal/entity"
	"currency-conversion/internal/repo"
)

var db *gorm.DB

func setupTestContainer(t *testing.T) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "mysql:8.0",
		ExposedPorts: []string{"3306/tcp"},
		Env: map[string]string{
			"MYSQL_ROOT_PASSWORD": "password",
			"MYSQL_DATABASE":      "testdb",
		},
		WaitingFor: wait.ForLog("port: 3306  MySQL Community Server - GPL."),
	}

	mysqlContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Fatalf("Could not start container: %s", err)
	}

	host, err := mysqlContainer.Host(ctx)
	if err != nil {
		t.Fatalf("Could not get container host: %s", err)
	}

	port, err := mysqlContainer.MappedPort(ctx, "3306")
	if err != nil {
		t.Fatalf("Could not get container port: %s", err)
	}

	dsn := "root:password@tcp(" + host + ":" + port.Port() + ")/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Could not connect to database: %s", err)
	}

	sqlCreateTable := `
    CREATE TABLE IF NOT EXISTS testdb.currencies (
        code VARCHAR(255) PRIMARY KEY,
        name VARCHAR(255),
        symbol_native VARCHAR(10),
        decimal_digits INT,
        active BOOLEAN,
        main_area_id INT,
		created_at DATETIME,
    	updated_at DATETIME,
    	deleted_at DATETIME
    );
    `
	if err := db.Exec(sqlCreateTable).Error; err != nil {
		t.Fatalf("Could not initialize database schema: %s", err)
	}

	sqlInsertData := `
    INSERT INTO testdb.currencies (code, name, symbol_native, decimal_digits, active, main_area_id, created_at, updated_at, deleted_at)
    VALUES ('USD', 'US Dollar', '$', 2, 0, 1, NOW(), NOW(), NULL),
           ('EUR', 'Euro', '€', 2, 0, 2, NOW(), NOW(), NULL);
    `
	if err := db.Exec(sqlInsertData).Error; err != nil {
		t.Fatalf("Could not insert initial data: %s", err)
	}

	time.Sleep(2 * time.Second)
}

func TestGetCurrencies(t *testing.T) {
	setupTestContainer(t)
	repos := repo.NewRepositories(db)

	ctx := context.Background()
	currencies, err := repos.Currency.GetCurrencies(ctx)
	if err != nil {
		t.Fatalf("Failed to get currencies: %s", err)
	}

	if len(currencies.Data) != 2 {
		t.Fatalf("Expected 2 currencies, got %d", len(currencies.Data))
	}

	if _, ok := currencies.Data["USD"]; !ok {
		t.Fatal("Expected currency USD not found")
	}

	if _, ok := currencies.Data["EUR"]; !ok {
		t.Fatal("Expected currency EUR not found")
	}
}

func TestAddCurrencies(t *testing.T) {
	setupTestContainer(t)
	repos := repo.NewRepositories(db)

	newCurrency := &dto.Currencies{
		Data: map[string]entity.Currency{
			"JPY": {
				Code:          "JPY",
				Name:          "Japanese Yen",
				SymbolNative:  "¥",
				DecimalDigits: 0,
				Active:        true,
				MainAreaId:    3,
			},
		},
	}

	ctx := context.Background()
	if err := repos.Currency.AddCurrencies(ctx, newCurrency); err != nil {
		t.Fatalf("Failed to add currency: %s", err)
	}

	currencies, err := repos.Currency.GetCurrencies(ctx)
	if err != nil {
		t.Fatalf("Failed to get currencies: %s", err)
	}

	if len(currencies.Data) != 3 {
		t.Fatalf("Expected 3 currencies, got %d", len(currencies.Data))
	}

	if _, ok := currencies.Data["JPY"]; !ok {
		t.Fatal("Expected currency JPY not found")
	}
}
