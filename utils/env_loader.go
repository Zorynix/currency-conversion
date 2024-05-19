package utils

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		logrus.Panic("---failed to load .env file---")
	}
}
