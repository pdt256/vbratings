package cbva_test

import (
	"os"
	"testing"

	"github.com/pdt256/vbratings/cbva"
	"github.com/pdt256/vbratings/sqlite"
	"github.com/stretchr/testify/assert"
)

func Test_Importer_ImportMatches(t *testing.T) {
	// Given
	db := sqlite.NewInMemoryDB()
	tournamentRepository := sqlite.NewTournamentRepository(db)
	importer := cbva.NewImporter(tournamentRepository, nil)
	reader, _ := os.Open("./testdata/2018-09-23-marine-street-mens-aa.json")

	// When
	totalImported := importer.ImportTournamentResults(reader)

	// Then
	actualTournamentResults := tournamentRepository.GetAllTournamentResults()
	assert.Equal(t, 15, len(actualTournamentResults))
	assert.Equal(t, 15, totalImported)
}
