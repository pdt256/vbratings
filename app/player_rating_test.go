package app_test

import (
	"testing"

	"github.com/pdt256/vbratings/app"
	"github.com/stretchr/testify/assert"
)

func Test_PlayerRating_GetTopPlayerRatings(t *testing.T) {
	// Given
	dbPath := ":memory:"
	configuration := app.NewConfiguration(dbPath)
	application := app.New(configuration)
	player1Id := "b9f0a4ed-81b8-45b3-9d79-d2020053ac66"
	player2Id := "7ed33e0b-3aca-4ec9-8382-5d9e1e08e667"
	application.PlayerRating.Create(player1Id, 2018, "female", 1500, 2400, 1)
	application.PlayerRating.Create(player2Id, 2018, "female", 1500, 2000, 1)
	application.Player.Create(player1Id, "Jane Doe", "")
	application.Player.Create(player2Id, "Jane Smith", "")

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
