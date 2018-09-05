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
	importer := vbscraper.NewBvbInfoImporter(matchRepository, nil)
	matchesReader, _ := os.Open("./assets/2017-avp-manhattan-beach-mens-matches.html")

	// When
	totalImported := importer.ImportMatches(matchesReader)

	// Then
	assert.Equal(t, 159, totalImported)
}
