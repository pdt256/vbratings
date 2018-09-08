package bvbinfo

import (
	"io"

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
