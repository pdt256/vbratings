package vbratings

import (
	"regexp"
	"strings"
)

type TournamentResult struct {
	Player1Name  string
	Player2Name  string
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
