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
	matchRepository   vbratings.MatchRepository
	playerRepository  vbratings.PlayerRepository
	bvbInfoRepository Repository
	uuidGenerator     uuid.Generator
}

func NewImporter(
	matchRepository vbratings.MatchRepository,
	playerRepository vbratings.PlayerRepository,
	bvbInfoRepository Repository,
	uuidGenerator uuid.Generator,
) *Importer {
	return &Importer{
		matchRepository:   matchRepository,
		playerRepository:  playerRepository,
		bvbInfoRepository: bvbInfoRepository,
		uuidGenerator:     uuidGenerator,
	}
}

func (i *Importer) ImportAllSeasons() (int, int) {
	allSeasonsUrl := "http://bvbinfo.com/season.asp"
	allSeasonsResponse, err := http.Get(allSeasonsUrl)
	checkError(err)

	seasons := GetSeasons(allSeasonsResponse.Body)
	allSeasonsResponse.Body.Close()

	totalMatches := 0
	totalPlayers := 0
	for _, season := range seasons {
		seasonUrl := fmt.Sprintf("http://bvbinfo.com/Season.asp?AssocID=%s&Year=%s", season.AssocID, season.Year)
		nMatches, nPlayers := i.ImportSeason(seasonUrl)

		totalMatches += nMatches
		totalPlayers += nPlayers
	}
	return totalMatches, totalPlayers
}

func (i *Importer) ImportSeason(seasonUrl string) (int, int) {
	seasonResponse, err := http.Get(seasonUrl)
	checkError(err)

	tournaments := GetTournaments(seasonResponse.Body)
	seasonResponse.Body.Close()

	totalMatches := 0
	totalPlayers := 0
	for _, tournament := range tournaments {
		tournamentUrl := fmt.Sprintf("http://bvbinfo.com/Tournament.asp?ID=%d&Process=Matches", tournament.Id)
		nMatches, nPlayers := i.ImportTournament(tournamentUrl)

		totalMatches += nMatches
		totalPlayers += nPlayers
	}
	return totalMatches, totalPlayers
}

func (i *Importer) ImportTournament(tournamentUrl string) (int, int) {
	tournamentResponse, err := http.Get(tournamentUrl)
	checkError(err)

	defer tournamentResponse.Body.Close()

	fmt.Print(".")
	return i.ImportMatches(tournamentResponse.Body)
}

func (i *Importer) ImportMatches(reader io.Reader) (int, int) {
	matches := GetMatches(reader)

	totalMatches := 0
	totalPlayers := 0
	for _, bvbInfoMatch := range matches {
		id := i.uuidGenerator.NewV4()

		playerAId, playerACreated := i.getPlayerIdFromBvbInfoPlayer(bvbInfoMatch.PlayerA)
		playerBId, playerBCreated := i.getPlayerIdFromBvbInfoPlayer(bvbInfoMatch.PlayerB)
		playerCId, playerCCreated := i.getPlayerIdFromBvbInfoPlayer(bvbInfoMatch.PlayerC)
		playerDId, playerDCreated := i.getPlayerIdFromBvbInfoPlayer(bvbInfoMatch.PlayerD)

		totalPlayers += playerACreated + playerBCreated + playerCCreated + playerDCreated

		match := vbratings.Match{
			PlayerAId: playerAId,
			PlayerBId: playerBId,
			PlayerCId: playerCId,
			PlayerDId: playerDId,
			IsForfeit: bvbInfoMatch.IsForfeit,
			Year:      bvbInfoMatch.Year,
			Set1:      bvbInfoMatch.Set1,
			Set2:      bvbInfoMatch.Set2,
			Set3:      bvbInfoMatch.Set3,
			Gender:    vbratings.GenderFromString(bvbInfoMatch.Gender),
		}

		i.matchRepository.Create(match, id)
		totalMatches++
	}

	return totalMatches, totalPlayers
}

func (i *Importer) getPlayerIdFromBvbInfoPlayer(bvbInfoPlayer Player) (string, int) {
	var playerId string
	playersCreated := 0

	playerId, err := i.bvbInfoRepository.GetPlayerId(bvbInfoPlayer.Id)
	if err == PlayerNotFoundError {
		playerId = i.uuidGenerator.NewV4()

		newPlayer := vbratings.Player{
			Id:     playerId,
			Name:   bvbInfoPlayer.Name,
			ImgUrl: fmt.Sprintf("http://bvbinfo.com/images/photos/%d.jpg", bvbInfoPlayer.Id),
		}

		createErr := i.playerRepository.Create(newPlayer)
		if createErr != nil {
			log.Printf("Unable to add player: %+v", newPlayer)
		}

		addPlayerIdErr := i.bvbInfoRepository.AddPlayerId(playerId, bvbInfoPlayer.Id)
		if addPlayerIdErr != nil {
			log.Printf("Unable to add playerId -> BvbId mapping: %v", err)
		}

		playersCreated++
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
