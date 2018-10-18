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

		c.updateCachedPlayerRating(playerARating)
		c.updateCachedPlayerRating(playerBRating)
		c.updateCachedPlayerRating(playerCRating)
		c.updateCachedPlayerRating(playerDRating)
	}

	totalSavedRatings := c.saveAllCachedRatings()

	return totalSavedRatings
}

func (c *ratingCalculator) CalculateRatingsByYearFromTournamentResults(year int) int {
	for _, tournamentAndResults := range c.tournamentRepository.GetAllTournamentsAndResultsByYear(year) {

		c.eachTeamBeatsAShadowTeam(tournamentAndResults, year)
		c.eachTeamBeatsBelowTeam(tournamentAndResults, year)

		// TODO: every team in a single earned finish beats the 2 below teams
	}

	totalSavedRatings := c.saveAllCachedRatings()

	return totalSavedRatings
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

	i := 0
	for _, result := range tournamentAndResults.Results {
		player1Rating := c.getPlayerRating(result.Player1Id, year)
		player2Rating := c.getPlayerRating(result.Player2Id, year)

		player1Rating.Rating = newRatings[i]
		player2Rating.Rating = newRatings[i+1]
		i += 2

		c.updateCachedPlayerRating(player1Rating)
		c.updateCachedPlayerRating(player2Rating)
	}
}

func (c *ratingCalculator) eachTeamBeatsBelowTeam(tournamentAndResults *TournamentAndResults, year int) {
	var teamARatings []int
	var teamBRatings []int
	earnedFinishA := 1
	earnedFinishB := 2
	for _, result := range tournamentAndResults.Results {
		player1Rating := c.getPlayerRating(result.Player1Id, year)
		player2Rating := c.getPlayerRating(result.Player2Id, year)

		if result.EarnedFinish == earnedFinishA {
			teamARatings = append(teamARatings, player1Rating.Rating, player2Rating.Rating)
		} else if result.EarnedFinish == earnedFinishB {
			teamBRatings = append(teamBRatings, player1Rating.Rating, player2Rating.Rating)
		}
	}

	if len(teamBRatings) == 0 {
		return
	}

	newRatingsA, newRatingsB := c.duelingCalculator.GetNewRatings(teamARatings, teamBRatings, 1, 0)

	aIndex := 0
	bIndex := 0
	for _, result := range tournamentAndResults.Results {
		player1Rating := c.getPlayerRating(result.Player1Id, year)
		player2Rating := c.getPlayerRating(result.Player2Id, year)

		if result.EarnedFinish == earnedFinishA {
			player1Rating.Rating = newRatingsA[aIndex]
			player2Rating.Rating = newRatingsA[aIndex+1]
			aIndex += 2
		} else if result.EarnedFinish == earnedFinishB {
			player1Rating.Rating = newRatingsB[bIndex]
			player2Rating.Rating = newRatingsB[bIndex+1]
			bIndex += 2
		}

		c.updateCachedPlayerRating(player1Rating)
		c.updateCachedPlayerRating(player2Rating)
	}
}

func (c *ratingCalculator) saveAllCachedRatings() int {
	for _, playerRating := range c.playerRatings {
		c.playerRatingRepository.Create(playerRating)
	}

	return len(c.playerRatings)
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

func (c *ratingCalculator) updateCachedPlayerRating(playerRating PlayerRating) {
	c.playerRatings[playerRating.PlayerId] = playerRating
}
