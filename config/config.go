package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DSN      string `mapstructure:"dsn"`
	APIKey   string `mapstructure:"api_key"`
	URLs     URLs   `mapstructure:"urls"`
	LogLevel string `mapstructure:"log_level"`
}

type URLs struct {
	AllCurrencies       string `mapstructure:"all_currencies"`
	LatestExchangeRates string `mapstructure:"latest_exchange_rates"`
}

var Cfg *Config

func LoadConfig(configPath string) error {
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	if err := viper.Unmarshal(&Cfg); err != nil {
		return fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return nil
}
