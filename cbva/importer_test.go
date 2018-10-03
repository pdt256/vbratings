package cbva_test

import (
	"os"
	"testing"

	"github.com/pdt256/vbratings/cbva"
	"github.com/pdt256/vbratings/pkg/uuid"
	"github.com/pdt256/vbratings/sqlite"
	"github.com/stretchr/testify/assert"
)

func Test_Importer_ImportTournamentResults(t *testing.T) {
	// Given
	db := sqlite.NewInMemoryDB()
	tournamentRepository := sqlite.NewTournamentRepository(db)
	playerRepository := sqlite.NewPlayerRepository(db)
	cbvaRepository := cbva.NewSqliteRepository(db)
	uuidGenerator := uuid.NewService()
	importer := cbva.NewImporter(
		tournamentRepository,
		playerRepository,
		cbvaRepository,
		uuidGenerator,
	)
	reader, _ := os.Open("./testdata/2018-09-23-marine-street-mens-aa.json")
	tournamentId := "A14CC0CB1B90719A"

	// When
	totalResults, totalPlayers := importer.ImportTournamentResults(reader, tournamentId)

	// Then
	actualTournamentResults := tournamentRepository.GetAllTournamentResults()
	assert.Equal(t, 15, len(actualTournamentResults))
	assert.Equal(t, 15, totalResults)
	assert.Equal(t, 30, totalPlayers)
	assert.Equal(t, tournamentId, actualTournamentResults[0].TournamentId)
}
