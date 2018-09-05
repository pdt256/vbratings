package vbscraper

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type PlayerRating struct {
	PlayerId   int
	Year       int
	SeedRating int
	Rating     int
}

type PlayerRatingRepository interface {
	Create(playerRating PlayerRating) error
	GetPlayerRatingByYear(playerId int, year int) *PlayerRating
}

type sqlitePlayerRatingRepository struct {
	db *sql.DB
}

var _ PlayerRatingRepository = (*sqlitePlayerRatingRepository)(nil)

func NewSqlitePlayerRatingRepository(db *sql.DB) *sqlitePlayerRatingRepository {
	return &sqlitePlayerRatingRepository{db}
}

func (r *sqlitePlayerRatingRepository) InitDB() {
	sqlStmt := `CREATE TABLE player_rating (
			playerId INT NOT NULL
			,year INT NOT NULL
			,seedRating INT NOT NULL
			,rating INT NOT NULL
		);
		DELETE FROM player_rating;`

	_, createError := r.db.Exec(sqlStmt)
	if createError != nil {
		log.Printf("%q: %s\n", createError, sqlStmt)
		return
	}
}

func (r *sqlitePlayerRatingRepository) Create(playerRating PlayerRating) error {
	_, err := r.db.Exec(
		"INSERT INTO player_rating(playerId, year, seedRating, rating) VALUES ($1, $2, $3, $4)",
		playerRating.PlayerId,
		playerRating.Year,
		playerRating.SeedRating,
		playerRating.Rating,
	)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (r *sqlitePlayerRatingRepository) GetPlayerRatingByYear(playerId int, year int) *PlayerRating {
	var pr PlayerRating
	row := r.db.QueryRow("SELECT playerId, year, seedRating, rating FROM player_rating WHERE playerId = $1 AND year = $2", playerId, year)
	err := row.Scan(
		&pr.PlayerId,
		&pr.Year,
		&pr.SeedRating,
		&pr.Rating,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		log.Fatalf("%#v", err)
	}

	return &pr
}
