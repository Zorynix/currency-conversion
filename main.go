package main

import (
	"currency-conversion/routes"
	"currency-conversion/utils"
	"flag"

	_ "github.com/go-sql-driver/mysql"
)

var (
	addr = flag.String("addr", ":8000", "TCP address to listen to")
)

func main() {

	flag.Parse()
	utils.InitLogger()
	routes.Routes(addr)

}
