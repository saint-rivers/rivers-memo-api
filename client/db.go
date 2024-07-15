package client

import (
	"database/sql"
	"log"
	"rivers-memo-cli/config"
)

func InitDatabase() *sql.DB {
	connStr := config.Envget("POSTGRES_URL")
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}
	return db
}