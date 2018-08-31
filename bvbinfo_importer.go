package vbscraper

import (
	"io"

	"github.com/satori/go.uuid"
)

type bvbinfoImporter struct {
	matchRepository MatchRepository
}

func (i *bvbinfoImporter) ImportMatches(reader io.Reader) int {
	matches := GetMatches(reader)

	totalImported := 0
	for _, match := range matches {
		id := uuid.NewV4().String()
		i.matchRepository.Create(match, id)
		totalImported++
	}

	return totalImported
}

func NewBvbInfoImporter(matchRepository MatchRepository) *bvbinfoImporter {
	return &bvbinfoImporter{matchRepository}
}
