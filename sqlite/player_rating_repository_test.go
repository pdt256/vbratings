package sqlite_test

import (
	"testing"

	"github.com/pdt256/vbratings"
	"github.com/pdt256/vbratings/sqlite"
	"github.com/stretchr/testify/assert"
)

func Test_PlayerRatingRepository_GetTopPlayerRatings(t *testing.T) {
	// Given
	topPlayerRating := vbratings.PlayerRating{
		PlayerId:   1,
		Year:       2018,
		SeedRating: 1500,
		Rating:     2400,
	}
	topPlayer := vbratings.Player{
		BvbId:  1,
		Name:   "John Doe",
		ImgUrl: "http://example.com/1.jpg",
	}
	secondPlayerRating := vbratings.PlayerRating{
		PlayerId:   2,
		Year:       2018,
		SeedRating: 1500,
		Rating:     2000,
	}
	secondPlayer := vbratings.Player{
		BvbId:  2,
		Name:   "John Smith",
		ImgUrl: "http://example.com/2.jpg",
	}
	db := sqlite.NewInMemoryDB()
	playerRepository := sqlite.NewPlayerRepository(db)
	playerRepository.InitDB()
	playerRepository.Create(topPlayer)
	playerRepository.Create(secondPlayer)
	playerRatingRepository := sqlite.NewPlayerRatingRepository(db)
	playerRatingRepository.InitDB()
	playerRatingRepository.Create(topPlayerRating)
	playerRatingRepository.Create(secondPlayerRating)

	// When
	playerAndRatings := playerRatingRepository.GetTopPlayerRatings(2018)

	// Then
	actualTopPlayer := playerAndRatings[0]
	assert.Equal(t, 1, actualTopPlayer.PlayerId)
	assert.Equal(t, "John Doe", actualTopPlayer.Name)
	assert.Equal(t, 2400, actualTopPlayer.Rating)
}
