package sqlite

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pdt256/vbratings"
)

type playerRatingRepository struct {
	db *sql.DB
}

var _ vbratings.PlayerRatingRepository = (*playerRatingRepository)(nil)

func NewPlayerRatingRepository(db *sql.DB) *playerRatingRepository {
	r := &playerRatingRepository{db}
	r.migrateDB()
	return r
}

func (r *playerRatingRepository) migrateDB() {
	sqlStmt := `CREATE TABLE IF NOT EXISTS player_rating (
			playerId TEXT NOT NULL
			,year INT NOT NULL
			,seedRating INT NOT NULL
			,rating TEXT NOT NULL
			,totalMatches INT NOT NULL
		);`

	_, createError := r.db.Exec(sqlStmt)
	checkError(createError)
}

func (r *playerRatingRepository) Create(playerRating vbratings.PlayerRating) error {
	_, err := r.db.Exec(
		"INSERT OR REPLACE INTO player_rating(playerId, year, seedRating, rating, totalMatches) VALUES ($1, $2, $3, $4, $5)",
		playerRating.PlayerId,
		playerRating.Year,
		playerRating.SeedRating,
		playerRating.Rating,
		playerRating.TotalMatches,
	)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (r *playerRatingRepository) GetPlayerRatingByYear(playerId string, year int) (*vbratings.PlayerRating, error) {
	var pr vbratings.PlayerRating
	row := r.db.QueryRow("SELECT playerId, year, seedRating, rating, totalMatches FROM player_rating WHERE playerId = $1 AND year = $2", playerId, year)
	err := row.Scan(
		&pr.PlayerId,
		&pr.Year,
		&pr.SeedRating,
		&pr.Rating,
		&pr.TotalMatches,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, vbratings.PlayerRatingNotFoundError
		}
		checkError(err)
	}

	return &pr, nil
}

func (r *playerRatingRepository) GetTopPlayerRatings(year int, gender string, limit int) []vbratings.PlayerAndRating {
	var playerAndRatings []vbratings.PlayerAndRating

	rows, queryErr := r.db.Query(`SELECT
		p.id, p.name,
		pr.playerId, pr.year, pr.seedRating, pr.rating, pr.totalMatches
		FROM player_rating AS pr
		INNER JOIN player AS p ON p.id = pr.playerId
		WHERE pr.year = $1 AND p.gender = $2
		ORDER BY rating DESC
		LIMIT $3;`, year, gender, limit)
	checkError(queryErr)

	defer rows.Close()

	for rows.Next() {
		var par vbratings.PlayerAndRating
		checkError(rows.Scan(
			&par.Id,
			&par.Name,
			&par.PlayerId,
			&par.Year,
			&par.SeedRating,
			&par.Rating,
			&par.TotalMatches,
		))

		playerAndRatings = append(playerAndRatings, par)
	}
	checkError(rows.Err())

	return playerAndRatings
}
