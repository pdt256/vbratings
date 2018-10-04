package app_test

import (
	"testing"

	"github.com/pdt256/vbratings/app"
	"github.com/pdt256/vbratings/sqlite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_PlayerRating_GetTopPlayerRatings(t *testing.T) {
	// Given
	configuration := app.NewConfiguration(sqlite.InMemorySharedDbPath)
	application := app.New(configuration)
	player1Id := "b9f0a4ed81b845b39d79d2020053ac66"
	player2Id := "7ed33e0b3aca4ec983825d9e1e08e667"
	err1 := application.PlayerRating.Create(player1Id, 2018, 1500, 2400, 1)
	err2 := application.PlayerRating.Create(player2Id, 2018, 1500, 2000, 1)
	err3 := application.Player.Create(player1Id, "Jane Doe", "female")
	err4 := application.Player.Create(player2Id, "Jane Smith", "female")
	require.NoError(t, err1)
	require.NoError(t, err2)
	require.NoError(t, err3)
	require.NoError(t, err4)

	// When
	playerAndRatings := application.PlayerRating.GetTopPlayerRatings(
		2018,
		"female",
		5)

	// Then
	actualTopPlayer := playerAndRatings[0]
	assert.Equal(t, player1Id, actualTopPlayer.PlayerId)
	assert.Equal(t, "Jane Doe", actualTopPlayer.Name)
	assert.Equal(t, 2400, actualTopPlayer.Rating)
}
