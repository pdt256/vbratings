package main

import (
	"fmt"
	"os"

	"github.com/namsral/flag"
	"github.com/pdt256/vbratings"
	"github.com/pdt256/vbratings/sqlite"
)

func main() {
	fmt.Println("Volleyball Ratings")
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	dbPath := flag.String("dbPath", "./_data/vb.db", "sqlite db path")
	year := flag.Int("year", 2018, "year")
	shouldInitDb := flag.Bool("init", false, "init db")

	flag.Parse()

	db := sqlite.NewFileDB(*dbPath)
	matchRepository := sqlite.NewMatchRepository(db)
	playerRatingRepository := sqlite.NewPlayerRatingRepository(db)

	if *shouldInitDb {
		fmt.Println("Initializing player_rating database")
		playerRatingRepository.InitDB()
		return
	}

	ratingCalculator := vbratings.NewRatingCalculator(matchRepository, playerRatingRepository)

	totalCalculated := ratingCalculator.CalculateRatingsByYear(*year)

	fmt.Printf("%d ratings calculated\n", totalCalculated)
}
