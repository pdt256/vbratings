package graphql_test

import (
	"testing"

	"github.com/pdt256/vbratings"
	"github.com/pdt256/vbratings/graphql"
	"github.com/pdt256/vbratings/sqlite"
	"github.com/stretchr/testify/assert"
)

func Test_TopPlayers(t *testing.T) {
	// Given
	player := vbratings.Player{
		BvbId: 123,
		Name:  "John Doe",
	}
	playerRating := vbratings.PlayerRating{
		PlayerId:     123,
		Rating:       1500,
		Year:         2018,
		Gender:       vbratings.Male,
		TotalMatches: 5,
	}
	db := sqlite.NewInMemoryDB()
	playerRepository := sqlite.NewPlayerRepository(db)
	playerRepository.InitDB()
	playerRepository.Create(player)
	playerRatingRepository := sqlite.NewPlayerRatingRepository(db)
	playerRatingRepository.InitDB()
	playerRatingRepository.Create(playerRating)
	query := graphql.NewQuery(playerRatingRepository)

	// When
	playerRatingResolvers := query.TopPlayers(graphql.TopPlayersArgs{
		Year:   int32P(2018),
		Gender: stringP("male"),
		Limit:  int32P(1),
	})

	// Then
	assert.Equal(t, 1, len(playerRatingResolvers))
	assert.Equal(t, int32(1500), playerRatingResolvers[0].Rating())
	assert.Equal(t, "John Doe", playerRatingResolvers[0].PlayerName())
	assert.Equal(t, int32(5), playerRatingResolvers[0].TotalMatches())
}

func int32P(input int) *int32 {
	output := int32(input)
	return &output
}

func stringP(input string) *string {
	return &input
}
