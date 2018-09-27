package graphql

import (
	"github.com/pdt256/vbratings"
)

type PlayerRatingResolver struct {
	pr vbratings.PlayerAndRating
}

func (p *PlayerRatingResolver) Rating() int32 {
	return int32(p.pr.Rating)
}

func (p *PlayerRatingResolver) PlayerName() string {
	return p.pr.Name
}

func (p *PlayerRatingResolver) TotalMatches() int32 {
	return int32(p.pr.TotalMatches)
}

type query struct {
	playerRatingRepository vbratings.PlayerRatingRepository
}

func NewQuery(playerRatingRepository vbratings.PlayerRatingRepository) *query {
	return &query{playerRatingRepository}
}

type TopPlayersArgs struct {
	Year   *int32
	Gender *string
	Limit  *int32
}

func (a *TopPlayersArgs) GetYear() int {
	if a.Year != nil {
		return int(*a.Year)
	}

	return 2018
}

func (a *TopPlayersArgs) GetGender() string {
	if a.Gender != nil {
		return *a.Gender
	}

	return "male"
}

func (a *TopPlayersArgs) GetLimit() int {
	if a.Limit != nil {
		return int(*a.Limit)
	}

	return 10
}

func (q *query) TopPlayers(args TopPlayersArgs) []*PlayerRatingResolver {
	playerAndRatings := q.playerRatingRepository.GetTopPlayerRatings(
		args.GetYear(),
		vbratings.GenderFromString(args.GetGender()),
		args.GetLimit())

	var r []*PlayerRatingResolver
	for _, value := range playerAndRatings {
		r = append(r, &PlayerRatingResolver{value})
	}

	return r
}
