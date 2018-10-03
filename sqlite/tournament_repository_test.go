package sqlite_test

import (
	"testing"

	"github.com/pdt256/vbratings"
	"github.com/pdt256/vbratings/sqlite"
	"github.com/stretchr/testify/assert"
)

func Test_TournamentRepository_GetAllTournamentResults(t *testing.T) {

	tournamentResult := vbratings.TournamentResult{
		Id:           "1c64ac801a4c427981397cdf62de98c6",
		TournamentId: "9e1732da1cf64446aeecd7c468458efd",
		Player1Id:    "6d4d0b4365044a80af3c7788f6bafabc",
		Player2Id:    "0193c5d0185b484daba21aae35c78e8d",
		EarnedFinish: 1,
	}
	db := sqlite.NewInMemoryDB()
	repository := sqlite.NewTournamentRepository(db)
	repository.AddTournamentResult(tournamentResult)

	// When
	tournamentResults := repository.GetAllTournamentResults()

	// Then
	assert.Equal(t, 1, len(tournamentResults))
	assert.Equal(t, "1c64ac801a4c427981397cdf62de98c6", tournamentResults[0].Id)
	assert.Equal(t, "9e1732da1cf64446aeecd7c468458efd", tournamentResults[0].TournamentId)
	assert.Equal(t, "6d4d0b4365044a80af3c7788f6bafabc", tournamentResults[0].Player1Id)
	assert.Equal(t, "0193c5d0185b484daba21aae35c78e8d", tournamentResults[0].Player2Id)
	assert.Equal(t, 1, tournamentResults[0].EarnedFinish)
}
