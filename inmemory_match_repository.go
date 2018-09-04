package vbscraper

import (
	"log"
	"sync"
)

type InMemoryMatchRepository struct {
	matches sync.Map
}

var _ MatchRepository = (*InMemoryMatchRepository)(nil)

func (r *InMemoryMatchRepository) Create(match Match, id string) error {
	r.matches.Store(id, &match)
	return nil
}

func (r *InMemoryMatchRepository) Find(id string) *Match {
	match, ok := r.matches.Load(id)
	if !ok {
		log.Fatal("match not found")
	}

	return match.(*Match)
}

func (r *InMemoryMatchRepository) GetAllPlayerIds() []int {
	var playerIds []int

	r.matches.Range(func(k, v interface{}) bool {
		playerIds = append(
			playerIds,
			v.(*Match).PlayerAId,
			v.(*Match).PlayerBId,
			v.(*Match).PlayerCId,
			v.(*Match).PlayerDId,
		)
		return true
	})

	return playerIds
}

func (r *InMemoryMatchRepository) TotalMatches() interface{} {
	totalMatches := 0
	r.matches.Range(func(k, v interface{}) bool {
		totalMatches++
		return true
	})

	return totalMatches
}
