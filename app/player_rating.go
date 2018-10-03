package app

import (
	"github.com/pdt256/vbratings"
)

type PlayerRating struct {
	playerRatingRepository vbratings.PlayerRatingRepository
}

// Top player ratings by year and gender
func (pr *PlayerRating) GetTopPlayerRatings(year int, gender string, limit int) []vbratings.PlayerAndRating {
	return pr.playerRatingRepository.GetTopPlayerRatings(
		year,
		gender,
		limit)
}

// Create player rating
func (pr *PlayerRating) Create(
	playerId string,
	year int,
	gender string,
	seedRating int,
	rating int,
	totalMatches int) error {
	playerRating := vbratings.PlayerRating{
		PlayerId:     playerId,
		Year:         year,
		Gender:       gender,
		SeedRating:   seedRating,
		Rating:       rating,
		TotalMatches: totalMatches,
	}
	return pr.playerRatingRepository.Create(playerRating)
}
