package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

var (
	url_all_currencies        = "https://api.currencyapi.com/v3/currencies"
	url_latest_exchange_rates = "https://api.currencyapi.com/v3/latest"
	methodGet                 = "GET"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Panic().Msg("No .env file found")
	}
}

func TestApi() ([]byte, error) {

	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error when uploading a file .env")
		panic(err)
	}

	client := &http.Client{}

	req, err := http.NewRequest(methodGet, url_all_currencies, nil)
	if err != nil {
		return nil, fmt.Errorf("Error when creating a request: %s", err)
	}

	apiKEY := os.Getenv("API_KEY")

	req.Header.Add("apikey", apiKEY)

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error when executing a request: %s", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading data from the response: %s", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("Error during JSON parsing: %s", err)
	}

	prettiedJSON, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("Error when formatting JSON: %s", err)
	}

	return prettiedJSON, nil
}
