package vbscraper_test

import (
	"fmt"
	"log"
	"os"
	"sync"
	"testing"

	"github.com/pdt256/vbscraper"
	"github.com/stretchr/testify/assert"
)

func Test_BvbInfo_ImportMatches(t *testing.T) {
	// Given
	dbPath := "./_data/vb_test.db"
	os.Remove(dbPath)
	inMemoryRepository := &inMemoryMatchRepository{}
	importer := vbscraper.NewBvbInfoImporter(inMemoryRepository)
	matchesReader, _ := os.Open("./assets/2017-avp-manhattan-beach-mens-matches.html")

	// When
	totalImported := importer.ImportMatches(matchesReader)

	// Then
	assert.Equal(t, 156, totalImported)
	assert.Equal(t, 156, inMemoryRepository.TotalMatches())
}

type inMemoryMatchRepository struct {
	matches sync.Map
}

func (r *inMemoryMatchRepository) Create(match vbscraper.Match, id string) error {
	r.matches.Store(id, &match)
	return nil
}

func (r *inMemoryMatchRepository) Find(id string) *vbscraper.Match {
	match, ok := r.matches.Load(id)
	if !ok {
		log.Fatal("match not found")
	}

	return match.(*vbscraper.Match)
}

func (r *inMemoryMatchRepository) TotalMatches() interface{} {
	totalMatches := 0
	r.matches.Range(func(k, v interface{}) bool {
		totalMatches++
		return true
	})

	return totalMatches
}

func (r *inMemoryMatchRepository) MatchExists(match *vbscraper.Match) bool {
	wasFound := false
	r.matches.Range(func(k, v interface{}) bool {
		value := v.(*vbscraper.Match)
		if fmt.Sprintf("%+v", value) == fmt.Sprintf("%+v", match) {
			wasFound = true
			return false
		}
		return true
	})

	return wasFound
}
