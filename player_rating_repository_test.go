package vbscraper_test

import (
	"testing"

	"github.com/pdt256/vbscraper"
	"github.com/stretchr/testify/assert"
)

func Test_PlayerRatingRepository_GetTopPlayerRatings(t *testing.T) {
	// Given
	topPlayerRating := vbscraper.PlayerRating{
		PlayerId:   1,
		Year:       2018,
		SeedRating: 1500,
		Rating:     2400,
	}
	topPlayer := vbscraper.Player{
		BvbId:  1,
		Name:   "John Doe",
		ImgUrl: "http://example.com/1.jpg",
	}
	secondPlayerRating := vbscraper.PlayerRating{
		PlayerId:   2,
		Year:       2018,
		SeedRating: 1500,
		Rating:     2000,
	}
	secondPlayer := vbscraper.Player{
		BvbId:  2,
		Name:   "John Smith",
		ImgUrl: "http://example.com/2.jpg",
	}
	db := NewInMemoryDB()
	playerRepository := vbscraper.NewSqlitePlayerRepository(db)
	playerRepository.InitDB()
	playerRepository.Create(topPlayer)
	playerRepository.Create(secondPlayer)
	playerRatingRepository := vbscraper.NewSqlitePlayerRatingRepository(db)
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
