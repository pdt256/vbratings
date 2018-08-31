package vbscraper_test

import (
	"os"
	"testing"

	"github.com/pdt256/vbscraper"
	"github.com/stretchr/testify/assert"
)

func Test_BvbInfo_ImportMatches(t *testing.T) {
	// Given
	inMemoryRepository := &vbscraper.InMemoryMatchRepository{}
	importer := vbscraper.NewBvbInfoImporter(inMemoryRepository)
	matchesReader, _ := os.Open("./assets/2017-avp-manhattan-beach-mens-matches.html")

	// When
	totalImported := importer.ImportMatches(matchesReader)

	// Then
	assert.Equal(t, 156, totalImported)
	assert.Equal(t, 156, inMemoryRepository.TotalMatches())
}
