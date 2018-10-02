package sqlite_test

import (
	"fmt"
	"testing"

	"github.com/pdt256/vbratings"
	"github.com/pdt256/vbratings/sqlite"
	"github.com/stretchr/testify/assert"
)

var (
	playerAId = "98556e2665224abc99c2d07d621befa7"
	playerBId = "7a30fab9631a442d83b70c9bf1293be8"
	playerCId = "91f67d94a9a54c91b9f0ee0efc497c28"
	playerDId = "ee655d72d148459ca10d05cce939bcab"
)

func Test_MatchRepository_CreateAndFindForfeit(t *testing.T) {
	// Given
	db := sqlite.NewInMemoryDB()
	matchRepository := sqlite.NewMatchRepository(db)
	repository := matchRepository
	match := vbratings.Match{
		PlayerAId: playerAId,
		PlayerBId: playerBId,
		PlayerCId: playerCId,
		PlayerDId: playerDId,
		IsForfeit: true,
		Year:      2018,
	}
	id := "123-abc"

	// When
	repository.Create(match, id)

	// Then
	actualMatch := repository.Find(id)
	assert.Equal(t, playerAId, actualMatch.PlayerAId)
	assert.Equal(t, playerBId, actualMatch.PlayerBId)
	assert.Equal(t, playerCId, actualMatch.PlayerCId)
	assert.Equal(t, playerDId, actualMatch.PlayerDId)
	assert.True(t, actualMatch.IsForfeit)
	assert.Equal(t, "", actualMatch.Set1)
	assert.Equal(t, "", actualMatch.Set2)
	assert.Equal(t, "", actualMatch.Set3)
	assert.Equal(t, 2018, actualMatch.Year)
}

func Test_MatchRepository_CreateAndFind3SetMatch(t *testing.T) {
	// Given
	db := sqlite.NewInMemoryDB()
	repository := sqlite.NewMatchRepository(db)
	match := vbratings.Match{
		PlayerAId: playerAId,
		PlayerBId: playerBId,
		PlayerCId: playerCId,
		PlayerDId: playerDId,
		IsForfeit: false,
		Set1:      "17-21",
		Set2:      "21-15",
		Set3:      "15-7",
	}
	id := "123-abc"

	// When
	repository.Create(match, id)

	// Then
	actualMatch := repository.Find(id)
	assert.Equal(t, playerAId, actualMatch.PlayerAId)
	assert.Equal(t, playerBId, actualMatch.PlayerBId)
	assert.Equal(t, playerCId, actualMatch.PlayerCId)
	assert.Equal(t, playerDId, actualMatch.PlayerDId)
	assert.Equal(t, "17-21", actualMatch.Set1)
	assert.Equal(t, "21-15", actualMatch.Set2)
	assert.Equal(t, "15-7", actualMatch.Set3)
	assert.False(t, actualMatch.IsForfeit)
}

func Test_MatchRepository_GetAllPlayerIds(t *testing.T) {
	// Given
	match := vbratings.Match{
		PlayerAId: playerAId,
		PlayerBId: playerBId,
		PlayerCId: playerCId,
		PlayerDId: playerDId,
		IsForfeit: false,
		Set1:      "17-21",
		Set2:      "21-15",
		Set3:      "15-7",
	}
	db := sqlite.NewInMemoryDB()
	repository := sqlite.NewMatchRepository(db)
	repository.Create(match, "a5ddebfd-dc00-4596-b39d-cb1c68f378df")

	// When
	actualPlayerIds := repository.GetAllPlayerIds()

	// Then
	expectedPlayerIds := []string{
		playerBId,
		playerCId,
		playerAId,
		playerDId,
	}
	assert.Equal(t, fmt.Sprintf("%+v", expectedPlayerIds), fmt.Sprintf("%+v", actualPlayerIds))
}
