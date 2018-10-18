package vbratings_test

import (
	"testing"

	"github.com/pdt256/vbratings"
	"github.com/pdt256/vbratings/sqlite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	playerAId = "98556e2665224abc99c2d07d621befa7"
	playerBId = "7a30fab9631a442d83b70c9bf1293be8"
	playerCId = "91f67d94a9a54c91b9f0ee0efc497c28"
	playerDId = "ee655d72d148459ca10d05cce939bcab"
)

func Test_RatingCalculator_CalculateRatingsByYear_SingleMatch(t *testing.T) {
	tournament := vbratings.Tournament{
		Id:     "1b97c1b593f84566bb932678eaf8a30d",
		Gender: "male",
		Year:   2018,
	}
	match := vbratings.Match{
		Id:           "0b9fb995e28a451fa5aaca68d397c1e0",
		PlayerAId:    playerAId,
		PlayerBId:    playerBId,
		PlayerCId:    playerCId,
		PlayerDId:    playerDId,
		IsForfeit:    false,
		Set1:         "17-21",
		Set2:         "21-15",
		Set3:         "15-7",
		TournamentId: tournament.Id,
	}
	db := sqlite.NewInMemoryDB()
	tournamentRepository := sqlite.NewTournamentRepository(db)
	tournamentRepository.Create(tournament)
	matchRepository := sqlite.NewMatchRepository(db)
	matchRepository.Create(match)
	playerRatingRepository := sqlite.NewPlayerRatingRepository(db)
	ratingCalculator := vbratings.NewRatingCalculator(matchRepository, tournamentRepository, playerRatingRepository)

	// When
	totalCalculated := ratingCalculator.CalculateRatingsByYearFromMatches(2018)

	// Then
	assert.Equal(t, 4, totalCalculated)
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
	tournament := vbratings.Tournament{
		Id:     "383cea98af1f40ebbf6fdfd6deec7270",
		Gender: "female",
		Year:   2018,
	}
	match1 := vbratings.Match{
		Id:           "bb5ff2b66aba45998fea0d8c0dc8bf52",
		PlayerAId:    playerAId,
		PlayerBId:    playerBId,
		PlayerCId:    playerCId,
		PlayerDId:    playerDId,
		IsForfeit:    false,
		Set1:         "17-21",
		Set2:         "21-15",
		Set3:         "15-7",
		TournamentId: tournament.Id,
	}
	match2 := vbratings.Match{
		Id:           "f87317add92e452f877e6b43e31c0a16",
		PlayerAId:    playerAId,
		PlayerBId:    playerBId,
		PlayerCId:    playerCId,
		PlayerDId:    playerDId,
		IsForfeit:    false,
		Set1:         "17-21",
		Set2:         "21-15",
		Set3:         "15-7",
		TournamentId: tournament.Id,
	}
	db := sqlite.NewInMemoryDB()
	tournamentRepository := sqlite.NewTournamentRepository(db)
	tournamentRepository.Create(tournament)
	matchRepository := sqlite.NewMatchRepository(db)
	matchRepository.Create(match1)
	matchRepository.Create(match2)
	playerRatingRepository := sqlite.NewPlayerRatingRepository(db)
	ratingCalculator := vbratings.NewRatingCalculator(matchRepository, tournamentRepository, playerRatingRepository)

	// When
	totalCalculated := ratingCalculator.CalculateRatingsByYearFromMatches(2018)

	// Then
	assert.Equal(t, 4, totalCalculated)
	playerRatingA, _ := playerRatingRepository.GetPlayerRatingByYear(playerAId, 2018)
	playerRatingB, _ := playerRatingRepository.GetPlayerRatingByYear(playerBId, 2018)
	playerRatingC, _ := playerRatingRepository.GetPlayerRatingByYear(playerCId, 2018)
	playerRatingD, _ := playerRatingRepository.GetPlayerRatingByYear(playerDId, 2018)
	assertPlayerRating(t, playerRatingA, 1500, 1530, 2018)
	assertPlayerRating(t, playerRatingB, 1500, 1530, 2018)
	assertPlayerRating(t, playerRatingC, 1500, 1469, 2018)
	assertPlayerRating(t, playerRatingD, 1500, 1469, 2018)
	assert.Equal(t, 2, playerRatingA.TotalMatches)
}

func Test_RatingCalculator_CalculateRatingsByYear_SeededWithPreviousYearPlayerRating(t *testing.T) {
	// Given
	playerRating := vbratings.PlayerRating{
		PlayerId:   playerAId,
		Year:       2017,
		SeedRating: 1500,
		Rating:     1600,
	}
	tournament := vbratings.Tournament{
		Id:     "3f955b24ef804c198b9ddcfaf330ff8f",
		Gender: "male",
		Year:   2018,
	}
	match := vbratings.Match{
		Id:           "34d659c28edb4a8b94ea5c39dea32534",
		PlayerAId:    playerAId,
		PlayerBId:    playerBId,
		PlayerCId:    playerCId,
		PlayerDId:    playerDId,
		IsForfeit:    false,
		Set1:         "17-21",
		Set2:         "21-15",
		Set3:         "15-7",
		TournamentId: tournament.Id,
	}
	db := sqlite.NewInMemoryDB()
	tournamentRepository := sqlite.NewTournamentRepository(db)
	tournamentRepository.Create(tournament)
	matchRepository := sqlite.NewMatchRepository(db)
	matchRepository.Create(match)
	playerRatingRepository := sqlite.NewPlayerRatingRepository(db)
	playerRatingRepository.Create(playerRating)
	ratingCalculator := vbratings.NewRatingCalculator(matchRepository, tournamentRepository, playerRatingRepository)

	// When
	totalCalculated := ratingCalculator.CalculateRatingsByYearFromMatches(2018)

	// Then
	assert.Equal(t, 4, totalCalculated)
	playerRatingA, _ := playerRatingRepository.GetPlayerRatingByYear(playerAId, 2018)
	playerRatingB, _ := playerRatingRepository.GetPlayerRatingByYear(playerBId, 2018)
	playerRatingC, _ := playerRatingRepository.GetPlayerRatingByYear(playerCId, 2018)
	playerRatingD, _ := playerRatingRepository.GetPlayerRatingByYear(playerDId, 2018)
	assertPlayerRating(t, playerRatingA, 1600, 1611, 2018)
	assertPlayerRating(t, playerRatingB, 1500, 1516, 2018)
	assertPlayerRating(t, playerRatingC, 1500, 1486, 2018)
	assertPlayerRating(t, playerRatingD, 1500, 1486, 2018)
}

func Test_RatingCalculator_CalculateRatingsByYearFromTournamentResults_1TeamWith1EarnedFinish(t *testing.T) {
	// Given
	const year = 2018
	tournament := vbratings.Tournament{
		Id:   "3366b167a109496db63f43169e4ac1a7",
		Year: year,
	}

	result1 := vbratings.TournamentResult{
		Id:           "e57fd803a81349b69f0e7160fa13e919",
		Player1Id:    playerAId,
		Player2Id:    playerBId,
		EarnedFinish: 1,
		TournamentId: tournament.Id,
	}

	db := sqlite.NewInMemoryDB()
	tournamentRepository := sqlite.NewTournamentRepository(db)
	tournamentRepository.Create(tournament)
	tournamentRepository.AddTournamentResult(result1)
	matchRepository := sqlite.NewMatchRepository(db)
	playerRatingRepository := sqlite.NewPlayerRatingRepository(db)
	ratingCalculator := vbratings.NewRatingCalculator(
		matchRepository,
		tournamentRepository,
		playerRatingRepository,
	)

	// When
	totalCalculated := ratingCalculator.CalculateRatingsByYearFromTournamentResults(year)

	// Then
	assert.Equal(t, 2, totalCalculated)
	assertNewRating(t, playerRatingRepository, result1.Player1Id, 1545, year)
	assertNewRating(t, playerRatingRepository, result1.Player2Id, 1545, year)
}

func Test_RatingCalculator_CalculateRatingsByYearFromTournamentResults_2TeamsWith2EarnedFinishes(t *testing.T) {
	// Given
	const year = 2018
	tournament := vbratings.Tournament{
		Id:   "3366b167a109496db63f43169e4ac1a7",
		Year: year,
	}

	result1 := vbratings.TournamentResult{
		Id:           "e57fd803a81349b69f0e7160fa13e919",
		Player1Id:    playerAId,
		Player2Id:    playerBId,
		EarnedFinish: 1,
		TournamentId: tournament.Id,
	}

	result2 := vbratings.TournamentResult{
		Id:           "d80ad4fe3bc2447abb44388702d0f791",
		Player1Id:    playerCId,
		Player2Id:    playerDId,
		EarnedFinish: 2,
		TournamentId: tournament.Id,
	}

	db := sqlite.NewInMemoryDB()
	tournamentRepository := sqlite.NewTournamentRepository(db)
	tournamentRepository.Create(tournament)
	tournamentRepository.AddTournamentResult(result1)
	tournamentRepository.AddTournamentResult(result2)
	matchRepository := sqlite.NewMatchRepository(db)
	playerRatingRepository := sqlite.NewPlayerRatingRepository(db)
	ratingCalculator := vbratings.NewRatingCalculator(
		matchRepository,
		tournamentRepository,
		playerRatingRepository,
	)

	// When
	totalCalculated := ratingCalculator.CalculateRatingsByYearFromTournamentResults(year)

	// Then
	assert.Equal(t, 4, totalCalculated)
	assertNewRating(t, playerRatingRepository, result1.Player1Id, 1561, year)
	assertNewRating(t, playerRatingRepository, result1.Player2Id, 1561, year)

	assertNewRating(t, playerRatingRepository, result2.Player1Id, 1529, year)
	assertNewRating(t, playerRatingRepository, result2.Player2Id, 1529, year)
}

func assertNewRating(t *testing.T, playerRatingRepository vbratings.PlayerRatingRepository, playerId string, expectedRating int, expectedYear int) {
	playerRatingResult, err := playerRatingRepository.GetPlayerRatingByYear(playerId, expectedYear)
	require.NoError(t, err)
	assertPlayerRating(t, playerRatingResult, 1500, expectedRating, expectedYear)
}

func assertPlayerRating(t *testing.T, playerRating *vbratings.PlayerRating, expectedSeedRating int, expectedRating int, expectedYear int) {
	assert.Equal(t, expectedSeedRating, playerRating.SeedRating)
	assert.Equal(t, expectedRating, playerRating.Rating)
	assert.Equal(t, expectedYear, playerRating.Year)
}
