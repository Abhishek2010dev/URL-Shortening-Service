package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func New(url string) *sql.DB {
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatalf("failed to open connection: %s", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping connection: %s", err)
	}

	log.Printf("Successfully connected to postgresql")

	return db
}
