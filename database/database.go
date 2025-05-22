package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var createShortenTableQuery = `
CREATE TABLE IF NOT EXISTS shorten (
	id SERIAL PRIMARY KEY,
	url TEXT NOT NULL UNIQUE,
	short_code TEXT NOT NULL UNIQUE,
	access_count INT DEFAULT 0,
	created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
`

func Connect() *sql.DB {
	url := "postgresql://postgres@localhost:5432/shorten?sslmode=disable"
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatalf("failed to open connection: %s", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping connection: %s", err)
	}

	log.Printf("Successfully connected to postgresql")

	if _, err := db.Exec(createShortenTableQuery); err != nil {
		log.Fatalf("failed to create `shorten` table: %s", err)
	}

	return db
}
