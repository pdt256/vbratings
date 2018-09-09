package vbratings

import (
	"github.com/pdt256/skill"
)

const defaultSeedRating = 1500

type ratingCalculator struct {
	matchRepository        MatchRepository
	playerRatingRepository PlayerRatingRepository
	playerRatings          map[int]PlayerRating
}

func NewRatingCalculator(matchRepository MatchRepository, playerRatingRepository PlayerRatingRepository) *ratingCalculator {
	return &ratingCalculator{
		matchRepository:        matchRepository,
		playerRatingRepository: playerRatingRepository,
		playerRatings:          make(map[int]PlayerRating),
	}
}

func (c *ratingCalculator) CalculateRatingsByYear(year int) int {

	ratingCalculator := skill.NewEloCalculator(32)
	duelingCalculator := skill.NewDuelingCalculator(ratingCalculator)

	for _, match := range c.matchRepository.GetAllMatchesByYear(year) {

		playerARating := c.getPlayerRating(match.PlayerAId, year)
		playerBRating := c.getPlayerRating(match.PlayerBId, year)
		playerCRating := c.getPlayerRating(match.PlayerCId, year)
		playerDRating := c.getPlayerRating(match.PlayerDId, year)

		teamARatings := []int{playerARating.Rating, playerBRating.Rating}
		teamBRatings := []int{playerCRating.Rating, playerDRating.Rating}

		newRatingsA, newRatingsB := duelingCalculator.GetNewRatings(teamARatings, teamBRatings, 1, 0)

		playerARating.Rating = newRatingsA[0]
		playerBRating.Rating = newRatingsA[1]
		playerCRating.Rating = newRatingsB[0]
		playerDRating.Rating = newRatingsB[1]

		c.playerRatings[match.PlayerAId] = playerARating
		c.playerRatings[match.PlayerBId] = playerBRating
		c.playerRatings[match.PlayerCId] = playerCRating
		c.playerRatings[match.PlayerDId] = playerDRating
	}

	for _, playerRating := range c.playerRatings {
		c.playerRatingRepository.Create(playerRating)
	}

	return len(c.playerRatings)
}

func (c *ratingCalculator) getPlayerRating(playerId int, year int) PlayerRating {
	if rating, ok := c.playerRatings[playerId]; ok {
		return rating
	}

	var playerRating *PlayerRating

	previousYear := year - 1
	playerRating, getPlayerRatingErr := c.playerRatingRepository.GetPlayerRatingByYear(playerId, previousYear)
	if getPlayerRatingErr == nil {
		playerRating.Year = year
		playerRating.SeedRating = playerRating.Rating
	} else if getPlayerRatingErr == PlayerRatingNotFoundError {
		playerRating = &PlayerRating{
			PlayerId:   playerId,
			Year:       year,
			SeedRating: defaultSeedRating,
			Rating:     defaultSeedRating,
		}
	}

	c.playerRatings[playerId] = *playerRating

	return *playerRating
}
