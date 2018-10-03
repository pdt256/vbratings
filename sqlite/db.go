package sqlite

import (
	"database/sql"
	"log"
)

var InMemorySharedDbPath = "file:db?mode=memory&cache=shared"
var InMemoryDbPath = ":memory:"

func NewInMemoryDB() *sql.DB {
	db, err := sql.Open("sqlite3", InMemoryDbPath)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func NewFileDB(dbPath string) *sql.DB {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
