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
	db *sql.DB
}

func NewSqliteMatchRepository(dbPath string) *sqliteMatchRepository {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}

	return &sqliteMatchRepository{db}
}

func (r *sqliteMatchRepository) InitDB() {
	sqlStmt := `CREATE TABLE match (
			id TEXT NOT NULL PRIMARY KEY
			,playerA_id TEXT NOT NULL
			,playerB_id TEXT NOT NULL
			,playerC_id TEXT NOT NULL
			,playerD_id TEXT NOT NULL
			,isForfeit BOOLEAN NOT NULL
		);
		DELETE FROM match;`

	_, createError := r.db.Exec(sqlStmt)
	if createError != nil {
		log.Printf("%q: %s\n", createError, sqlStmt)
		return
	}
}

func (r *sqliteMatchRepository) Create(match Match, id string) error {
	_, err := r.db.Exec(
		"INSERT INTO match(id, playerA_id, playerB_id, playerC_id, playerD_id, isForfeit) VALUES ($1, $2, $3, $4, $5, $6)",
		id,
		match.PlayerA.BvbId,
		match.PlayerB.BvbId,
		match.PlayerC.BvbId,
		match.PlayerD.BvbId,
		match.IsForfeit,
	)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (r *sqliteMatchRepository) Find(id string) *Match {
	var m Match
	row := r.db.QueryRow("SELECT playerA_id, playerB_id, playerC_id, playerD_id, isForfeit FROM match WHERE id = $1", id)
	err := row.Scan(
		&m.PlayerA.BvbId,
		&m.PlayerB.BvbId,
		&m.PlayerC.BvbId,
		&m.PlayerD.BvbId,
		&m.IsForfeit,
	)
	if err != nil {
		log.Fatal(err)
	}

	return &m
}
