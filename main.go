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
