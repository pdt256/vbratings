package cbva

import (
	"io"
	"io/ioutil"
	"regexp"
	"strconv"
)

var tournamentResultRegexp = regexp.MustCompile(`(?sm:EarnedFinish":"([^"]+)".+?Player1_Name":"([^"]+)".+?Player2_Name":"([^"]+)")`)

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
