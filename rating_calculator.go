package vbratings

import (
	"github.com/pdt256/skill"
)

const defaultSeedRating = 1500

type ratingCalculator struct {
	matchRepository        MatchRepository
	tournamentRepository   TournamentRepository
	playerRatingRepository PlayerRatingRepository
	playerRatings          map[string]PlayerRating
	duelingCalculator      skill.DuelingCalculator
}

func NewRatingCalculator(
	matchRepository MatchRepository,
	tournamentRepository TournamentRepository,
	playerRatingRepository PlayerRatingRepository) *ratingCalculator {

	duelingCalculator := skill.NewDuelingCalculator(
		skill.NewEloCalculator(32),
	)

	return &ratingCalculator{
		matchRepository:        matchRepository,
		tournamentRepository:   tournamentRepository,
		playerRatingRepository: playerRatingRepository,
		playerRatings:          make(map[string]PlayerRating),
		duelingCalculator:      duelingCalculator,
	}
}

func (c *ratingCalculator) CalculateRatingsByYearFromMatches(year int) int {

	for _, match := range c.matchRepository.GetAllMatchesByYear(year) {
		playerARating := c.getPlayerRating(match.PlayerAId, year)
		playerBRating := c.getPlayerRating(match.PlayerBId, year)
		playerCRating := c.getPlayerRating(match.PlayerCId, year)
		playerDRating := c.getPlayerRating(match.PlayerDId, year)

		teamARatings := []int{playerARating.Rating, playerBRating.Rating}
		teamBRatings := []int{playerCRating.Rating, playerDRating.Rating}

		newRatingsA, newRatingsB := c.duelingCalculator.GetNewRatings(teamARatings, teamBRatings, 1, 0)

		playerARating.Rating = newRatingsA[0]
		playerBRating.Rating = newRatingsA[1]
		playerCRating.Rating = newRatingsB[0]
		playerDRating.Rating = newRatingsB[1]

		playerARating.TotalMatches++
		playerBRating.TotalMatches++
		playerCRating.TotalMatches++
		playerDRating.TotalMatches++

		c.playerRatings[match.PlayerAId] = playerARating
		c.playerRatings[match.PlayerBId] = playerBRating
		c.playerRatings[match.PlayerCId] = playerCRating
		c.playerRatings[match.PlayerDId] = playerDRating
	}

	c.saveAllCachedRatings()

	return len(c.playerRatings)
}

func (c *ratingCalculator) CalculateRatingsByYearFromTournamentResults(year int) int {
	for _, tournamentAndResults := range c.tournamentRepository.GetAllTournamentsAndResultsByYear(year) {

		c.eachTeamBeatsAShadowTeam(tournamentAndResults, year)

		// TODO: every team in a single earned finish beats the 2 below teams

		c.saveAllCachedRatings()
	}

	totalRatingsCalculated := 4

	return totalRatingsCalculated
}

func (c *ratingCalculator) saveAllCachedRatings() {
	for _, playerRating := range c.playerRatings {
		c.playerRatingRepository.Create(playerRating)
	}
}

func (c *ratingCalculator) eachTeamBeatsAShadowTeam(tournamentAndResults *TournamentAndResults, year int) {
	var teamARatings []int
	for _, result := range tournamentAndResults.Results {
		player1Rating := c.getPlayerRating(result.Player1Id, year)
		player2Rating := c.getPlayerRating(result.Player2Id, year)

		teamARatings = append(teamARatings, player1Rating.Rating, player2Rating.Rating)
	}
	newRatings := teamARatings
	shadowRatings := []int{1500}
	newRatings, _ = c.duelingCalculator.GetNewRatings(newRatings, shadowRatings, 1, 0)
	newRatings, _ = c.duelingCalculator.GetNewRatings(newRatings, shadowRatings, 1, 0)
	newRatings, _ = c.duelingCalculator.GetNewRatings(newRatings, shadowRatings, 1, 0)
	for i, result := range tournamentAndResults.Results {
		player1Rating := c.getPlayerRating(result.Player1Id, year)
		player2Rating := c.getPlayerRating(result.Player2Id, year)

		player1Rating.Rating = newRatings[i]
		player2Rating.Rating = newRatings[i+1]

		c.playerRatings[result.Player1Id] = player1Rating
		c.playerRatings[result.Player2Id] = player2Rating
	}
}

func (c *ratingCalculator) getPlayerRating(playerId string, year int) PlayerRating {
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
