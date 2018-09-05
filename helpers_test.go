package vbscraper_test

import (
	"database/sql"
	"log"
)

func NewInMemoryDB() *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}

	return db
}
