package vbratings_test

import (
	"testing"

	"github.com/pdt256/vbratings"
	"github.com/pdt256/vbratings/sqlite"
	"github.com/stretchr/testify/assert"
)

var (
	playerAId = "98556e2665224abc99c2d07d621befa7"
	playerBId = "7a30fab9631a442d83b70c9bf1293be8"
	playerCId = "91f67d94a9a54c91b9f0ee0efc497c28"
	playerDId = "ee655d72d148459ca10d05cce939bcab"
)

func Test_RatingCalculator_CalculateRatingsByYear_SingleMatch(t *testing.T) {
	match := vbratings.Match{
		PlayerAId: playerAId,
		PlayerBId: playerBId,
		PlayerCId: playerCId,
		PlayerDId: playerDId,
		IsForfeit: false,
		Set1:      "17-21",
		Set2:      "21-15",
		Set3:      "15-7",
		Year:      2018,
	}
	db := sqlite.NewInMemoryDB()
	matchRepository := sqlite.NewMatchRepository(db)
	matchRepository.Create(match, "123-abc")
	playerRatingRepository := sqlite.NewPlayerRatingRepository(db)
	ratingCalculator := vbratings.NewRatingCalculator(matchRepository, playerRatingRepository)

	// When
	ratingCalculator.CalculateRatingsByYear(2018)

	// Then
	playerRatingA, _ := playerRatingRepository.GetPlayerRatingByYear(playerAId, 2018)
	playerRatingB, _ := playerRatingRepository.GetPlayerRatingByYear(playerBId, 2018)
	playerRatingC, _ := playerRatingRepository.GetPlayerRatingByYear(playerCId, 2018)
	playerRatingD, _ := playerRatingRepository.GetPlayerRatingByYear(playerDId, 2018)
	assertPlayerRating(t, playerRatingA, 1500, 1516, 2018)
	assertPlayerRating(t, playerRatingB, 1500, 1516, 2018)
	assertPlayerRating(t, playerRatingC, 1500, 1484, 2018)
	assertPlayerRating(t, playerRatingD, 1500, 1484, 2018)
}

func Test_RatingCalculator_CalculateRatingsByYear_SeededWithPlayerRatingFromPreviousMatch(t *testing.T) {
	// Given
	match1 := vbratings.Match{
		PlayerAId: playerAId,
		PlayerBId: playerBId,
		PlayerCId: playerCId,
		PlayerDId: playerDId,
		IsForfeit: false,
		Set1:      "17-21",
		Set2:      "21-15",
		Set3:      "15-7",
		Year:      2018,
		Gender:    "female",
	}
	match2 := vbratings.Match{
		PlayerAId: playerAId,
		PlayerBId: playerBId,
		PlayerCId: playerCId,
		PlayerDId: playerDId,
		IsForfeit: false,
		Set1:      "17-21",
		Set2:      "21-15",
		Set3:      "15-7",
		Year:      2018,
		Gender:    "female",
	}
	db := sqlite.NewInMemoryDB()
	matchRepository := sqlite.NewMatchRepository(db)
	matchRepository.Create(match1, "match1")
	matchRepository.Create(match2, "match2")
	playerRatingRepository := sqlite.NewPlayerRatingRepository(db)
	ratingCalculator := vbratings.NewRatingCalculator(matchRepository, playerRatingRepository)

	// When
	ratingCalculator.CalculateRatingsByYear(2018)

	// Then
	playerRatingA, _ := playerRatingRepository.GetPlayerRatingByYear(playerAId, 2018)
	playerRatingB, _ := playerRatingRepository.GetPlayerRatingByYear(playerBId, 2018)
	playerRatingC, _ := playerRatingRepository.GetPlayerRatingByYear(playerCId, 2018)
	playerRatingD, _ := playerRatingRepository.GetPlayerRatingByYear(playerDId, 2018)
	assertPlayerRating(t, playerRatingA, 1500, 1530, 2018)
	assertPlayerRating(t, playerRatingB, 1500, 1530, 2018)
	assertPlayerRating(t, playerRatingC, 1500, 1469, 2018)
	assertPlayerRating(t, playerRatingD, 1500, 1469, 2018)
	assert.Equal(t, 2, playerRatingA.TotalMatches)
	assert.Equal(t, "female", playerRatingA.Gender)
}

func Test_RatingCalculator_CalculateRatingsByYear_SeededWithPreviousYearPlayerRating(t *testing.T) {
	// Given
	playerRating := vbratings.PlayerRating{
		PlayerId:   playerAId,
		Year:       2017,
		SeedRating: 1500,
		Rating:     1600,
	}
	match := vbratings.Match{
		PlayerAId: playerAId,
		PlayerBId: playerBId,
		PlayerCId: playerCId,
		PlayerDId: playerDId,
		IsForfeit: false,
		Set1:      "17-21",
		Set2:      "21-15",
		Set3:      "15-7",
		Year:      2018,
	}
	db := sqlite.NewInMemoryDB()
	matchRepository := sqlite.NewMatchRepository(db)
	matchRepository.Create(match, "123-abc")
	playerRatingRepository := sqlite.NewPlayerRatingRepository(db)
	playerRatingRepository.Create(playerRating)
	ratingCalculator := vbratings.NewRatingCalculator(matchRepository, playerRatingRepository)

	// When
	ratingCalculator.CalculateRatingsByYear(2018)

	// Then
	playerRatingA, _ := playerRatingRepository.GetPlayerRatingByYear(playerAId, 2018)
	playerRatingB, _ := playerRatingRepository.GetPlayerRatingByYear(playerBId, 2018)
	playerRatingC, _ := playerRatingRepository.GetPlayerRatingByYear(playerCId, 2018)
	playerRatingD, _ := playerRatingRepository.GetPlayerRatingByYear(playerDId, 2018)
	assertPlayerRating(t, playerRatingA, 1600, 1611, 2018)
	assertPlayerRating(t, playerRatingB, 1500, 1516, 2018)
	assertPlayerRating(t, playerRatingC, 1500, 1486, 2018)
	assertPlayerRating(t, playerRatingD, 1500, 1486, 2018)
}

func assertPlayerRating(t *testing.T, playerRating *vbratings.PlayerRating, expectedSeedRating int, expectedRating int, expectedYear int) {
	assert.Equal(t, expectedSeedRating, playerRating.SeedRating)
	assert.Equal(t, expectedRating, playerRating.Rating)
	assert.Equal(t, expectedYear, playerRating.Year)
}
