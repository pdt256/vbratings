package vbscraper_test

import (
	"os"
	"testing"

	"github.com/pdt256/vbscraper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_MatchRepository_CreateAndFindForfeit(t *testing.T) {
	// Given
	dbPath := "./_data/vb_test.db"
	os.Remove(dbPath)
	repository := vbscraper.NewSqliteMatchRepository(dbPath)
	repository.InitDB()
	match := vbscraper.Match{
		PlayerA:   vbscraper.Player{BvbId: "1", Name: "John"},
		PlayerB:   vbscraper.Player{BvbId: "2", Name: "James"},
		PlayerC:   vbscraper.Player{BvbId: "3", Name: "Jeremy"},
		PlayerD:   vbscraper.Player{BvbId: "4", Name: "Johnathan"},
		IsForfeit: true,
	}
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
	assert.True(t, actualMatch.IsForfeit)
	assert.Equal(t, "", actualMatch.Set1)
	assert.Equal(t, "", actualMatch.Set2)
	assert.Equal(t, "", actualMatch.Set3)
}

func Test_MatchRepository_CreateAndFind3SetMatch(t *testing.T) {
	// Given
	dbPath := "./_data/vb_test.db"
	os.Remove(dbPath)
	repository := vbscraper.NewSqliteMatchRepository(dbPath)
	repository.InitDB()
	match := vbscraper.Match{
		PlayerA:   vbscraper.Player{BvbId: "1", Name: "John"},
		PlayerB:   vbscraper.Player{BvbId: "2", Name: "James"},
		PlayerC:   vbscraper.Player{BvbId: "3", Name: "Jeremy"},
		PlayerD:   vbscraper.Player{BvbId: "4", Name: "Johnathan"},
		IsForfeit: false,
		Set1:      "17-21",
		Set2:      "21-15",
		Set3:      "15-7",
	}
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
	assert.Equal(t, "17-21", actualMatch.Set1)
	assert.Equal(t, "21-15", actualMatch.Set2)
	assert.Equal(t, "15-7", actualMatch.Set3)
	assert.False(t, actualMatch.IsForfeit)
}
