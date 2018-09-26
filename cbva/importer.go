package cbva

import (
	"io"

	"github.com/pdt256/vbratings"
	"github.com/satori/go.uuid"
)

type Importer struct {
	tournamentRepository vbratings.TournamentRepository
	playerRepository     vbratings.PlayerRepository
}

func NewImporter(tournamentRepository vbratings.TournamentRepository, playerRepository vbratings.PlayerRepository) *Importer {
	return &Importer{tournamentRepository, playerRepository}
}

func (i *Importer) ImportTournamentResults(reader io.Reader) int {
	tournamentResults := GetTournamentResults(reader)

	totalImported := 0
	for _, tournamentResult := range tournamentResults {
		id := uuid.NewV4().String()
		i.tournamentRepository.AddTournamentResult(tournamentResult, id)
		totalImported++
	}

	return totalImported
}
