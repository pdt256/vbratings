package vbscraper

import (
	"github.com/pdt256/skill"
)

type ratingCalculator struct {
	matchRepository        MatchRepository
	playerRatingRepository PlayerRatingRepository
}

func NewRatingCalculator(matchRepository MatchRepository, playerRatingRepository PlayerRatingRepository) *ratingCalculator {
	return &ratingCalculator{matchRepository, playerRatingRepository}
}

func (c *ratingCalculator) CalculateRatingsByYear(year int) {
	playerRatings := make(map[int]PlayerRating)
	_ = playerRatings

	const defaultSeedRating = 1500

	ratingCalculator := skill.NewEloCalculator(32)
	duelingCalculator := skill.NewDuelingCalculator(ratingCalculator)

	// get all matches for $year
	for _, match := range c.matchRepository.GetAllMatchesByYear(year) {
		// TODO:
		//  get playerRating from $playerRatings
		//  if not exist, find in playerRatingRepository where (year - 1)
		//  if not exist, new playerRating to $playerRatings with default SeedRating (1500)

		playerARating := PlayerRating{
			PlayerId:   match.PlayerAId,
			Year:       year,
			SeedRating: defaultSeedRating,
			Rating:     defaultSeedRating,
		}

		playerBRating := PlayerRating{
			PlayerId:   match.PlayerBId,
			Year:       year,
			SeedRating: defaultSeedRating,
			Rating:     defaultSeedRating,
		}

		playerCRating := PlayerRating{
			PlayerId:   match.PlayerCId,
			Year:       year,
			SeedRating: defaultSeedRating,
			Rating:     defaultSeedRating,
		}

		playerDRating := PlayerRating{
			PlayerId:   match.PlayerDId,
			Year:       year,
			SeedRating: defaultSeedRating,
			Rating:     defaultSeedRating,
		}

		teamARatings := []int{playerARating.Rating, playerBRating.Rating}
		teamBRatings := []int{playerCRating.Rating, playerDRating.Rating}

		newRatingsA, newRatingsB := duelingCalculator.GetNewRatings(teamARatings, teamBRatings, 1, 0)

		playerARating.Rating = newRatingsA[0]
		playerBRating.Rating = newRatingsA[1]

		playerCRating.Rating = newRatingsB[0]
		playerDRating.Rating = newRatingsB[1]

		playerRatings[match.PlayerAId] = playerARating
		playerRatings[match.PlayerBId] = playerBRating
		playerRatings[match.PlayerCId] = playerCRating
		playerRatings[match.PlayerDId] = playerDRating
	}

	for _, playerRating := range playerRatings {
		c.playerRatingRepository.Create(playerRating)
	}
}
