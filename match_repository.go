package vbscraper

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type MatchRepository interface {
	Create(match Match, id string) error
	Find(id string) *Match
}

type sqliteMatchRepository struct {
	dbPath string
}

func NewSqliteMatchRepository(dbPath string) *sqliteMatchRepository {
	return &sqliteMatchRepository{
		dbPath: dbPath,
	}
}

func (r *sqliteMatchRepository) InitDB() {
	db := r.getDB()
	defer db.Close()

	sqlStmt := `CREATE TABLE match (
			id TEXT NOT NULL PRIMARY KEY
			,playerA_id TEXT NOT NULL
			,playerB_id TEXT NOT NULL
			,playerC_id TEXT NOT NULL
			,playerD_id TEXT NOT NULL
		);
		DELETE FROM match;`

	_, createError := db.Exec(sqlStmt)
	if createError != nil {
		log.Printf("%q: %s\n", createError, sqlStmt)
		return
	}
}

func (r *sqliteMatchRepository) Create(match Match, id string) error {
	db := r.getDB()

	_, err := db.Exec(
		"INSERT INTO match(id, playerA_id, playerB_id, playerC_id, playerD_id) VALUES ($1, $2, $3, $4, $5)",
		id,
		match.PlayerA.BvbId,
		match.PlayerB.BvbId,
		match.PlayerC.BvbId,
		match.PlayerD.BvbId,
	)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (r *sqliteMatchRepository) getDB() *sql.DB {
	db, err := sql.Open("sqlite3", r.dbPath)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func (r *sqliteMatchRepository) Find(id string) *Match {
	db := r.getDB()

	var m Match
	row := db.QueryRow("SELECT playerA_id, playerB_id, playerC_id, playerD_id FROM match WHERE id = $1", id)
	err := row.Scan(
		&m.PlayerA.BvbId,
		&m.PlayerB.BvbId,
		&m.PlayerC.BvbId,
		&m.PlayerD.BvbId,
	)
	if err != nil {
		log.Fatal(err)
	}

	return &m
}
