package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/pdt256/vbratings"
	"github.com/pdt256/vbratings/sqlite"
)

func main() {
	fmt.Println("Volleyball Ratings Calculator")
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	dbPath := flag.String("dbPath", "./_data/vb.db", "sqlite db path")
	allYears := flag.Bool("allYears", false, "calculate for all years")
	year := flag.Int("year", 2018, "year")
	shouldInitDb := flag.Bool("init", false, "init db")

	flag.Parse()

	db := sqlite.NewFileDB(*dbPath)
	matchRepository := sqlite.NewMatchRepository(db)
	playerRatingRepository := sqlite.NewPlayerRatingRepository(db)

	if *shouldInitDb {
		fmt.Println("Initializing player_rating DB")
		playerRatingRepository.InitDB()
		return
	}

	var totalCalculated int
	if *allYears {
		years := []int{
			2000, 2001, 2003, 2004, 2005, 2006, 2007, 2008, 2009,
			2010, 2011, 2012, 2013, 2014, 2015, 2016, 2017, 2018,
		}

		for _, singleYear := range years {
			ratingCalculator := vbratings.NewRatingCalculator(matchRepository, playerRatingRepository)
			totalCalculated += ratingCalculator.CalculateRatingsByYear(singleYear)
			fmt.Print(".")
		}
		fmt.Println()
	} else {
		ratingCalculator := vbratings.NewRatingCalculator(matchRepository, playerRatingRepository)
		totalCalculated = ratingCalculator.CalculateRatingsByYear(*year)
	}

	fmt.Printf("%d ratings calculated\n", totalCalculated)
}
