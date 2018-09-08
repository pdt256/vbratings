package vbratings

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

type PlayerRatingNotFoundError struct{ Err error }

func (e *PlayerRatingNotFoundError) Error() string { return "player rating not found" }
