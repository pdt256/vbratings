package cbva_test

import (
	"os"
	"strings"
	"testing"

	"github.com/pdt256/vbratings/cbva"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_GetTournamentResults_Handles1Result(t *testing.T) {
	// Given
	input := `{"EarnedFinish":"1","EarnedRating":"AAA","EarnedRatingCodeID": 6,"Player1_Name":"JEFF STEFFENS","Player1_Rating":"AAA","Player2_Name":"Leo Williams","Player2_Rating":"AAA","TeamStatusCodeID": 1,"TournamentResultID":"83575","TournamentTeamID":"231829"}`

	// When
	tournamentResults := cbva.GetTournamentResults(strings.NewReader(input))

	// Then
	tournamentResult := tournamentResults[0]
	assert.Equal(t, "JEFF STEFFENS", tournamentResult.Player1.Name)
	assert.Equal(t, "Leo Williams", tournamentResult.Player2.Name)
	assert.Equal(t, 1, tournamentResult.EarnedFinish)
}

func Test_GetTournamentResults_Handles3Results(t *testing.T) {
	// Given
	input := `{"EarnedFinish":"1","EarnedRating":"AAA","EarnedRatingCodeID": 6,"Player1_Name":"JEFF STEFFENS","Player1_Rating":"AAA","Player2_Name":"Leo Williams","Player2_Rating":"AAA","TeamStatusCodeID": 1,"TournamentResultID":"83575","TournamentTeamID":"231829"}{"EarnedFinish":"2","EarnedRating":"AA","EarnedRatingCodeID": 5,"Player1_Name":"J.J. MOSOLF","Player1_Rating":"AA","Player2_Name":"max states","Player2_Rating":"AA","TeamStatusCodeID": 1,"TournamentResultID":"83581","TournamentTeamID":"231709"},{"EarnedFinish":"3","EarnedRating":"AA","EarnedRatingCodeID": 5,"Player1_Name":"Mark Raphael","Player1_Rating":"AA","Player2_Name":"Jamie Isaacs","Player2_Rating":"AA","TeamStatusCodeID": 1,"TournamentResultID":"83574","TournamentTeamID":"231677"}`

	// When
	tournamentResults := cbva.GetTournamentResults(strings.NewReader(input))

	// Then
	tournamentResult1 := tournamentResults[0]
	assert.Equal(t, "JEFF STEFFENS", tournamentResult1.Player1.Name)
	assert.Equal(t, "Leo Williams", tournamentResult1.Player2.Name)
	assert.Equal(t, 1, tournamentResult1.EarnedFinish)
	tournamentResult2 := tournamentResults[1]
	assert.Equal(t, "J.J. MOSOLF", tournamentResult2.Player1.Name)
	assert.Equal(t, "max states", tournamentResult2.Player2.Name)
	assert.Equal(t, 2, tournamentResult2.EarnedFinish)
	tournamentResult3 := tournamentResults[2]
	assert.Equal(t, "Mark Raphael", tournamentResult3.Player1.Name)
	assert.Equal(t, "Jamie Isaacs", tournamentResult3.Player2.Name)
	assert.Equal(t, 3, tournamentResult3.EarnedFinish)
}

func Test_GetTournamentResults_ReturnsCorrectTournamentResultCounts(t *testing.T) {
	// Given
	var tournaments = []struct {
		expectedTotalMatches int
		filePath             string
	}{
		{15, "2018-09-23-marine-street-mens-aa.json"},
		{6, "2018-09-01-hermosa-mens-open.json"},
	}

	for _, tt := range tournaments {
		t.Run(tt.filePath, func(t *testing.T) {
			// Given
			file, err := os.Open("./testdata/" + tt.filePath)
			require.NoError(t, err)

			// When
			tournamentResults := cbva.GetTournamentResults(file)

			// Then
			assert.Equal(t, tt.expectedTotalMatches, len(tournamentResults))
		})
	}
}

func Test_GetTournaments_Handles1Result(t *testing.T) {
	// Given
	input := `<div class=\"cbva-btn\"><a class=\"btn btn-info\" href=\"#\" data-id=\"1CA73A1A527D6384\" data-Date=\"09-30-2018\" data-Rating=\"16U\" data-Gender=\"Female\" data-TournamentDirector=\"Sinjin Smith\" data-Location=\"Belmont Shore, Long Beach\" data-toggle=\"modal\" data-target=\"#tournament-result\">View Results</a></div>`

	// When
	tournaments := cbva.GetTournaments(strings.NewReader(input))

	// Then
	tournament := tournaments[0]
	assert.Equal(t, "1CA73A1A527D6384", tournament.Id)
	assert.Equal(t, "09-30-2018", tournament.Date)
	assert.Equal(t, "16U", tournament.Rating)
	assert.Equal(t, "Female", tournament.Gender)
	assert.Equal(t, "Belmont Shore, Long Beach", tournament.Location)
}

func Test_GetTournaments_Handles10Results(t *testing.T) {
	// Given
	file, err := os.Open("./testdata/10-tournaments-list.json")
	require.NoError(t, err)

	// When
	tournaments := cbva.GetTournaments(file)

	// Then
	assert.Equal(t, 10, len(tournaments))
}
