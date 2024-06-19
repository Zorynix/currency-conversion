package main

import (
	"currency-conversion/internal/app"
)

const configPath = "config/config.yaml"

func main() {
	app.Run(configPath)
}
