package vbscraper

import (
	"fmt"
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

var playerNameRegexp = regexp.MustCompile(`(?m:clsPlayerName">([^<]+)</td>)`)

type Season struct {
	AssocID string
	Year    string
}

type Tournament struct {
	BvbId string
	Name  string
}

type Match struct {
	PlayerAId int
	PlayerBId int
	PlayerCId int
	PlayerDId int
	IsForfeit bool
	Set1      string
	Set2      string
	Set3      string
	Year      int
}

type Player struct {
	BvbId  int
	Name   string
	ImgUrl string
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
		playerAId, _ := strconv.Atoi(value[1])
		playerBId, _ := strconv.Atoi(value[2])
		playerCId, _ := strconv.Atoi(value[3])
		playerDId, _ := strconv.Atoi(value[4])

		isForfeit := value[5] == "Forfeit"
		isRetired := value[6] == "retired"

		set1 := value[7]
		set2 := value[8]
		set3 := value[9]

		matches = append(matches, Match{
			playerAId,
			playerBId,
			playerCId,
			playerDId,
			isForfeit || isRetired,
			set1,
			set2,
			set3,
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

func GetPlayer(reader io.Reader, playerId int) Player {
	var name string

	bytes, _ := ioutil.ReadAll(reader)
	nameMatch := playerNameRegexp.FindStringSubmatch(string(bytes))
	if len(nameMatch) > 0 {
		name = nameMatch[1]
	}

	imgUrl := fmt.Sprintf("http://bvbinfo.com/images/photos/%d.jpg", playerId)

	return Player{
		BvbId:  playerId,
		Name:   name,
		ImgUrl: imgUrl,
	}
}
