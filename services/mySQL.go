package services

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type Mysql struct {
	db *sql.DB
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Panic().Msg("No .env file found")
	}
}

// Database settings
var (
	host     = os.Getenv("host")
	port     = os.Getenv("port")
	user     = os.Getenv("user")
	password = os.Getenv("password")
	dbname   = os.Getenv("dbname")
)

func NewMySQL(ctx context.Context) (*Mysql, error) {
	conn, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", user, password, dbname))

	if err != nil {
		log.Fatal().Interface("unable to create mysql connection: %v", err).Msg("")
	}

	err = conn.Ping()

	if err != nil {
		log.Fatal().Interface("unable to ping mysql connection: %v", err).Msg("")

	} else {
		fmt.Println("mysqld is alive")
	}

	return &Mysql{db: conn}, nil
}

func (mys *Mysql) Close() {
	mys.db.Close()
}
