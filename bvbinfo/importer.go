package bvbinfo

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/pdt256/vbratings"
	"github.com/satori/go.uuid"
)

type Importer struct {
	matchRepository  vbratings.MatchRepository
	playerRepository vbratings.PlayerRepository
}

func NewImporter(matchRepository vbratings.MatchRepository, playerRepository vbratings.PlayerRepository) *Importer {
	return &Importer{matchRepository, playerRepository}
}

func (i *Importer) ImportAllSeasons() int {
	allSeasonsUrl := "http://bvbinfo.com/season.asp"
	allSeasonsResponse, err := http.Get(allSeasonsUrl)
	checkError(err)

	seasons := GetSeasons(allSeasonsResponse.Body)
	allSeasonsResponse.Body.Close()

	totalMatchesImported := 0
	for _, season := range seasons {
		seasonUrl := fmt.Sprintf("http://bvbinfo.com/Season.asp?AssocID=%s&Year=%s", season.AssocID, season.Year)
		totalMatchesImported += i.ImportSeason(seasonUrl)
	}
	return totalMatchesImported
}

func (i *Importer) ImportAllPlayers() int {
	totalImported := 0
	for _, playerId := range i.matchRepository.GetAllPlayerIds() {
		playerUrl := fmt.Sprintf("http://bvbinfo.com/player.asp?ID=%d", playerId)

		playerResponse, err := http.Get(playerUrl)
		if err != nil {
			log.Fatal(err)
		}

		player := GetPlayer(playerResponse.Body, playerId)
		playerResponse.Body.Close()

		i.playerRepository.Create(player)
		totalImported++
		fmt.Print(".")
	}
	return totalImported
}

func (i *Importer) ImportSeason(seasonUrl string) int {
	seasonResponse, err := http.Get(seasonUrl)
	checkError(err)

	tournaments := GetTournaments(seasonResponse.Body)
	seasonResponse.Body.Close()

	totalMatchesImported := 0
	for _, tournament := range tournaments {
		tournamentUrl := fmt.Sprintf("http://bvbinfo.com/Tournament.asp?ID=%s&Process=Matches", tournament.BvbId)
		totalMatchesImported += i.ImportTournament(tournamentUrl)
	}
	return totalMatchesImported
}

func (i *Importer) ImportTournament(tournamentUrl string) int {
	tournamentResponse, err := http.Get(tournamentUrl)
	checkError(err)

	defer tournamentResponse.Body.Close()

	totalMatchesImported := i.ImportMatches(tournamentResponse.Body)
	fmt.Print(".")

	return totalMatchesImported
}

func (i *Importer) ImportMatches(reader io.Reader) int {
	matches := GetMatches(reader)

	totalImported := 0
	for _, match := range matches {
		id := uuid.NewV4().String()
		i.matchRepository.Create(match, id)
		totalImported++
	}

	return totalImported
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
