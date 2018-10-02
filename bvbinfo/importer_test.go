package bvbinfo_test

import (
	"os"
	"testing"

	"github.com/pdt256/vbratings/bvbinfo"
	"github.com/pdt256/vbratings/pkg/uuid"
	"github.com/pdt256/vbratings/sqlite"
	"github.com/stretchr/testify/assert"
)

func Test_Importer_ImportMatches(t *testing.T) {
	// Given
	db := sqlite.NewInMemoryDB()
	matchRepository := sqlite.NewMatchRepository(db)
	bvbInfoRepository := bvbinfo.NewRepositoryWithCaching(db)
	playerRepository := sqlite.NewPlayerRepository(db)
	uuidGenerator := uuid.NewService()
	importer := bvbinfo.NewImporter(
		matchRepository,
		playerRepository,
		bvbInfoRepository,
		uuidGenerator,
	)
	matchesReader, _ := os.Open("./testdata/2017-avp-manhattan-beach-mens-matches.html")

	// When
	totalMatches, totalPlayers := importer.ImportMatches(matchesReader)

	// Then
	actualMatches := matchRepository.GetAllMatchesByYear(2017)
	assert.Equal(t, 159, len(actualMatches))
	assert.Equal(t, 159, totalMatches)
	assert.Equal(t, 260, totalPlayers)
}
