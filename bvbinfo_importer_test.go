package vbscraper_test

import (
	"os"
	"testing"

	"github.com/pdt256/vbscraper"
	"github.com/stretchr/testify/assert"
)

func Test_Bvbinfo_ImportMatches(t *testing.T) {
	// Given
	db := NewInMemoryDB()
	matchRepository := vbscraper.NewSqliteMatchRepository(db)
	matchRepository.InitDB()
	importer := vbscraper.NewBvbInfoImporter(matchRepository, nil)
	matchesReader, _ := os.Open("./assets/2017-avp-manhattan-beach-mens-matches.html")

	// When
	totalImported := importer.ImportMatches(matchesReader)

	// Then
	actualMatches := matchRepository.GetAllMatchesByYear(2017)
	assert.Equal(t, 159, len(actualMatches))
	assert.Equal(t, 159, totalImported)
}
