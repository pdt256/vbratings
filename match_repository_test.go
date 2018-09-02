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
		PlayerAId: "1",
		PlayerBId: "2",
		PlayerCId: "3",
		PlayerDId: "4",
		IsForfeit: true,
	}
	id := "123-abc"

	// When
	err := repository.Create(match, id)

	// Then
	require.NoError(t, err)
	actualMatch := repository.Find(id)
	assert.Equal(t, "1", actualMatch.PlayerAId)
	assert.Equal(t, "2", actualMatch.PlayerBId)
	assert.Equal(t, "3", actualMatch.PlayerCId)
	assert.Equal(t, "4", actualMatch.PlayerDId)
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
		PlayerAId: "1",
		PlayerBId: "2",
		PlayerCId: "3",
		PlayerDId: "4",
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
	assert.Equal(t, "1", actualMatch.PlayerAId)
	assert.Equal(t, "2", actualMatch.PlayerBId)
	assert.Equal(t, "3", actualMatch.PlayerCId)
	assert.Equal(t, "4", actualMatch.PlayerDId)
	assert.Equal(t, "17-21", actualMatch.Set1)
	assert.Equal(t, "21-15", actualMatch.Set2)
	assert.Equal(t, "15-7", actualMatch.Set3)
	assert.False(t, actualMatch.IsForfeit)
}
