package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/namsral/flag"
	"github.com/pdt256/vbscraper"
)

func main() {
	fmt.Println("BVBInfo Ratings")
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	dbPath := flag.String("dbPath", "./_data/vb.db", "sqlite db path")
	year := flag.Int("year", 2018, "year")
	shouldInitDb := flag.Bool("init", false, "init db")

	flag.Parse()

	db, err := sql.Open("sqlite3", *dbPath)
	if err != nil {
		log.Fatal(err)
	}
	matchRepository := vbscraper.NewSqliteMatchRepository(db)
	playerRatingRepository := vbscraper.NewSqlitePlayerRatingRepository(db)

	if *shouldInitDb {
		fmt.Println("Initializing player_rating database")
		playerRatingRepository.InitDB()
		return
	}

	ratingCalculator := vbscraper.NewRatingCalculator(matchRepository, playerRatingRepository)

	totalCalculated := ratingCalculator.CalculateRatingsByYear(*year)

	fmt.Printf("%d ratings calculated\n", totalCalculated)
}
