package vbscraper

import (
	"io"
	"io/ioutil"
	"regexp"
)

type Season struct {
	AssocID string
	Year    string
}

type Tournament struct {
	BvbId string
	Name  string
}

type Match struct {
	PlayerA Player
	PlayerB Player
	PlayerC Player
	PlayerD Player
}

type Player struct {
	BvbId string
	Name  string
}

func GetMatches(reader io.Reader) []Match {
	playerRegex := `<a href="player.asp\?ID=(\d+)">(.+?)</a>`
	re := regexp.MustCompile(playerRegex + `.*?` + playerRegex + `.*?` + playerRegex + `.*?` + playerRegex)
	bytes, _ := ioutil.ReadAll(reader)
	regexMatches := re.FindAllStringSubmatch(string(bytes), -1)

	var matches []Match
	for _, value := range regexMatches {
		matches = append(matches, Match{
			Player{value[1], value[2]},
			Player{value[3], value[4]},
			Player{value[5], value[6]},
			Player{value[7], value[8]},
		})
	}

	return matches
}

func GetTournaments(reader io.Reader) []Tournament {
	re := regexp.MustCompile(`<a href="Tournament.asp\?ID=(\d+)">(.+?)</a>`)
	bytes, _ := ioutil.ReadAll(reader)
	regexMatches := re.FindAllStringSubmatch(string(bytes), -1)

	var tournaments []Tournament
	for _, value := range regexMatches {
		tournaments = append(tournaments, Tournament{value[1], value[2]})
	}

	return tournaments
}

func GetSeasons(reader io.Reader) []Season {
	re := regexp.MustCompile(`<a href="Season.asp\?AssocID=(\d+)&Year=(\d+)">`)
	bytes, _ := ioutil.ReadAll(reader)
	regexMatches := re.FindAllStringSubmatch(string(bytes), -1)

	var seasons []Season
	for _, value := range regexMatches {
		seasons = append(seasons, Season{value[1], value[2]})
	}

	return seasons
}
