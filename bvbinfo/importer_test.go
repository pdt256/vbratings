package bvbinfo_test

import (
	"os"
	"testing"

	"github.com/pdt256/vbratings/bvbinfo"
	"github.com/pdt256/vbratings/pkg/uuid"
	"github.com/pdt256/vbratings/sqlite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Importer_ImportMatches(t *testing.T) {
	// Given
	tournamentId := "5bbdbfb027d14ef5ac464c21ee90ac1d"
	db := sqlite.NewInMemoryDB()
	tournamentRepository := sqlite.NewTournamentRepository(db)
	matchRepository := sqlite.NewMatchRepository(db)
	playerRepository := sqlite.NewPlayerRepository(db)
	bvbInfoRepository := bvbinfo.NewRepositoryWithCaching(db)
	uuidGenerator := uuid.NewStaticService([]string{tournamentId})
	importer := bvbinfo.NewImporter(
		tournamentRepository,
		matchRepository,
		playerRepository,
		bvbInfoRepository,
		uuidGenerator,
	)
	matchesReader, _ := os.Open("./testdata/2017-avp-manhattan-beach-mens-matches.html")

	// When
	bvbTournamentId := 3332
	totalMatches, totalPlayers := importer.ImportMatches(matchesReader, bvbTournamentId)

	// Then
	actualBvbInfoTournament, gtErr := bvbInfoRepository.GetTournament(bvbTournamentId)
	require.NoError(t, gtErr)
	assert.Equal(t, bvbTournamentId, actualBvbInfoTournament.Id)
	assert.Equal(t, "Men's AVP $112,500 Manhattan Beach Open", actualBvbInfoTournament.Name)
	assert.Equal(t, 2017, actualBvbInfoTournament.Year)
	assert.Equal(t, "August 17-20, 2017", actualBvbInfoTournament.Dates)
	assert.Equal(t, "male", actualBvbInfoTournament.Gender)
	bvbPlayer, gpErr := bvbInfoRepository.GetPlayer(5214)
	require.NoError(t, gpErr)
	assert.Equal(t, 5214, bvbPlayer.Id)
	actualTournament, err := tournamentRepository.GetTournament(tournamentId)
	require.NoError(t, err)
	assert.Equal(t, tournamentId, actualTournament.Id)
	assert.Equal(t, "Men's AVP $112,500 Manhattan Beach Open", actualTournament.Name)
	assert.Equal(t, "August 17-20, 2017", actualTournament.Date)
	assert.Equal(t, 2017, actualTournament.Year)
	assert.Equal(t, "male", actualTournament.Gender)
	actualMatches := matchRepository.GetAllMatchesByYear(2017)
	assert.Equal(t, 159, len(actualMatches))
	assert.Equal(t, 159, totalMatches)
	assert.Equal(t, 260, totalPlayers)
}
