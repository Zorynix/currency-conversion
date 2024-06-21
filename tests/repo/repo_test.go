package repo_test

import (
	"currency-conversion/internal/repo"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestNewRepositories(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}

	repos := repo.NewRepositories(db)

	assert.NotNil(t, repos.Currency)
	assert.NotNil(t, repos.ExchangeRates)
	assert.NotNil(t, repos.RateHistories)
}
