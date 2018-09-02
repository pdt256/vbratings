package vbscraper

import (
	"io"
	"io/ioutil"
	"regexp"
	"strconv"
)

var tournamentRegexp = regexp.MustCompile(`<a href="Tournament.asp\?ID=(\d+)">(.+?)</a>`)
var seasonRegexp = regexp.MustCompile(`<a href="Season.asp\?AssocID=(\d+)&Year=(\d+)">`)

var playerExpression = `<a href="player.asp\?ID=(\d+)">[^<]+</a>`
var matchRegexp = regexp.MustCompile(`(?m:<br>Match\s\d+:[^?]+` +
	playerExpression + `[^?]+` +
	playerExpression + `[^?]+` +
	playerExpression + `[^?]+` +
	playerExpression + `[^?]+\)` +
	`(?:` + `(?:\sby\s(Forfeit))` + `|` + `(?:[^?]+(retired))` + `|` + `(?:\s(\d+-\d+),\s(\d+-\d+)(?:,\s(\d+-\d+))?` + `\s\((\d+:\d+)\))` + `)` +
	`)`)

var tournamentInfoRegexp = regexp.MustCompile(`(?m:clsTournHeader[^<]+<BR>\s+[^,]+,\s([^\r|\n]+))`)

type Season struct {
	AssocID string
	Year    string
}

type Tournament struct {
	BvbId string
	Name  string
}

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
}

func GetMatches(reader io.Reader) []Match {
	bytes, _ := ioutil.ReadAll(reader)
	tournamentInfoMatches := tournamentInfoRegexp.FindAllStringSubmatch(string(bytes), -1)

	var year int
	if len(tournamentInfoMatches) > 0 {
		year, _ = strconv.Atoi(tournamentInfoMatches[0][1])
	}

	regexMatches := matchRegexp.FindAllStringSubmatch(string(bytes), -1)

	var matches []Match
	for _, value := range regexMatches {
		//fmt.Printf("%#v", value[1:])
		isForfeit := value[5] == "Forfeit"
		isRetired := value[6] == "retired"

		matches = append(matches, Match{
			value[1],
			value[2],
			value[3],
			value[4],
			isForfeit || isRetired,
			value[7],
			value[8],
			value[9],
			year,
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
