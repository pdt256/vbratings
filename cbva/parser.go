package cbva

import (
	"io"
	"io/ioutil"
	"regexp"
	"strconv"

	"github.com/pdt256/vbratings"
)

var tournamentResultRegexp = regexp.MustCompile(`(?sm:EarnedFinish":"([^"]+)".+?Player1_Name":"([^"]+)".+?Player2_Name":"([^"]+)")`)

func GetTournamentResults(reader io.Reader) []vbratings.TournamentResult {
	bytes, _ := ioutil.ReadAll(reader)
	body := string(bytes)

	tournamentResultMatches := tournamentResultRegexp.FindAllStringSubmatch(body, -1)

	var tournamentResults []vbratings.TournamentResult
	for _, value := range tournamentResultMatches {
		earnedFinish, _ := strconv.Atoi(value[1])
		player1Name := value[2]
		player2Name := value[3]

		tournamentResults = append(tournamentResults, vbratings.TournamentResult{
			Player1Name:  vbratings.GetPlayerNameSlug(player1Name),
			Player2Name:  vbratings.GetPlayerNameSlug(player2Name),
			EarnedFinish: earnedFinish,
		})
	}

	return tournamentResults
}
