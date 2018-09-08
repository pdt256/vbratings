package bvbinfo_test

import (
	"os"
	"strings"
	"testing"

	"github.com/pdt256/vbratings"
	"github.com/pdt256/vbratings/bvbinfo"
	"github.com/stretchr/testify/assert"
)

func Test_GetMatches_Handles3SetMatch(t *testing.T) {
	// Given
	input := `<br>Match 61: <b><a href="player.asp?ID=5214">Phil Dalhausser</a> / <a href="player.asp?ID=1931">Nick Lucena</a> (3)</b> def. <a href="player.asp?ID=13453">Trevor Crabb</a> / <a href="player.asp?ID=1163">Sean Rosenthal</a> (4) 23-25, 21-18, 15-10 (1:15)`

	// When
	matches := bvbinfo.GetMatches(strings.NewReader(input))

	// Then
	match := matches[0]
	assert.Equal(t, 5214, match.PlayerAId)
	assert.Equal(t, 1931, match.PlayerBId)
	assert.Equal(t, 13453, match.PlayerCId)
	assert.Equal(t, 1163, match.PlayerDId)
	assert.False(t, match.IsForfeit)
	assert.Equal(t, "23-25", match.Set1)
	assert.Equal(t, "21-18", match.Set2)
	assert.Equal(t, "15-10", match.Set3)
	assert.Equal(t, match.Gender, vbratings.Male)
}

func Test_GetMatches_Handles2ndSetRetired(t *testing.T) {
	// Given
	input := `<br>Match 12: <b><a href="player.asp?ID=16546">Andrea Abbiati</a> / <a href="player.asp?ID=10736">Tiziano Andreatta</a> Italy (31, Q27)</b> def. <a href="player.asp?ID=7145">Lombardo Ontiveros</a> / <a href="player.asp?ID=8011">Juan Virgen</a> Mexico (Q6) 26-24 retired (0:29)`

	// When
	matches := bvbinfo.GetMatches(strings.NewReader(input))

	// Then
	match := matches[0]
	assert.Equal(t, 16546, match.PlayerAId)
	assert.Equal(t, 10736, match.PlayerBId)
	assert.Equal(t, 7145, match.PlayerCId)
	assert.Equal(t, 8011, match.PlayerDId)
	assert.True(t, match.IsForfeit)
	assert.Equal(t, "", match.Set1)
	assert.Equal(t, "", match.Set2)
	assert.Equal(t, "", match.Set3)
}

func Test_GetMatches_Handles3rdSetRetired(t *testing.T) {
	// Given
	input := `<br>Match 30: <b><a href="player.asp?ID=7710">Leonardo Lunardi</a> / <a href="player.asp?ID=11131">Daniel Virkus</a> (Q18)</b> def. <a href="player.asp?ID=7960">Wayne Leever</a> / <a href="player.asp?ID=8777">Jared Tucker</a> (Q47) 21-16, 8-2 retired (0:32)`

	// When
	matches := bvbinfo.GetMatches(strings.NewReader(input))

	// Then
	match := matches[0]
	assert.Equal(t, 7710, match.PlayerAId)
	assert.Equal(t, 11131, match.PlayerBId)
	assert.Equal(t, 7960, match.PlayerCId)
	assert.Equal(t, 8777, match.PlayerDId)
	assert.True(t, match.IsForfeit)
	assert.Equal(t, "", match.Set1)
	assert.Equal(t, "", match.Set2)
	assert.Equal(t, "", match.Set3)
}

func Test_GetMatches_HandlesForfeit(t *testing.T) {
	// Given
	input := `<br>Match 2: <b><a href="player.asp?ID=13513">Juan Beltran</a> / <a href="player.asp?ID=14187">Zack Kweder</a> (Q32)</b> def. <a href="player.asp?ID=10935">Alex Pepke</a> / <a href="player.asp?ID=15591">Drew Pitlik</a> (Q33) by Forfeit`

	// When
	matches := bvbinfo.GetMatches(strings.NewReader(input))

	// Then
	match := matches[0]
	assert.Equal(t, 13513, match.PlayerAId)
	assert.Equal(t, 14187, match.PlayerBId)
	assert.Equal(t, 10935, match.PlayerCId)
	assert.Equal(t, 15591, match.PlayerDId)
	assert.True(t, match.IsForfeit)
	assert.Equal(t, "", match.Set1)
	assert.Equal(t, "", match.Set2)
	assert.Equal(t, "", match.Set3)
}

func Test_GetMatches_GetsYear(t *testing.T) {
	// Given
	file, _ := os.Open("./testdata/2018-fivb-gstaad-major-mens-matches.html")

	// When
	matches := bvbinfo.GetMatches(file)

	// Then
	match := matches[0]
	assert.Equal(t, 2018, match.Year)
}

func Test_GetMatches_GetsFemaleGender(t *testing.T) {
	// Given
	file, _ := os.Open("./testdata/2017-avp-manhattan-beach-womens-matches.html")

	// When
	matches := bvbinfo.GetMatches(file)

	// Then
	match := matches[0]
	assert.Equal(t, vbratings.Female, match.Gender)
}

func Test_GetMatches_ReturnsCorrectMatchCounts(t *testing.T) {
	// Given
	var tournaments = []struct {
		filePath             string
		expectedTotalMatches int
	}{
		{"./testdata/2014-avp-st-petersburg-mens-matches.html", 76},
		{"./testdata/2015-avp-manhattan-beach-mens-matches.html", 107},
		{"./testdata/2017-avp-manhattan-beach-mens-matches.html", 159},
		{"./testdata/2018-fivb-gstaad-major-mens-matches.html", 79},
	}

	for _, tt := range tournaments {
		t.Run(tt.filePath, func(t *testing.T) {
			// Given
			file, _ := os.Open(tt.filePath)

			// When
			matches := bvbinfo.GetMatches(file)

			// Then
			assert.Equal(t, tt.expectedTotalMatches, len(matches))
		})
	}
}

func Test_GetTournaments(t *testing.T) {
	// Given
	file, _ := os.Open("./testdata/2017-avp-tournaments.html")

	// When
	tournaments := bvbinfo.GetTournaments(file)

	// Then
	assert.Equal(t, 16, len(tournaments))
	assert.Equal(t, "3320", tournaments[0].BvbId)
	assert.Equal(t, "Huntington Beach", tournaments[0].Name)
}

func Test_GetSeasons(t *testing.T) {
	// Given
	file, _ := os.Open("./testdata/all-seasons.html")

	// When
	seasons := bvbinfo.GetSeasons(file)

	// Then
	assert.Equal(t, 269, len(seasons))
	assert.Equal(t, "3", seasons[0].AssocID)
	assert.Equal(t, "2019", seasons[0].Year)
}

func Test_GetPlayer(t *testing.T) {
	// Given
	file, _ := os.Open("./testdata/misty-may-player.html")

	// When
	player := bvbinfo.GetPlayer(file, 1256)

	// Then
	assert.Equal(t, 1256, player.BvbId)
	assert.Equal(t, "Misty May-Treanor", player.Name)
	assert.Equal(t, "http://bvbinfo.com/images/photos/1256.jpg", player.ImgUrl)
}
