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
	id := "b3a00549f00d4c3e880349526e9ec1e8"
	application.Player.Create(id, "Jane Doe", "http://example.com/1.jpg")

	// When
	player, err := application.Player.GetPlayer(id)

	// Then
	require.NoError(t, err)
	assert.Equal(t, id, player.Id)
	assert.Equal(t, "Jane Doe", player.Name)
	assert.Equal(t, "http://example.com/1.jpg", player.ImgUrl)
}
