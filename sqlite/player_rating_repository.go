package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pdt256/vbratings"
)

type playerRatingRepository struct {
	db *sql.DB
}

var _ vbratings.PlayerRatingRepository = (*playerRatingRepository)(nil)

func NewPlayerRatingRepository(db *sql.DB) *playerRatingRepository {
	return &playerRatingRepository{db}
}

func (r *playerRatingRepository) InitDB() {
	sqlStmt := `CREATE TABLE IF NOT EXISTS player_rating (
			playerId INT NOT NULL
			,year INT NOT NULL
			,gender INT NOT NULL
			,seedRating INT NOT NULL
			,rating INT NOT NULL
			,totalMatches INT NOT NULL
		);
		DELETE FROM player_rating;`

	_, createError := r.db.Exec(sqlStmt)
	checkError(createError)
}

func (r *playerRatingRepository) Create(playerRating vbratings.PlayerRating) {
	_, err := r.db.Exec(
		"INSERT OR REPLACE INTO player_rating(playerId, year, gender, seedRating, rating, totalMatches) VALUES ($1, $2, $3, $4, $5, $6)",
		playerRating.PlayerId,
		playerRating.Year,
		playerRating.Gender,
		playerRating.SeedRating,
		playerRating.Rating,
		playerRating.TotalMatches,
	)
	checkError(err)
}

func (r *playerRatingRepository) GetPlayerRatingByYear(playerId int, year int) (*vbratings.PlayerRating, error) {
	var pr vbratings.PlayerRating
	row := r.db.QueryRow("SELECT playerId, year, gender, seedRating, rating, totalMatches FROM player_rating WHERE playerId = $1 AND year = $2", playerId, year)
	err := row.Scan(
		&pr.PlayerId,
		&pr.Year,
		&pr.Gender,
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

func (r *playerRatingRepository) GetTopPlayerRatings(year int, gender vbratings.Gender, limit int) []vbratings.PlayerAndRating {
	var playerAndRatings []vbratings.PlayerAndRating

	rows, queryErr := r.db.Query(`SELECT
		p.bvbId, p.name, p.imgUrl,
		pr.playerId, pr.year, pr.seedRating, pr.rating, pr.totalMatches
		FROM player_rating AS pr
		INNER JOIN player AS p ON p.bvbId = pr.playerId
		WHERE pr.year = $1 AND pr.gender = $2
		ORDER BY rating DESC
		LIMIT $3;`, year, gender, limit)
	checkError(queryErr)

	defer rows.Close()

	for rows.Next() {
		var par vbratings.PlayerAndRating
		checkError(rows.Scan(
			&par.BvbId,
			&par.Name,
			&par.ImgUrl,
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
