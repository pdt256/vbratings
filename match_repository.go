package vbscraper

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type matchRepository struct {
	dbPath string
}

func NewMatchRepository(dbPath string) *matchRepository {
	return &matchRepository{
		dbPath: dbPath,
	}
}

func (r *matchRepository) InitDB() {
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

func (r *matchRepository) Create(match match, id string) error {
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

func (r *matchRepository) getDB() *sql.DB {
	db, err := sql.Open("sqlite3", r.dbPath)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func (r *matchRepository) Find(id string) match {
	db := r.getDB()

	var m match
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

	return m
}
