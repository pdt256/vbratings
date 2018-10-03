package vbratings

import (
	"regexp"
	"strings"
)

type Tournament struct {
	Id       string
	Date     string
	Rating   string
	Gender   string
	Location string
}

type TournamentResult struct {
	Id           string
	TournamentId string
	Player1Id    string
	Player2Id    string
	EarnedFinish int
}

type TournamentRepository interface {
	AddTournament(tournament Tournament)
	AddTournamentResult(tournamentResult TournamentResult)
	GetAllTournamentResults() []TournamentResult
}

var slugRegexp = regexp.MustCompile("[^a-z0-9]+")

func GetPlayerNameSlug(s string) string {
	return strings.Trim(slugRegexp.ReplaceAllString(strings.ToLower(s), "-"), "-")
}
