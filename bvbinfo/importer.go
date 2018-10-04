package bvbinfo

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/pdt256/vbratings"
	"github.com/pdt256/vbratings/pkg/uuid"
)

type Importer struct {
	tournamentRepository vbratings.TournamentRepository
	matchRepository      vbratings.MatchRepository
	playerRepository     vbratings.PlayerRepository
	bvbInfoRepository    Repository
	uuidGenerator        uuid.Generator
}

func NewImporter(
	tournamentRepository vbratings.TournamentRepository,
	matchRepository vbratings.MatchRepository,
	playerRepository vbratings.PlayerRepository,
	bvbInfoRepository Repository,
	uuidGenerator uuid.Generator,
) *Importer {
	return &Importer{
		tournamentRepository: tournamentRepository,
		matchRepository:      matchRepository,
		playerRepository:     playerRepository,
		bvbInfoRepository:    bvbInfoRepository,
		uuidGenerator:        uuidGenerator,
	}
}

func (i *Importer) ImportAllSeasons() (int, int, int) {
	allSeasonsUrl := "http://bvbinfo.com/season.asp"
	allSeasonsResponse, err := http.Get(allSeasonsUrl)
	checkError(err)

	seasons := GetSeasons(allSeasonsResponse.Body)
	allSeasonsResponse.Body.Close()

	totalTournaments := 0
	totalMatches := 0
	totalPlayers := 0
	for _, season := range seasons {
		seasonUrl := fmt.Sprintf("http://bvbinfo.com/Season.asp?AssocID=%s&Year=%s", season.AssocID, season.Year)
		nTournaments, nMatches, nPlayers := i.ImportSeason(seasonUrl)

		totalTournaments += nTournaments
		totalMatches += nMatches
		totalPlayers += nPlayers
	}

	return totalTournaments, totalMatches, totalPlayers
}

func (i *Importer) ImportSeason(seasonUrl string) (int, int, int) {
	seasonResponse, err := http.Get(seasonUrl)
	checkError(err)

	tournaments := GetTournaments(seasonResponse.Body)
	seasonResponse.Body.Close()

	totalTournaments := 0
	totalMatches := 0
	totalPlayers := 0
	for _, tournament := range tournaments {
		tournamentUrl := fmt.Sprintf("http://bvbinfo.com/Tournament.asp?ID=%d&Process=Matches", tournament.Id)
		nMatches, nPlayers := i.ImportTournament(tournamentUrl, tournament.Id)

		totalTournaments++
		totalMatches += nMatches
		totalPlayers += nPlayers
	}

	return totalTournaments, totalMatches, totalPlayers
}

func (i *Importer) ImportTournament(tournamentUrl string, tournamentId int) (int, int) {
	tournamentResponse, err := http.Get(tournamentUrl)
	checkError(err)

	defer tournamentResponse.Body.Close()

	fmt.Print(".")
	return i.ImportMatches(tournamentResponse.Body, tournamentId)
}

func (i *Importer) ImportMatches(reader io.Reader, tournamentId int) (int, int) {
	tournament, matches := GetMatches(reader, tournamentId)

	newTournament := vbratings.Tournament{
		Id:     i.uuidGenerator.NewV4(),
		Date:   tournament.Dates,
		Gender: tournament.Gender,
		Year:   tournament.Year,
		Name:   tournament.Name,
	}

	i.tournamentRepository.Create(newTournament)

	tournament.TournamentId = newTournament.Id
	i.bvbInfoRepository.AddTournament(tournament)

	totalMatches := 0
	totalPlayers := 0
	for _, bvbInfoMatch := range matches {
		playerAId, playerACreated := i.getPlayerIdFromBvbInfoPlayer(bvbInfoMatch.PlayerA, tournament.Gender)
		playerBId, playerBCreated := i.getPlayerIdFromBvbInfoPlayer(bvbInfoMatch.PlayerB, tournament.Gender)
		playerCId, playerCCreated := i.getPlayerIdFromBvbInfoPlayer(bvbInfoMatch.PlayerC, tournament.Gender)
		playerDId, playerDCreated := i.getPlayerIdFromBvbInfoPlayer(bvbInfoMatch.PlayerD, tournament.Gender)

		totalPlayers += playerACreated + playerBCreated + playerCCreated + playerDCreated

		newMatch := vbratings.Match{
			Id:           i.uuidGenerator.NewV4(),
			PlayerAId:    playerAId,
			PlayerBId:    playerBId,
			PlayerCId:    playerCId,
			PlayerDId:    playerDId,
			IsForfeit:    bvbInfoMatch.IsForfeit,
			Set1:         bvbInfoMatch.Set1,
			Set2:         bvbInfoMatch.Set2,
			Set3:         bvbInfoMatch.Set3,
			TournamentId: newTournament.Id,
		}

		i.matchRepository.Create(newMatch)
		totalMatches++
	}

	return totalMatches, totalPlayers
}

func (i *Importer) getPlayerIdFromBvbInfoPlayer(bvbInfoPlayer Player, gender string) (string, int) {
	var playerId string
	playersCreated := 0

	playerId, err := i.bvbInfoRepository.GetPlayerId(bvbInfoPlayer.Id)
	if err == PlayerNotFoundError {
		newPlayer := vbratings.Player{
			Id:     i.uuidGenerator.NewV4(),
			Name:   bvbInfoPlayer.Name,
			Gender: gender,
		}

		createErr := i.playerRepository.Create(newPlayer)
		if createErr != nil {
			log.Printf("Unable to add player: %+v", newPlayer)
		}

		playerId = newPlayer.Id
		bvbInfoPlayer.PlayerId = newPlayer.Id
		addPlayerIdErr := i.bvbInfoRepository.AddPlayer(bvbInfoPlayer)
		if addPlayerIdErr != nil {
			log.Printf("Unable to add playerId -> BvbId mapping: %v", err)
		}

		playersCreated = 1
	} else if err != nil {
		log.Printf("Unknown error: %v", err)
	}

	return playerId, playersCreated
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
