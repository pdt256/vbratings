package vbscraper_test

import (
	"os"
	"testing"

	"github.com/pdt256/vbscraper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_MatchRepository_CreateAndFind(t *testing.T) {
	// Given
	dbPath := "./_data/vb_test.db"
	os.Remove(dbPath)
	repository := vbscraper.NewMatchRepository(dbPath)
	repository.InitDB()
	match := *vbscraper.NewMatch(
		*vbscraper.NewPlayer("1", "John"),
		*vbscraper.NewPlayer("2", "James"),
		*vbscraper.NewPlayer("3", "Jeremy"),
		*vbscraper.NewPlayer("4", "Johnathan"),
	)
	id := "123-abc"

	// When
	err := repository.Create(match, id)

	// Then
	require.NoError(t, err)
	actualMatch := repository.Find(id)
	assert.Equal(t, "1", actualMatch.PlayerA.BvbId)
	assert.Equal(t, "2", actualMatch.PlayerB.BvbId)
	assert.Equal(t, "3", actualMatch.PlayerC.BvbId)
	assert.Equal(t, "4", actualMatch.PlayerD.BvbId)
}
