package vbscraper

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type MatchRepository interface {
	Create(match Match, id string) error
	Find(id string) *Match
	GetAllPlayerIds() []int
}

type sqliteMatchRepository struct {
	db *sql.DB
}

func NewSqliteMatchRepository(db *sql.DB) *sqliteMatchRepository {
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
			,set1 TEXT NOT NULL
			,set2 TEXT NOT NULL
			,set3 TEXT NOT NULL
			,year INT NOT NULL
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
		"INSERT INTO match(id, playerA_id, playerB_id, playerC_id, playerD_id, isForfeit, set1, set2, set3, year) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",
		id,
		match.PlayerAId,
		match.PlayerBId,
		match.PlayerCId,
		match.PlayerDId,
		match.IsForfeit,
		match.Set1,
		match.Set2,
		match.Set3,
		match.Year,
	)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (r *sqliteMatchRepository) Find(id string) *Match {
	var m Match
	row := r.db.QueryRow("SELECT playerA_id, playerB_id, playerC_id, playerD_id, isForfeit, set1, set2, set3, year FROM match WHERE id = $1", id)
	err := row.Scan(
		&m.PlayerAId,
		&m.PlayerBId,
		&m.PlayerCId,
		&m.PlayerDId,
		&m.IsForfeit,
		&m.Set1,
		&m.Set2,
		&m.Set3,
		&m.Year,
	)
	if err != nil {
		log.Fatal(err)
	}

	return &m
}

func (r *sqliteMatchRepository) GetAllPlayerIds() []int {
	var playerIds []int

	rows, queryErr := r.db.Query("SELECT *" +
		" FROM (SELECT playerA_id AS id FROM match" +
		" UNION SELECT playerB_id AS id FROM match" +
		" UNION SELECT playerC_id AS id FROM match" +
		" UNION SELECT playerD_id AS id FROM match)" +
		" GROUP BY id")
	if queryErr != nil {
		log.Fatal(queryErr)
	}
	defer rows.Close()

	for rows.Next() {
		var playerId int
		scanErr := rows.Scan(&playerId)
		if scanErr != nil {
			log.Fatal(scanErr)
		}
		playerIds = append(playerIds, playerId)
	}
	err := rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return playerIds
}
