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
	application.PlayerRating.Create(1, 2018, "female", 1500, 2400, 1)
	application.PlayerRating.Create(2, 2018, "female", 1500, 2000, 1)
	application.Player.Create(1, "Jane Doe", "")
	application.Player.Create(2, "Jane Smith", "")

	// When
	playerAndRatings := application.PlayerRating.GetTopPlayerRatings(
		2018,
		"female",
		5)

	// Then
	actualTopPlayer := playerAndRatings[0]
	assert.Equal(t, 1, actualTopPlayer.PlayerId)
	assert.Equal(t, "Jane Doe", actualTopPlayer.Name)
	assert.Equal(t, 2400, actualTopPlayer.Rating)
}
