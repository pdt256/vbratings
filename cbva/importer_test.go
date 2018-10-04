package cbva_test

import (
	"os"
	"testing"

	"github.com/pdt256/vbratings/cbva"
	"github.com/pdt256/vbratings/pkg/uuid"
	"github.com/pdt256/vbratings/sqlite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	cbvaTournament := cbva.Tournament{
		Id:           "1CA73A1A527D6384",
		Date:         "09-30-2018",
		Rating:       "16U",
		Gender:       "Female",
		Location:     "Belmont Shore, Long Beach",
		TournamentId: "f6d480b080f044f88174ad3c7e1de669",
	}

	// When
	totalResults, totalPlayers := importer.ImportTournamentResults(reader, cbvaTournament)

	// Then
	actualCBVATournament, err := cbvaRepository.GetTournament(cbvaTournament.Id)
	require.NoError(t, err)
	assert.Equal(t, "1CA73A1A527D6384", actualCBVATournament.Id)
	assert.Equal(t, "09-30-2018", actualCBVATournament.Date)
	assert.Equal(t, "16U", actualCBVATournament.Rating)
	assert.Equal(t, "Female", actualCBVATournament.Gender)
	assert.Equal(t, "Belmont Shore, Long Beach", actualCBVATournament.Location)
	actualTournamentResults := tournamentRepository.GetAllTournamentResults()
	assert.Equal(t, 15, len(actualTournamentResults))
	assert.Equal(t, 15, totalResults)
	assert.Equal(t, 30, totalPlayers)
	assert.Equal(t, cbvaTournament.TournamentId, actualTournamentResults[0].TournamentId)
}
