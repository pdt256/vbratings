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
	Gender    Gender
}

type MatchRepository interface {
	Create(match Match, id string)
	Find(id string) *Match
	GetAllPlayerIds() []string
	GetAllMatchesByYear(year int) []Match
}

type Gender uint

func (gender Gender) String() string {
	names := [...]string{
		"Male",
		"Female",
		"Code",
	}

	return names[gender]
}

const (
	Male Gender = iota
	Female
)

func GenderFromString(input string) Gender {
	if strings.ToLower(input) == "female" {
		return Female
	}

	return Male
}
