package vbratings

import (
	"strings"
)

type Match struct {
	PlayerAId string
	PlayerBId string
	PlayerCId string
	PlayerDId string
	IsForfeit bool
	Set1      string
	Set2      string
	Set3      string
	Year      int
	Gender    string
}

type MatchRepository interface {
	Create(match Match, id string)
	Find(id string) *Match
	GetAllPlayerIds() []string
	GetAllMatchesByYear(year int) []Match
}

func GenderFromString(input string) string {
	if strings.ToLower(input) == "female" {
		return "female"
	}

	return "male"
}
