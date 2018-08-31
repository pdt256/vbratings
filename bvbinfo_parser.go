package vbscraper

import (
	"io"
	"io/ioutil"
	"regexp"
)

var tournamentRegexp = regexp.MustCompile(`<a href="Tournament.asp\?ID=(\d+)">(.+?)</a>`)
var seasonRegexp = regexp.MustCompile(`<a href="Season.asp\?AssocID=(\d+)&Year=(\d+)">`)

var playerExpression = `<a href="player.asp\?ID=(\d+)">([^<]+)</a>`
var matchRegexp = regexp.MustCompile(`(?m:<br>Match\s\d+:[^?]+` +
	playerExpression + `[^?]+` +
	playerExpression + `[^?]+` +
	playerExpression + `[^?]+` +
	playerExpression + `[^?]+\)` +
	`(?:` + `(?:\sby\s(Forfeit))` + `|` + `(?:\s(\d+-\d+),\s(\d+-\d+)(?:,\s(\d+-\d+))?\s\((\d+:\d+)\))` + `)` +
	`)`)

type Season struct {
	AssocID string
	Year    string
}

type Tournament struct {
	BvbId string
	Name  string
}

type Match struct {
	PlayerA   Player
	PlayerB   Player
	PlayerC   Player
	PlayerD   Player
	IsForfeit bool
}

type Player struct {
	BvbId string
	Name  string
}

func GetMatches(reader io.Reader) []Match {
	bytes, _ := ioutil.ReadAll(reader)
	regexMatches := matchRegexp.FindAllStringSubmatch(string(bytes), -1)

	var matches []Match
	for _, value := range regexMatches {
		//fmt.Printf("%#v", value[1:])
		isForfeit := value[9] == "Forfeit"
		matches = append(matches, Match{
			Player{value[1], value[2]},
			Player{value[3], value[4]},
			Player{value[5], value[6]},
			Player{value[7], value[8]},
			isForfeit,
		})
	}

	return matches
}

func GetTournaments(reader io.Reader) []Tournament {
	bytes, _ := ioutil.ReadAll(reader)
	regexMatches := tournamentRegexp.FindAllStringSubmatch(string(bytes), -1)

	var tournaments []Tournament
	for _, value := range regexMatches {
		tournaments = append(tournaments, Tournament{value[1], value[2]})
	}

	return tournaments
}

func GetSeasons(reader io.Reader) []Season {
	bytes, _ := ioutil.ReadAll(reader)
	regexMatches := seasonRegexp.FindAllStringSubmatch(string(bytes), -1)

	var seasons []Season
	for _, value := range regexMatches {
		seasons = append(seasons, Season{value[1], value[2]})
	}

	return seasons
}
