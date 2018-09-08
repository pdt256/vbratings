package vbscraper

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type PlayerRating struct {
	PlayerId   int
	Year       int
	SeedRating int
	Rating     int
}

type PlayerAndRating struct {
	Player
	PlayerRating
}

type PlayerRatingRepository interface {
	Create(playerRating PlayerRating)
	GetPlayerRatingByYear(playerId int, year int) (*PlayerRating, error)
	GetTopPlayerRatings(year int) []PlayerAndRating
}

type sqlitePlayerRatingRepository struct {
	db *sql.DB
}

var _ PlayerRatingRepository = (*sqlitePlayerRatingRepository)(nil)

func NewSqlitePlayerRatingRepository(db *sql.DB) *sqlitePlayerRatingRepository {
	return &sqlitePlayerRatingRepository{db}
}

func (r *sqlitePlayerRatingRepository) InitDB() {
	sqlStmt := `CREATE TABLE IF NOT EXISTS player_rating (
			playerId INT NOT NULL
			,year INT NOT NULL
			,seedRating INT NOT NULL
			,rating INT NOT NULL
		);
		DELETE FROM player_rating;`

	_, createError := r.db.Exec(sqlStmt)
	checkError(createError)
}

func (r *sqlitePlayerRatingRepository) Create(playerRating PlayerRating) {
	_, err := r.db.Exec(
		"INSERT OR REPLACE INTO player_rating(playerId, year, seedRating, rating) VALUES ($1, $2, $3, $4)",
		playerRating.PlayerId,
		playerRating.Year,
		playerRating.SeedRating,
		playerRating.Rating,
	)
	checkError(err)
}

func (r *sqlitePlayerRatingRepository) GetPlayerRatingByYear(playerId int, year int) (*PlayerRating, error) {
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
			return nil, &NotFoundError{}
		}
		checkError(err)
	}

	return &pr, nil
}

func (r *sqlitePlayerRatingRepository) GetTopPlayerRatings(year int) []PlayerAndRating {
	var playerAndRatings []PlayerAndRating

	rows, queryErr := r.db.Query(`SELECT
		p.bvbId, p.name, p.imgUrl,
		pr.playerId, pr.year, pr.seedRating, pr.rating
		FROM player_rating AS pr
		INNER JOIN player AS p ON p.bvbId = pr.playerId
		WHERE year = 2018
		ORDER BY rating DESC;`)
	checkError(queryErr)

	defer rows.Close()

	for rows.Next() {
		var par PlayerAndRating
		checkError(rows.Scan(
			&par.BvbId,
			&par.Name,
			&par.ImgUrl,
			&par.PlayerId,
			&par.Year,
			&par.SeedRating,
			&par.Rating,
		))

		playerAndRatings = append(playerAndRatings, par)
	}
	checkError(rows.Err())

	return playerAndRatings
}
