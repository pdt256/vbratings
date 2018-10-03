package bvbinfo

import (
	"io"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var tournamentRegexp = regexp.MustCompile(`<a href="Tournament.asp\?ID=(\d+)">(.+?)</a>`)
var seasonRegexp = regexp.MustCompile(`<a href="Season.asp\?AssocID=(\d+)&Year=(\d+)">`)

var playerExpression = `<a href="player.asp\?ID=(\d+)">([^<]+)</a>`
var matchRegexp = regexp.MustCompile(`(?m:<br>Match\s\d+:[^?]+` +
	playerExpression + `[^?]+` +
	playerExpression + `[^?]+` +
	playerExpression + `[^?]+` +
	playerExpression + `[^?]+\)` +
	`(?:` + `(?:\sby\s(Forfeit))` + `|` + `(?:[^?]+(retired))` + `|` + `(?:\s(\d+-\d+),\s(\d+-\d+)(?:,\s(\d+-\d+))?` + `\s\((\d+:\d+)\))` + `)` +
	`)`)

var tournamentInfoRegexp = regexp.MustCompile(`(?m:clsTournHeader[^<]+<BR>\s+[^,]+,\s([^\r|\n]+))`)
var tournamentGenderRegexp = regexp.MustCompile(`(?m:clsTournHeader[^>]+>\s+([^\s]+)\s)`)

type Season struct {
	AssocID string
	Year    string
}

type Tournament struct {
	Id   int
	Name string
}

type Match struct {
	PlayerA   Player
	PlayerB   Player
	PlayerC   Player
	PlayerD   Player
	IsForfeit bool
	Set1      string
	Set2      string
	Set3      string
	Year      int
	Gender    string
}

type Player struct {
	Id   int
	Name string
}

func GetMatches(reader io.Reader) []Match {
	bytes, _ := ioutil.ReadAll(reader)
	body := string(bytes)
	tournamentInfoMatches := tournamentInfoRegexp.FindAllStringSubmatch(body, -1)

	var year int
	if len(tournamentInfoMatches) > 0 {
		year, _ = strconv.Atoi(tournamentInfoMatches[0][1])
	}

	tournamentGenderMatches := tournamentGenderRegexp.FindAllStringSubmatch(body, -1)
	var gender string
	if len(tournamentGenderMatches) > 0 {
		gender = tournamentGenderMatches[0][1]
	}
	gender = normalizeGender(gender)

	regexMatches := matchRegexp.FindAllStringSubmatch(body, -1)

	var matches []Match
	for _, value := range regexMatches {

		isForfeit := value[9] == "Forfeit"
		isRetired := value[10] == "retired"

		set1 := value[11]
		set2 := value[12]
		set3 := value[13]

		idA, _ := strconv.Atoi(value[1])
		idB, _ := strconv.Atoi(value[3])
		idC, _ := strconv.Atoi(value[5])
		idD, _ := strconv.Atoi(value[7])

		matches = append(matches, Match{
			PlayerA: Player{
				Id:   idA,
				Name: value[2],
			},
			PlayerB: Player{
				Id:   idB,
				Name: value[4],
			},
			PlayerC: Player{
				Id:   idC,
				Name: value[6],
			},
			PlayerD: Player{
				Id:   idD,
				Name: value[8],
			},
			IsForfeit: isForfeit || isRetired,
			Set1:      set1,
			Set2:      set2,
			Set3:      set3,
			Year:      year,
			Gender:    gender,
		})
	}

	return matches
}

func normalizeGender(input string) string {
	lowerInput := strings.ToLower(input)
	if lowerInput == "women's" || lowerInput == "female" {
		return "female"
	}

	return "male"
}

func GetTournaments(reader io.Reader) []Tournament {
	bytes, _ := ioutil.ReadAll(reader)
	regexMatches := tournamentRegexp.FindAllStringSubmatch(string(bytes), -1)

	var tournaments []Tournament
	for _, value := range regexMatches {
		id, _ := strconv.Atoi(value[1])
		tournaments = append(tournaments, Tournament{id, value[2]})
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
