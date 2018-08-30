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
	matchesFile, _ := os.Open("./assets/2017-avp-manhattan-beach-mens-matches.html")

	// When
	matches := vbscraper.GetMatches(matchesFile)

	// Then
	assert.Equal(t, 156, len(matches))
	expectedMatch0 := "{PlayerA:{Id:7376 Name:Matt Motter} PlayerB:{Id:2037 Name:Matt Olson} PlayerC:{Id:17858 Name:Parker Anderson} PlayerD:{Id:17468 Name:Kailum Rinaldi}}"
	expectedMatch155 := "{PlayerA:{Id:5214 Name:Phil Dalhausser} PlayerB:{Id:1931 Name:Nick Lucena} PlayerC:{Id:13453 Name:Trevor Crabb} PlayerD:{Id:1163 Name:Sean Rosenthal}}"
	assert.Equal(t, expectedMatch0, fmt.Sprintf("%+v", matches[0]))
	assert.Equal(t, expectedMatch155, fmt.Sprintf("%+v", matches[155]))
}
