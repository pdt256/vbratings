package app_test

import (
	"testing"

	"github.com/pdt256/vbratings/app"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Player_GetPlayer(t *testing.T) {
	// Given
	dbPath := ":memory:"
	configuration := app.NewConfiguration(dbPath)
	application := app.New(configuration)
	application.Player.Create(1, "Jane Doe", "")

	// When
	player, err := application.Player.GetPlayer(1)

	// Then
	require.NoError(t, err)
	assert.Equal(t, "Jane Doe", player.Name)
}
