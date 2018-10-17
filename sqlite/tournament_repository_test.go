package sqlite_test

import (
	"testing"

	"github.com/pdt256/vbratings"
	"github.com/pdt256/vbratings/sqlite"
	"github.com/stretchr/testify/assert"
)

const (
	tournamentAId       = "d2cbb5464c2e4bf382464be2cd0152be"
	tournamentResultAId = "8b9c8925e4d943b3986975630c23ca99"
	tournamentResultBId = "aaf0d24b5a134dc5984a284f96375caa"
)

func Test_TournamentRepository_GetAllTournamentResults(t *testing.T) {

	tournamentResult := vbratings.TournamentResult{
		Id:           tournamentResultAId,
		TournamentId: tournamentAId,
		Player1Id:    playerAId,
		Player2Id:    playerBId,
		EarnedFinish: 1,
	}
	db := sqlite.NewInMemoryDB()
	repository := sqlite.NewTournamentRepository(db)
	repository.AddTournamentResult(tournamentResult)

	// When
	tournamentResults := repository.GetAllTournamentResults()

	// Then
	assert.Equal(t, 1, len(tournamentResults))
	assert.Equal(t, tournamentResultAId, tournamentResults[0].Id)
	assert.Equal(t, tournamentAId, tournamentResults[0].TournamentId)
	assert.Equal(t, playerAId, tournamentResults[0].Player1Id)
	assert.Equal(t, playerBId, tournamentResults[0].Player2Id)
	assert.Equal(t, 1, tournamentResults[0].EarnedFinish)
}

func Test_TournamentRepository_GetAllTournamentsAndResultsByYear(t *testing.T) {
	const year = 2018

	tournament := vbratings.Tournament{
		Id:   tournamentAId,
		Year: year,
	}
	result1st := vbratings.TournamentResult{
		Id:           tournamentResultAId,
		TournamentId: tournamentAId,
		Player1Id:    playerAId,
		Player2Id:    playerBId,
		EarnedFinish: 1,
	}
	result2nd := vbratings.TournamentResult{
		Id:           tournamentResultBId,
		TournamentId: tournamentAId,
		Player1Id:    playerCId,
		Player2Id:    playerDId,
		EarnedFinish: 2,
	}
	db := sqlite.NewInMemoryDB()
	repository := sqlite.NewTournamentRepository(db)
	repository.Create(tournament)
	repository.AddTournamentResult(result2nd)
	repository.AddTournamentResult(result1st)

	// When
	tournamentsAndResults := repository.GetAllTournamentsAndResultsByYear(year)

	// Then
	assert.Equal(t, 1, len(tournamentsAndResults))
	tournamentAndResults := tournamentsAndResults[0]
	assert.Equal(t, tournamentAId, tournamentAndResults.Tournament.Id)
	assert.Equal(t, tournamentResultAId, tournamentAndResults.Results[0].Id)
	assert.Equal(t, year, tournamentAndResults.Tournament.Year)
	assert.Equal(t, tournamentAId, tournamentAndResults.Tournament.Id)
	assert.Equal(t, tournamentResultBId, tournamentAndResults.Results[1].Id)
	assert.Equal(t, year, tournamentAndResults.Tournament.Year)
}
