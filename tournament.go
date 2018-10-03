package vbratings

import (
	"regexp"
	"strings"
)

type TournamentResult struct {
	Player1Id    string
	Player2Id    string
	EarnedFinish int
}

type TournamentRepository interface {
	AddTournamentResult(tournamentResult TournamentResult, id string)
	GetAllTournamentResults() []TournamentResult
}

var slugRegexp = regexp.MustCompile("[^a-z0-9]+")

func GetPlayerNameSlug(s string) string {
	return strings.Trim(slugRegexp.ReplaceAllString(strings.ToLower(s), "-"), "-")
}
