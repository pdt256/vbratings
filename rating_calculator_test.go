package vbratings_test

import (
	"fmt"
	"testing"

	"github.com/pdt256/skill"
	"github.com/pdt256/vbratings"
	"github.com/pdt256/vbratings/sqlite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	playerAId = "98556e2665224abc99c2d07d621befa7"
	playerBId = "7a30fab9631a442d83b70c9bf1293be8"
	playerCId = "91f67d94a9a54c91b9f0ee0efc497c28"
	playerDId = "ee655d72d148459ca10d05cce939bcab"
	playerEId = "0dbfed48bee240c2936a70a86663ed62"
	playerFId = "305f1a6f4ae64884956d0ed7ce13785b"
	playerGId = "8b3ac1934c6f4593bca0bb888196fa58"
	playerHId = "98c7546096ea4773b4c1bc5195b80cb7"
	playerIId = "6395114fbb8b4e589cabe50addfcc7fd"
	playerJId = "4fc8124ccd314172b0a4085346ba303c"
	playerKId = "44d08b8c252f47ed83a8677f6982af26"
	playerLId = "f6ab7b01caa94f18a87a69f3930be9bc"
	playerMId = "c4d08b8c252f47ed83a8677f6982af26"
	playerNId = "c6ab7b01caa94f18a87a69f3930be9bc"
	playerOId = "d4d08b8c252f47ed83a8677f6982af26"
	playerPId = "d6ab7b01caa94f18a87a69f3930be9bc"
)

func Test_RatingCalculator_CalculateRatingsByYear_SingleMatch(t *testing.T) {
	// Given
	tournament := vbratings.Tournament{
		Id:     "1b97c1b593f84566bb932678eaf8a30d",
		Gender: "male",
		Year:   2018,
	}
	match := vbratings.Match{
		Id:           "0b9fb995e28a451fa5aaca68d397c1e0",
		PlayerAId:    playerAId,
		PlayerBId:    playerBId,
		PlayerCId:    playerCId,
		PlayerDId:    playerDId,
		IsForfeit:    false,
		Set1:         "17-21",
		Set2:         "21-15",
		Set3:         "15-7",
		TournamentId: tournament.Id,
	}
	db := sqlite.NewInMemoryDB()
	tournamentRepository := sqlite.NewTournamentRepository(db)
	tournamentRepository.Create(tournament)
	matchRepository := sqlite.NewMatchRepository(db)
	matchRepository.Create(match)
	playerRatingRepository := sqlite.NewPlayerRatingRepository(db)
	ratingCalculator := vbratings.NewRatingCalculator(matchRepository, tournamentRepository, playerRatingRepository)

	// When
	totalCalculated := ratingCalculator.CalculateRatingsByYearFromMatches(2018)

	// Then
	assert.Equal(t, 4, totalCalculated)
	playerRatingA, _ := playerRatingRepository.GetPlayerRatingByYear(playerAId, 2018)
	playerRatingB, _ := playerRatingRepository.GetPlayerRatingByYear(playerBId, 2018)
	playerRatingC, _ := playerRatingRepository.GetPlayerRatingByYear(playerCId, 2018)
	playerRatingD, _ := playerRatingRepository.GetPlayerRatingByYear(playerDId, 2018)
	assertPlayerRating(t, playerRatingA, 1500, 1516, 2018)
	assertPlayerRating(t, playerRatingB, 1500, 1516, 2018)
	assertPlayerRating(t, playerRatingC, 1500, 1484, 2018)
	assertPlayerRating(t, playerRatingD, 1500, 1484, 2018)
}

func Test_RatingCalculator_CalculateRatingsByYear_SeededWithPlayerRatingFromPreviousMatch(t *testing.T) {
	// Given
	tournament := vbratings.Tournament{
		Id:     "383cea98af1f40ebbf6fdfd6deec7270",
		Gender: "female",
		Year:   2018,
	}
	match1 := vbratings.Match{
		Id:           "bb5ff2b66aba45998fea0d8c0dc8bf52",
		PlayerAId:    playerAId,
		PlayerBId:    playerBId,
		PlayerCId:    playerCId,
		PlayerDId:    playerDId,
		IsForfeit:    false,
		Set1:         "17-21",
		Set2:         "21-15",
		Set3:         "15-7",
		TournamentId: tournament.Id,
	}
	match2 := vbratings.Match{
		Id:           "f87317add92e452f877e6b43e31c0a16",
		PlayerAId:    playerAId,
		PlayerBId:    playerBId,
		PlayerCId:    playerCId,
		PlayerDId:    playerDId,
		IsForfeit:    false,
		Set1:         "17-21",
		Set2:         "21-15",
		Set3:         "15-7",
		TournamentId: tournament.Id,
	}
	db := sqlite.NewInMemoryDB()
	tournamentRepository := sqlite.NewTournamentRepository(db)
	tournamentRepository.Create(tournament)
	matchRepository := sqlite.NewMatchRepository(db)
	matchRepository.Create(match1)
	matchRepository.Create(match2)
	playerRatingRepository := sqlite.NewPlayerRatingRepository(db)
	ratingCalculator := vbratings.NewRatingCalculator(matchRepository, tournamentRepository, playerRatingRepository)

	// When
	totalCalculated := ratingCalculator.CalculateRatingsByYearFromMatches(2018)

	// Then
	assert.Equal(t, 4, totalCalculated)
	playerRatingA, _ := playerRatingRepository.GetPlayerRatingByYear(playerAId, 2018)
	playerRatingB, _ := playerRatingRepository.GetPlayerRatingByYear(playerBId, 2018)
	playerRatingC, _ := playerRatingRepository.GetPlayerRatingByYear(playerCId, 2018)
	playerRatingD, _ := playerRatingRepository.GetPlayerRatingByYear(playerDId, 2018)
	assertPlayerRating(t, playerRatingA, 1500, 1530, 2018)
	assertPlayerRating(t, playerRatingB, 1500, 1530, 2018)
	assertPlayerRating(t, playerRatingC, 1500, 1469, 2018)
	assertPlayerRating(t, playerRatingD, 1500, 1469, 2018)
	assert.Equal(t, 2, playerRatingA.TotalMatches)
}

func Test_RatingCalculator_CalculateRatingsByYear_SeededWithPreviousYearPlayerRating(t *testing.T) {
	// Given
	playerRating := vbratings.PlayerRating{
		PlayerId:   playerAId,
		Year:       2017,
		SeedRating: 1500,
		Rating:     1600,
	}
	tournament := vbratings.Tournament{
		Id:     "3f955b24ef804c198b9ddcfaf330ff8f",
		Gender: "male",
		Year:   2018,
	}
	match := vbratings.Match{
		Id:           "34d659c28edb4a8b94ea5c39dea32534",
		PlayerAId:    playerAId,
		PlayerBId:    playerBId,
		PlayerCId:    playerCId,
		PlayerDId:    playerDId,
		IsForfeit:    false,
		Set1:         "17-21",
		Set2:         "21-15",
		Set3:         "15-7",
		TournamentId: tournament.Id,
	}
	db := sqlite.NewInMemoryDB()
	tournamentRepository := sqlite.NewTournamentRepository(db)
	tournamentRepository.Create(tournament)
	matchRepository := sqlite.NewMatchRepository(db)
	matchRepository.Create(match)
	playerRatingRepository := sqlite.NewPlayerRatingRepository(db)
	playerRatingRepository.Create(playerRating)
	ratingCalculator := vbratings.NewRatingCalculator(matchRepository, tournamentRepository, playerRatingRepository)

	// When
	totalCalculated := ratingCalculator.CalculateRatingsByYearFromMatches(2018)

	// Then
	assert.Equal(t, 4, totalCalculated)
	playerRatingA, _ := playerRatingRepository.GetPlayerRatingByYear(playerAId, 2018)
	playerRatingB, _ := playerRatingRepository.GetPlayerRatingByYear(playerBId, 2018)
	playerRatingC, _ := playerRatingRepository.GetPlayerRatingByYear(playerCId, 2018)
	playerRatingD, _ := playerRatingRepository.GetPlayerRatingByYear(playerDId, 2018)
	assertPlayerRating(t, playerRatingA, 1600, 1611, 2018)
	assertPlayerRating(t, playerRatingB, 1500, 1516, 2018)
	assertPlayerRating(t, playerRatingC, 1500, 1486, 2018)
	assertPlayerRating(t, playerRatingD, 1500, 1486, 2018)
}

func Test_RatingCalculator_CalculateRatingsByYearFromTournamentResults_1TeamWith1EarnedFinish(t *testing.T) {
	// Given
	const year = 2018
	tournament := vbratings.Tournament{
		Id:   "3366b167a109496db63f43169e4ac1a7",
		Year: year,
	}

	result1 := vbratings.TournamentResult{
		Id:           "e57fd803a81349b69f0e7160fa13e919",
		Player1Id:    playerAId,
		Player2Id:    playerBId,
		EarnedFinish: 1,
		TournamentId: tournament.Id,
	}

	db := sqlite.NewInMemoryDB()
	tournamentRepository := sqlite.NewTournamentRepository(db)
	tournamentRepository.Create(tournament)
	tournamentRepository.AddTournamentResult(result1)
	matchRepository := sqlite.NewMatchRepository(db)
	playerRatingRepository := sqlite.NewPlayerRatingRepository(db)
	ratingCalculator := vbratings.NewRatingCalculator(
		matchRepository,
		tournamentRepository,
		playerRatingRepository,
	)

	// When
	totalCalculated := ratingCalculator.CalculateRatingsByYearFromTournamentResults(year)

	// Then
	assert.Equal(t, 2, totalCalculated)
	assertNewRating(t, playerRatingRepository, result1.Player1Id, 1545, year)
	assertNewRating(t, playerRatingRepository, result1.Player2Id, 1545, year)
}

func Test_RatingCalculator_CalculateRatingsByYearFromTournamentResults_2TeamsWith2EarnedFinishes(t *testing.T) {
	// Given
	const year = 2018
	tournament := vbratings.Tournament{
		Id:   "3366b167a109496db63f43169e4ac1a7",
		Year: year,
	}

	result1 := vbratings.TournamentResult{
		Id:           "e57fd803a81349b69f0e7160fa13e919",
		Player1Id:    playerAId,
		Player2Id:    playerBId,
		EarnedFinish: 1,
		TournamentId: tournament.Id,
	}

	result2 := vbratings.TournamentResult{
		Id:           "d80ad4fe3bc2447abb44388702d0f791",
		Player1Id:    playerCId,
		Player2Id:    playerDId,
		EarnedFinish: 2,
		TournamentId: tournament.Id,
	}

	db := sqlite.NewInMemoryDB()
	tournamentRepository := sqlite.NewTournamentRepository(db)
	tournamentRepository.Create(tournament)
	tournamentRepository.AddTournamentResult(result1)
	tournamentRepository.AddTournamentResult(result2)
	matchRepository := sqlite.NewMatchRepository(db)
	playerRatingRepository := sqlite.NewPlayerRatingRepository(db)
	ratingCalculator := vbratings.NewRatingCalculator(
		matchRepository,
		tournamentRepository,
		playerRatingRepository,
	)

	// When
	totalCalculated := ratingCalculator.CalculateRatingsByYearFromTournamentResults(year)

	// Then
	assert.Equal(t, 4, totalCalculated)
	assertNewRating(t, playerRatingRepository, result1.Player1Id, 1561, year)
	assertNewRating(t, playerRatingRepository, result1.Player2Id, 1561, year)

	assertNewRating(t, playerRatingRepository, result2.Player1Id, 1529, year)
	assertNewRating(t, playerRatingRepository, result2.Player2Id, 1529, year)
}

//func Test_RatingCalculator_CalculateRatingsByYearFromTournamentResults_8TeamsWith5EarnedFinishes(t *testing.T) {
//	// Given
//	const year = 2018
//	tournament := vbratings.Tournament{
//		Id:   "3366b167a109496db63f43169e4ac1a7",
//		Year: year,
//	}
//
//	result1 := vbratings.TournamentResult{
//		Id:           "e57fd803a81349b69f0e7160fa13e919",
//		Player1Id:    playerAId,
//		Player2Id:    playerBId,
//		EarnedFinish: 1,
//		TournamentId: tournament.Id,
//	}
//
//	result2 := vbratings.TournamentResult{
//		Id:           "d80ad4fe3bc2447abb44388702d0f791",
//		Player1Id:    playerCId,
//		Player2Id:    playerDId,
//		EarnedFinish: 2,
//		TournamentId: tournament.Id,
//	}
//
//	result3a := vbratings.TournamentResult{
//		Id:           "65e604d014f24461893354ef5cf34f94",
//		Player1Id:    playerEId,
//		Player2Id:    playerFId,
//		EarnedFinish: 3,
//		TournamentId: tournament.Id,
//	}
//
//	result3b := vbratings.TournamentResult{
//		Id:           "05280c19859343f9a60c7752055bb3e3",
//		Player1Id:    playerGId,
//		Player2Id:    playerHId,
//		EarnedFinish: 3,
//		TournamentId: tournament.Id,
//	}
//
//	result5a := vbratings.TournamentResult{
//		Id:           "ef96a5918ffd49e78998e1b4fff3413f",
//		Player1Id:    playerIId,
//		Player2Id:    playerJId,
//		EarnedFinish: 5,
//		TournamentId: tournament.Id,
//	}
//
//	result5b := vbratings.TournamentResult{
//		Id:           "3d31348210044b799bce73a379be3a8e",
//		Player1Id:    playerKId,
//		Player2Id:    playerLId,
//		EarnedFinish: 5,
//		TournamentId: tournament.Id,
//	}
//
//	result5c := vbratings.TournamentResult{
//		Id:           "cd31348210044b799bce73a379be3a8e",
//		Player1Id:    playerMId,
//		Player2Id:    playerNId,
//		EarnedFinish: 5,
//		TournamentId: tournament.Id,
//	}
//
//	result5d := vbratings.TournamentResult{
//		Id:           "dd31348210044b799bce73a379be3a8e",
//		Player1Id:    playerOId,
//		Player2Id:    playerPId,
//		EarnedFinish: 5,
//		TournamentId: tournament.Id,
//	}
//	db := sqlite.NewInMemoryDB()
//	tournamentRepository := sqlite.NewTournamentRepository(db)
//	tournamentRepository.Create(tournament)
//	tournamentRepository.AddTournamentResult(result1)
//	tournamentRepository.AddTournamentResult(result2)
//	tournamentRepository.AddTournamentResult(result3a)
//	tournamentRepository.AddTournamentResult(result3b)
//	tournamentRepository.AddTournamentResult(result5a)
//	tournamentRepository.AddTournamentResult(result5b)
//	tournamentRepository.AddTournamentResult(result5c)
//	tournamentRepository.AddTournamentResult(result5d)
//	matchRepository := sqlite.NewMatchRepository(db)
//	playerRatingRepository := sqlite.NewPlayerRatingRepository(db)
//	ratingCalculator := vbratings.NewRatingCalculator(
//		matchRepository,
//		tournamentRepository,
//		playerRatingRepository,
//	)
//
//	// When
//	totalCalculated := ratingCalculator.CalculateRatingsByYearFromTournamentResults(year)
//
//	// Then
//	assert.Equal(t, 4, totalCalculated)
//	assertNewRating(t, playerRatingRepository, result1.Player1Id, 1561, year)
//	assertNewRating(t, playerRatingRepository, result1.Player2Id, 1561, year)
//
//	assertNewRating(t, playerRatingRepository, result2.Player1Id, 1545, year)
//	assertNewRating(t, playerRatingRepository, result2.Player2Id, 1545, year)
//
//	assertNewRating(t, playerRatingRepository, result3a.Player1Id, 1529, year)
//	assertNewRating(t, playerRatingRepository, result3a.Player2Id, 1529, year)
//	assertNewRating(t, playerRatingRepository, result3b.Player1Id, 1529, year)
//	assertNewRating(t, playerRatingRepository, result3b.Player2Id, 1529, year)
//
//	assertNewRating(t, playerRatingRepository, result5a.Player1Id, 1511, year)
//	assertNewRating(t, playerRatingRepository, result5a.Player2Id, 1511, year)
//	assertNewRating(t, playerRatingRepository, result5b.Player1Id, 1511, year)
//	assertNewRating(t, playerRatingRepository, result5b.Player2Id, 1511, year)
//	assertNewRating(t, playerRatingRepository, result5c.Player1Id, 1511, year)
//	assertNewRating(t, playerRatingRepository, result5c.Player2Id, 1511, year)
//	assertNewRating(t, playerRatingRepository, result5d.Player1Id, 1511, year)
//	assertNewRating(t, playerRatingRepository, result5d.Player2Id, 1511, year)
//}

func assertNewRating(t *testing.T, playerRatingRepository vbratings.PlayerRatingRepository, playerId string, expectedRating int, expectedYear int) {
	playerRatingResult, err := playerRatingRepository.GetPlayerRatingByYear(playerId, expectedYear)
	require.NoError(t, err)
	assertPlayerRating(t, playerRatingResult, 1500, expectedRating, expectedYear)
}

func Test_Prototype(t *testing.T) {
	// Given
	tournament := vbratings.Tournament{
		Id:   "3366b167a109496db63f43169e4ac1a7",
		Year: 2018,
	}

	result1 := vbratings.TournamentResult{
		Id:           "e57fd803a81349b69f0e7160fa13e919",
		Player1Id:    "0ddbff7e97384e3a8280503a6b483b87",
		Player2Id:    "74c26a34b97e4377a05f04f2fdeac3f0",
		EarnedFinish: 1,
		TournamentId: tournament.Id,
	}

	result2 := vbratings.TournamentResult{
		Id:           "d80ad4fe3bc2447abb44388702d0f791",
		Player1Id:    "a152dbcf009948d89d5e62642f9dda8b",
		Player2Id:    "e123ddb0b1ad443bb019b39e5d3bb36e",
		EarnedFinish: 2,
		TournamentId: tournament.Id,
	}

	result3a := vbratings.TournamentResult{
		Id:           "65e604d014f24461893354ef5cf34f94",
		Player1Id:    "0dbfed48bee240c2936a70a86663ed62",
		Player2Id:    "305f1a6f4ae64884956d0ed7ce13785b",
		EarnedFinish: 3,
		TournamentId: tournament.Id,
	}

	result3b := vbratings.TournamentResult{
		Id:           "05280c19859343f9a60c7752055bb3e3",
		Player1Id:    "8b3ac1934c6f4593bca0bb888196fa58",
		Player2Id:    "98c7546096ea4773b4c1bc5195b80cb7",
		EarnedFinish: 3,
		TournamentId: tournament.Id,
	}

	result5a := vbratings.TournamentResult{
		Id:           "ef96a5918ffd49e78998e1b4fff3413f",
		Player1Id:    "6395114fbb8b4e589cabe50addfcc7fd",
		Player2Id:    "4fc8124ccd314172b0a4085346ba303c",
		EarnedFinish: 5,
		TournamentId: tournament.Id,
	}

	result5b := vbratings.TournamentResult{
		Id:           "3d31348210044b799bce73a379be3a8e",
		Player1Id:    "44d08b8c252f47ed83a8677f6982af26",
		Player2Id:    "f6ab7b01caa94f18a87a69f3930be9bc",
		EarnedFinish: 5,
		TournamentId: tournament.Id,
	}

	result5c := vbratings.TournamentResult{
		Id:           "cd31348210044b799bce73a379be3a8e",
		Player1Id:    "c4d08b8c252f47ed83a8677f6982af26",
		Player2Id:    "c6ab7b01caa94f18a87a69f3930be9bc",
		EarnedFinish: 5,
		TournamentId: tournament.Id,
	}

	result5d := vbratings.TournamentResult{
		Id:           "dd31348210044b799bce73a379be3a8e",
		Player1Id:    "d4d08b8c252f47ed83a8677f6982af26",
		Player2Id:    "d6ab7b01caa94f18a87a69f3930be9bc",
		EarnedFinish: 5,
		TournamentId: tournament.Id,
	}

	origRatings1 := ratings{1500, 1500}
	origRatings2 := ratings{1500, 1500}
	origRatings3a := ratings{1500, 1500}
	origRatings3b := ratings{1500, 1500}
	origRatings5a := ratings{1500, 1500}
	origRatings5b := ratings{1500, 1500}
	origRatings5c := ratings{1500, 1500}
	origRatings5d := ratings{1500, 1500}

	ratings1 := &ratings{}
	ratings2 := &ratings{}
	ratings3a := &ratings{}
	ratings3b := &ratings{}
	ratings5a := &ratings{}
	ratings5b := &ratings{}
	ratings5c := &ratings{}
	ratings5d := &ratings{}

	*ratings1 = origRatings1
	*ratings2 = origRatings2
	*ratings3a = origRatings3a
	*ratings3b = origRatings3b
	*ratings5a = origRatings5a
	*ratings5b = origRatings5b
	*ratings5c = origRatings5c
	*ratings5d = origRatings5d

	results := []vbratings.TournamentResult{
		result1,
		result2,
		result3a,
		result3b,
		result5a,
		result5b,
		result5c,
		result5d,
	}

	var groupA group
	var groupB group

	earnedFinishA := 3
	earnedFinishB := 5

	for _, result := range results {
		if result.EarnedFinish == earnedFinishA {
			tempARatings := &ratings{
				getPlayerRating(result.Player1Id).Rating,
				getPlayerRating(result.Player2Id).Rating,
			}
			groupA = append(groupA, tempARatings)
		} else if result.EarnedFinish == earnedFinishB {
			tempBRatings := &ratings{
				getPlayerRating(result.Player1Id).Rating,
				getPlayerRating(result.Player2Id).Rating,
			}
			groupB = append(groupB, tempBRatings)
		}
	}

	fmt.Printf(" groupA: ")
	for _, value := range groupA {
		fmt.Printf(" %+v", *value)

	}
	fmt.Println()

	fmt.Printf(" groupB: ")
	for _, value := range groupB {
		fmt.Printf(" %+v", *value)

	}
	fmt.Println()

	// Every playoff team beats shadow teams
	shadowTeam1 := &ratings{1500, 1500}
	shadowTeam2 := &ratings{1500, 1500}
	shadowTeam3 := &ratings{1500, 1500}
	beat(group{ratings1, ratings2, ratings3a, ratings3b, ratings5a, ratings5b, ratings5c, ratings5d}, group{shadowTeam1})
	beat(group{ratings1, ratings2, ratings3a, ratings3b, ratings5a, ratings5b, ratings5c, ratings5d}, group{shadowTeam2})
	beat(group{ratings1, ratings2, ratings3a, ratings3b, ratings5a, ratings5b, ratings5c, ratings5d}, group{shadowTeam3})

	// Top beats immediate lower
	//beat(group{ratings1}, group{ratings2})
	//beat(group{ratings2}, group{ratings3a, ratings3b})
	//beat(group{ratings3a, ratings3b}, group{ratings5a, ratings5b, ratings5c, ratings5d})

	// Top beats 2 below in 2 separate matches
	//beat(group{ratings1}, group{ratings2})
	//beat(group{ratings1}, group{ratings3a, ratings3b})
	//beat(group{ratings2}, group{ratings3a, ratings3b})
	//beat(group{ratings2}, group{ratings5a, ratings5b, ratings5c, ratings5d})
	//beat(group{ratings3a, ratings3b}, group{ratings5a, ratings5b, ratings5c, ratings5d})

	// Top beats 2 below
	beat(group{ratings1}, group{ratings2})
	//beat(group{ratings1}, group{ratings2, ratings3a, ratings3b})
	//beat(group{ratings2}, group{ratings3a, ratings3b, ratings5a, ratings5b, ratings5c, ratings5d})
	//beat(group{ratings3a, ratings3b}, group{ratings5a, ratings5b, ratings5c, ratings5d})

	fmt.Printf(" 1: %+v - %+v\n", origRatings1, ratings1)
	fmt.Printf(" 2: %+v - %+v\n", origRatings2, ratings2)
	fmt.Printf("3a: %+v - %+v\n", origRatings3a, ratings3a)
	fmt.Printf("3b: %+v - %+v\n", origRatings3b, ratings3b)
	fmt.Printf("5a: %+v - %+v\n", origRatings5a, ratings5a)
	fmt.Printf("5b: %+v - %+v\n", origRatings5b, ratings5b)
	fmt.Printf("5c: %+v - %+v\n", origRatings5c, ratings5c)
	fmt.Printf("5d: %+v - %+v\n", origRatings5d, ratings5d)
}

func getPlayerRating(playerId string) vbratings.PlayerRating {
	return vbratings.PlayerRating{
		PlayerId:   playerId,
		Year:       2018,
		SeedRating: 1500,
		Rating:     1500,
	}
}

type group []*ratings

func beat(g1 group, g2 group) {
	var r1 ratings
	var r2 ratings

	for i := range g1 {
		r1 = append(r1, *g1[i]...)
	}

	for i := range g2 {
		r2 = append(r2, *g2[i]...)
	}

	r1.beat(&r2)

	j := 0
	for i := range g1 {
		*g1[i] = ratings{r1[j], r1[j+1]}
		j += 2
	}

	j = 0
	for i := range g2 {
		*g2[i] = ratings{r2[j], r2[j+1]}
		j += 2
	}
}

type ratings []int

func (r1 *ratings) beat(r2 *ratings) {
	*r1, *r2 = duelingCalculator.GetNewRatings(*r1, *r2, 1, 0)
}

var ratingCalculator = skill.NewEloCalculator(32)
var duelingCalculator = skill.NewDuelingCalculator(ratingCalculator)

func assertPlayerRating(t *testing.T, playerRating *vbratings.PlayerRating, expectedSeedRating int, expectedRating int, expectedYear int) {
	assert.Equal(t, expectedSeedRating, playerRating.SeedRating)
	assert.Equal(t, expectedRating, playerRating.Rating)
	assert.Equal(t, expectedYear, playerRating.Year)
}
