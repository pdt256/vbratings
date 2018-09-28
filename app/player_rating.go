package app

import (
	"github.com/pdt256/vbratings"
)

type PlayerRating struct {
	playerRatingRepository vbratings.PlayerRatingRepository
}

// Top player ratings
// by year and gender
func (pr *PlayerRating) GetTopPlayerRatings(year int, gender string, limit int) []vbratings.PlayerAndRating {
	return pr.playerRatingRepository.GetTopPlayerRatings(
		year,
		vbratings.GenderFromString(gender),
		limit)
}

func (pr *PlayerRating) Create(
	playerId int,
	year int,
	gender string,
	seedRating int,
	rating int,
	totalMatches int) {
	playerRating := vbratings.PlayerRating{
		PlayerId:     playerId,
		Year:         year,
		Gender:       vbratings.GenderFromString(gender),
		SeedRating:   seedRating,
		Rating:       rating,
		TotalMatches: totalMatches,
	}
	pr.playerRatingRepository.Create(playerRating)
}

func (pr *PlayerRating) DoIt(isNow bool) {}

func (pr *PlayerRating) ReturnInts() []int {
	return []int{}
}
