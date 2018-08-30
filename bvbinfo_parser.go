package vbscraper

import (
	"io"
	"io/ioutil"
	"regexp"
)

type season struct {
	AssocID string
	Year    string
}

type tournament struct {
	BvbId string
	Name  string
}

type match struct {
	PlayerA player
	PlayerB player
	PlayerC player
	PlayerD player
}

type player struct {
	BvbId string
	Name  string
}

func NewMatch(playerA player, playerB player, playerC player, playerD player) *match {
	return &match{playerA, playerB, playerC, playerD}
}

func NewPlayer(id string, name string) *player {
	return &player{id, name}
}

func GetMatches(reader io.Reader) []match {
	playerRegex := `<a href="player.asp\?ID=(\d+)">([\w\s]+)</a>`
	re := regexp.MustCompile(playerRegex + `.*?` + playerRegex + `.*?` + playerRegex + `.*?` + playerRegex)
	bytes, _ := ioutil.ReadAll(reader)
	regexMatches := re.FindAllStringSubmatch(string(bytes), -1)

	var matches []match
	for _, value := range regexMatches {
		matches = append(matches, *NewMatch(
			player{value[1], value[2]},
			player{value[3], value[4]},
			player{value[5], value[6]},
			player{value[7], value[8]},
		))
	}

	return matches
}

func GetTournaments(reader io.Reader) []tournament {
	re := regexp.MustCompile(`<a href="Tournament.asp\?ID=(\d+)">([\w\s]+)</a>`)
	bytes, _ := ioutil.ReadAll(reader)
	regexMatches := re.FindAllStringSubmatch(string(bytes), -1)

	var tournaments []tournament
	for _, value := range regexMatches {
		tournaments = append(tournaments, tournament{value[1], value[2]})
	}

	return tournaments
}

func GetSeasons(reader io.Reader) []season {
	re := regexp.MustCompile(`<a href="Season.asp\?AssocID=(\d+)&Year=(\d+)">`)
	bytes, _ := ioutil.ReadAll(reader)
	regexMatches := re.FindAllStringSubmatch(string(bytes), -1)

	var seasons []season
	for _, value := range regexMatches {
		seasons = append(seasons, season{value[1], value[2]})
	}

	return seasons
}
