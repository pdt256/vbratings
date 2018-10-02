package vbratings

import (
	"errors"
)

type PlayerRating struct {
	PlayerId     string
	Year         int
	Gender       Gender
	SeedRating   int
	Rating       int
	TotalMatches int
}

type PlayerAndRating struct {
	Player
	PlayerRating
}

type PlayerRatingRepository interface {
	Create(playerRating PlayerRating) error
	GetPlayerRatingByYear(playerId string, year int) (*PlayerRating, error)
	GetTopPlayerRatings(year int, gender Gender, limit int) []PlayerAndRating
}

var PlayerRatingNotFoundError = errors.New("player rating not found")
