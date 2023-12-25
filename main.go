package main

import (
	"currency-conversion/routes"
	"flag"

	_ "github.com/go-sql-driver/mysql"
)

var (
	addr = flag.String("addr", ":8000", "TCP address to listen to")
)

func main() {

	flag.Parse()

	routes.Routes(addr)
	//currency_codes := []string{}

}

// for _, currencyCode := range currency_codes {
//go updateRate
// }

// fmt.Println("Hello, 世界")

func updateRates() {
	// Currencyapi Service # hourly
	// ECB Service # daily

	// rates Service interface

	// currency_codes := [180]string{"USD", "CAD", "CNY"}

	// for _, currencyCode := range currency_codes {
	// 	go updateRate
	// }
}

func updateRate( /* currencyCode */ ) {
	// get latest from API
	// update to DB
	// --- resave to history table

	// @TODO: add event to queue (rabbitmq)
}
