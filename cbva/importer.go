package cbva

import (
	"io"
	"log"

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

func (i *Importer) ImportTournamentResults(reader io.Reader) (int, int) {
	cbvaTournamentResults := GetTournamentResults(reader)

	totalResults := 0
	totalPlayers := 0
	for _, cbvaTournamentResult := range cbvaTournamentResults {
		id := uuid.NewService().NewV4()

		player1Id, player1Created := i.getPlayerIdFromCBVAPlayer(cbvaTournamentResult.Player1)
		player2Id, player2Created := i.getPlayerIdFromCBVAPlayer(cbvaTournamentResult.Player2)

		totalPlayers += player1Created + player2Created

		tournamentResult := vbratings.TournamentResult{
			Player1Id:    player1Id,
			Player2Id:    player2Id,
			EarnedFinish: cbvaTournamentResult.EarnedFinish,
		}

		i.tournamentRepository.AddTournamentResult(tournamentResult, id)
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
