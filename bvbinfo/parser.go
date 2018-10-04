package bvbinfo

import (
	"fmt"
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

var tournamentInfoRegexp = regexp.MustCompile(`(?m:clsTournHeader[^>]+>\s+([^<]+)<BR>\s+([^,]+, \d{4}))`)
var tournamentGenderRegexp = regexp.MustCompile(`(?m:clsTournHeader[^>]+>\s+([^\s]+)\s)`)

type Season struct {
	AssocID string
	Year    string
}

type TournamentLink struct {
	Id   int
	Name string
}

type Tournament struct {
	Id           int
	Name         string
	Year         int
	Dates        string
	Gender       string
	TournamentId string
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
}

type Player struct {
	Id       int
	Name     string
	ImgUrl   string
	PlayerId string
}

func GetMatches(reader io.Reader, tournamentId int) (Tournament, []Match) {
	bytes, _ := ioutil.ReadAll(reader)
	body := string(bytes)
	tournamentInfoMatches := tournamentInfoRegexp.FindAllStringSubmatch(body, -1)

	var dates string
	var name string
	var year int
	if len(tournamentInfoMatches) > 0 {
		name = tournamentInfoMatches[0][1]
		dates = tournamentInfoMatches[0][2]
		year, _ = strconv.Atoi(dates[len(dates)-4:])
	}

	tournamentGenderMatches := tournamentGenderRegexp.FindAllStringSubmatch(body, -1)
	var gender string
	if len(tournamentGenderMatches) > 0 {
		gender = tournamentGenderMatches[0][1]
	}
	gender = normalizeGender(gender)

	regexMatches := matchRegexp.FindAllStringSubmatch(body, -1)

	tournament := Tournament{
		Id:     tournamentId,
		Name:   name,
		Year:   year,
		Dates:  dates,
		Gender: gender,
	}

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
				Id:     idA,
				Name:   value[2],
				ImgUrl: getImgUrl(idA),
			},
			PlayerB: Player{
				Id:     idB,
				Name:   value[4],
				ImgUrl: getImgUrl(idB),
			},
			PlayerC: Player{
				Id:     idC,
				Name:   value[6],
				ImgUrl: getImgUrl(idC),
			},
			PlayerD: Player{
				Id:     idD,
				Name:   value[8],
				ImgUrl: getImgUrl(idD),
			},
			IsForfeit: isForfeit || isRetired,
			Set1:      set1,
			Set2:      set2,
			Set3:      set3,
		})
	}

	return tournament, matches
}

func getImgUrl(bvbId int) string {
	return fmt.Sprintf("http://bvbinfo.com/images/photos/%d.jpg", bvbId)
}

func normalizeGender(input string) string {
	lowerInput := strings.ToLower(input)
	if lowerInput == "women's" || lowerInput == "female" {
		return "female"
	}

	return "male"
}

func GetTournaments(reader io.Reader) []TournamentLink {
	bytes, _ := ioutil.ReadAll(reader)
	regexMatches := tournamentRegexp.FindAllStringSubmatch(string(bytes), -1)

	var tournamentLinks []TournamentLink
	for _, value := range regexMatches {
		id, _ := strconv.Atoi(value[1])
		tournamentLinks = append(tournamentLinks, TournamentLink{id, value[2]})
	}

	return tournamentLinks
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
