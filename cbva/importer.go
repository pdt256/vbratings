package cbva

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/pdt256/vbratings"
	"github.com/pdt256/vbratings/pkg/uuid"
)

type Importer struct {
	tournamentRepository vbratings.TournamentRepository
	playerRepository     vbratings.PlayerRepository
	cbvaRepository       Repository
	uuidGenerator        uuid.Generator
}

func NewImporter(
	tournamentRepository vbratings.TournamentRepository,
	playerRepository vbratings.PlayerRepository,
	cbvaRepository Repository,
	uuidGenerator uuid.Generator,
) *Importer {
	return &Importer{
		tournamentRepository: tournamentRepository,
		playerRepository:     playerRepository,
		cbvaRepository:       cbvaRepository,
		uuidGenerator:        uuidGenerator,
	}
}

func (i *Importer) ImportAllTournaments() (int, int, int) {
	totalTournaments := 0
	totalResults := 0
	totalPlayers := 0

	page := 1

	for {
		postData := fmt.Sprintf(`{"AgeType":"","Gender":"","Divisions":null,"Location":"","SortBy":"","IsDesc":"true","StartDate":"1990-01-01T08:00:00.000Z","PageNumber":%d,"PageSize":2000}`, page)
		req, _ := http.NewRequest("POST", "https://cbva.com/Results/SearchTournamentResult", bytes.NewReader([]byte(postData)))
		req.Header.Add("Content-Type", "application/json")
		client := &http.Client{}
		tournamentsResponse, err := client.Do(req)
		checkError(err)

		tournaments := GetTournaments(tournamentsResponse.Body)

		for _, cbvaTournament := range tournaments {
			nMatches, nPlayers := i.ImportTournament(cbvaTournament)

			totalTournaments++
			totalResults += nMatches
			totalPlayers += nPlayers
		}

		tournamentsResponse.Body.Close()

		if len(tournaments) == 0 {
			break
		}

		page++
	}

	return totalTournaments, totalResults, totalPlayers
}

func (i *Importer) ImportTournament(cbvaTournament Tournament) (int, int) {
	postData := fmt.Sprintf(`{"id":"%s"}`, cbvaTournament.Id)
	req, _ := http.NewRequest("POST", "https://cbva.com/Results/GetTournamentTeamResult", bytes.NewReader([]byte(postData)))
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	tournamentResponse, err := client.Do(req)
	checkError(err)
	defer tournamentResponse.Body.Close()

	tournament := vbratings.Tournament{
		Id:       i.uuidGenerator.NewV4(),
		Date:     cbvaTournament.Date,
		Rating:   cbvaTournament.Rating,
		Gender:   cbvaTournament.Gender,
		Location: cbvaTournament.Location,
	}
	i.tournamentRepository.AddTournament(tournament)

	i.cbvaRepository.AddTournamentId(tournament.Id, cbvaTournament.Id)

	fmt.Print(".")
	return i.ImportTournamentResults(tournamentResponse.Body, tournament.Id)
}

func (i *Importer) ImportTournamentResults(reader io.Reader, tournamentId string) (int, int) {
	cbvaTournamentResults := GetTournamentResults(reader)

	totalResults := 0
	totalPlayers := 0
	for _, cbvaTournamentResult := range cbvaTournamentResults {
		player1Id, player1Created := i.getPlayerIdFromCBVAPlayer(cbvaTournamentResult.Player1)
		player2Id, player2Created := i.getPlayerIdFromCBVAPlayer(cbvaTournamentResult.Player2)

		totalPlayers += player1Created + player2Created

		tournamentResult := vbratings.TournamentResult{
			Id:           i.uuidGenerator.NewV4(),
			Player1Id:    player1Id,
			Player2Id:    player2Id,
			EarnedFinish: cbvaTournamentResult.EarnedFinish,
			TournamentId: tournamentId,
		}

		i.tournamentRepository.AddTournamentResult(tournamentResult)

		totalResults++
	}

	return totalResults, totalPlayers
}

func (i *Importer) getPlayerIdFromCBVAPlayer(player Player) (string, int) {
	var playerId string
	playersCreated := 0

	playerId, err := i.cbvaRepository.GetPlayerId(player.Name)
	if err == PlayerNotFoundError {
		playerId = i.uuidGenerator.NewV4()

		newPlayer := vbratings.Player{
			Id:   playerId,
			Name: player.Name,
		}

		createErr := i.playerRepository.Create(newPlayer)
		if createErr != nil {
			log.Printf("Unable to add player: %+v", newPlayer)
		}

		addPlayerIdErr := i.cbvaRepository.AddPlayerId(playerId, player.Name)
		if addPlayerIdErr != nil {
			log.Printf("Unable to add playerId -> BvbId mapping: %v", err)
		}

		playersCreated++

	} else if err != nil {
		log.Printf("Unknown error: %v", err)
	}

	return playerId, playersCreated
}
