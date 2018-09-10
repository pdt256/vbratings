package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/namsral/flag"
	"github.com/olekukonko/tablewriter"
	"github.com/pdt256/vbratings"
	"github.com/pdt256/vbratings/sqlite"
)

func main() {
	fmt.Println("Volleyball Ratings")
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	dbPath := flag.String("dbPath", "./_data/vb.db", "sqlite db path")
	topPlayers := flag.Bool("topPlayers", false, "view top players")
	year := flag.Int("year", 2018, "year")
	gender := flag.String("gender", "male", "gender (male, female)")
	limit := flag.Int("limit", 10, "number of players to show")

	flag.Parse()

	db := sqlite.NewFileDB(*dbPath)
	playerRatingRepository := sqlite.NewPlayerRatingRepository(db)

	if *topPlayers {
		playerRatings := playerRatingRepository.GetTopPlayerRatings(*year, vbratings.GenderFromString(*gender), *limit)
		showTable(playerRatings)
	} else {
		fmt.Println("No query selected")
	}
}

func showTable(playerAndRatings []vbratings.PlayerAndRating) {
	var data [][]string

	for _, playerAndRating := range playerAndRatings {
		data = append(data, []string{
			strconv.Itoa(playerAndRating.Rating),
			playerAndRating.Name,
			strconv.Itoa(playerAndRating.TotalMatches),
		})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Rating", "Name", "TotalMatches"})

	for _, v := range data {
		table.Append(v)
	}

	table.Render()
}
