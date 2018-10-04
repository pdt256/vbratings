package cbva

import (
	"io"
	"io/ioutil"
	"regexp"
	"strconv"
)

var tournamentResultRegexp = regexp.MustCompile(`(?sm:EarnedFinish":"([^"]+)".+?Player1_Name":"([^"]+)".+?Player2_Name":"([^"]+)")`)

var tournamentsRegexp = regexp.MustCompile(`(?sm:data-id=\\"([^\\]+)\\" data-Date=\\"([^\\]+)\\" data-Rating=\\"([^\\]+)\\" data-Gender=\\"([^\\]+)\\".+?data-Location=\\"([^\\]+)\\")`)

type Tournament struct {
	Id           string
	Date         string
	Rating       string
	Gender       string
	Location     string
	TournamentId string
}

func (t Tournament) Year() int {
	year, _ := strconv.Atoi(t.Date[len(t.Date)-4:])
	return year
}

type TournamentResult struct {
	Player1      Player
	Player2      Player
	EarnedFinish int
}

type Player struct {
	Name string
}

func GetTournamentResults(reader io.Reader) []TournamentResult {
	bytes, _ := ioutil.ReadAll(reader)
	body := string(bytes)

	tournamentResultMatches := tournamentResultRegexp.FindAllStringSubmatch(body, -1)

	var tournamentResults []TournamentResult
	for _, value := range tournamentResultMatches {
		earnedFinish, _ := strconv.Atoi(value[1])
		tournamentResults = append(tournamentResults, TournamentResult{
			Player1: Player{
				Name: value[2],
			},
			Player2: Player{
				Name: value[3],
			},
			EarnedFinish: earnedFinish,
		})
	}

	return tournamentResults
}

func GetTournaments(reader io.Reader) []Tournament {
	bytes, _ := ioutil.ReadAll(reader)
	body := string(bytes)

	tournamentMatches := tournamentsRegexp.FindAllStringSubmatch(body, -1)

	var tournaments []Tournament
	for _, value := range tournamentMatches {
		tournaments = append(tournaments, Tournament{
			Id:       value[1],
			Date:     value[2],
			Rating:   value[3],
			Gender:   value[4],
			Location: value[5],
		})
	}

	return tournaments
}
