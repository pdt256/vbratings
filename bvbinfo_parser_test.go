package vbscraper_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/pdt256/vbscraper"
	"github.com/stretchr/testify/assert"
)

func Test_GetMatches(t *testing.T) {
	// Given
	file, _ := os.Open("./assets/2017-avp-manhattan-beach-mens-matches.html")

	// When
	matches := vbscraper.GetMatches(file)

	// Then
	assert.Equal(t, 159, len(matches))
	lastMatch := matches[158]
	assert.Equal(t, "5214", lastMatch.PlayerA.BvbId)
	assert.Equal(t, "Phil Dalhausser", lastMatch.PlayerA.Name)
	assert.Equal(t, "1931", lastMatch.PlayerB.BvbId)
	assert.Equal(t, "Nick Lucena", lastMatch.PlayerB.Name)
	assert.Equal(t, "13453", lastMatch.PlayerC.BvbId)
	assert.Equal(t, "Trevor Crabb", lastMatch.PlayerC.Name)
	assert.Equal(t, "1163", lastMatch.PlayerD.BvbId)
	assert.Equal(t, "Sean Rosenthal", lastMatch.PlayerD.Name)
	assert.False(t, lastMatch.IsForfeit)
	assert.Equal(t, "23-25", lastMatch.Set1)
	assert.Equal(t, "21-18", lastMatch.Set2)
	assert.Equal(t, "15-10", lastMatch.Set3)
	expectedLastMatch := "{PlayerA:{BvbId:5214 Name:Phil Dalhausser} PlayerB:{BvbId:1931 Name:Nick Lucena} PlayerC:{BvbId:13453 Name:Trevor Crabb} PlayerD:{BvbId:1163 Name:Sean Rosenthal} IsForfeit:false Set1:23-25 Set2:21-18 Set3:15-10}"
	assert.Equal(t, expectedLastMatch, fmt.Sprintf("%+v", lastMatch))
}

func Test_GetMatches_HandlesForfeit(t *testing.T) {
	// Given
	file, _ := os.Open("./assets/2015-avp-manhattan-beach-mens-matches.html")

	// When
	matches := vbscraper.GetMatches(file)

	// Then
	assert.Equal(t, 107, len(matches))
	forfeitMatch := matches[0]
	assert.True(t, forfeitMatch.IsForfeit)
	assert.Equal(t, "", forfeitMatch.Set1)
	assert.Equal(t, "", forfeitMatch.Set2)
	assert.Equal(t, "", forfeitMatch.Set3)
	expectedLastMatch := "{PlayerA:{BvbId:13513 Name:Juan Beltran} PlayerB:{BvbId:14187 Name:Zack Kweder} PlayerC:{BvbId:10935 Name:Alex Pepke} PlayerD:{BvbId:15591 Name:Drew Pitlik} IsForfeit:true Set1: Set2: Set3:}"
	assert.Equal(t, expectedLastMatch, fmt.Sprintf("%+v", forfeitMatch))
}

func Test_GetTournaments(t *testing.T) {
	// Given
	file, _ := os.Open("./assets/2017-avp-tournaments.html")

	// When
	tournaments := vbscraper.GetTournaments(file)

	// Then
	assert.Equal(t, 16, len(tournaments))
	assert.Equal(t, "3320", tournaments[0].BvbId)
	assert.Equal(t, "Huntington Beach", tournaments[0].Name)
	assert.Equal(t, "{BvbId:3335 Name:Chicago}", fmt.Sprintf("%+v", tournaments[15]))
}

func Test_GetSeasons(t *testing.T) {
	// Given
	file, _ := os.Open("./assets/all-seasons.html")

	// When
	seasons := vbscraper.GetSeasons(file)

	// Then
	assert.Equal(t, 269, len(seasons))
	assert.Equal(t, "3", seasons[0].AssocID)
	assert.Equal(t, "2019", seasons[0].Year)
	assert.Equal(t, "{AssocID:3 Year:2019}", fmt.Sprintf("%+v", seasons[0]))
}
