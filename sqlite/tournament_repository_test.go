package sqlite_test

import (
	"fmt"
	"testing"

	"github.com/pdt256/vbratings"
	"github.com/pdt256/vbratings/sqlite"
	"github.com/stretchr/testify/assert"
)

func Test_TournamentRepository_GetAllTournamentResults(t *testing.T) {

	tournamentResult := vbratings.TournamentResult{
		Player1Id:    "john-doe",
		Player2Id:    "james-smith",
		EarnedFinish: 1,
	}
	db := sqlite.NewInMemoryDB()
	repository := sqlite.NewTournamentRepository(db)
	repository.AddTournamentResult(tournamentResult, "123-abc")

	// When
	tournamentResults := repository.GetAllTournamentResults()

	// Then
	assert.Equal(t, "[{john-doe james-smith 1}]", fmt.Sprintf("%v", tournamentResults))
}
