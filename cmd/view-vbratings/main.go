package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/pdt256/vbratings"
	"github.com/pdt256/vbratings/sqlite"
	"github.com/spf13/cobra"
)

func main() {
	var dbPath string

	var year int
	var gender string
	var limit int
	var cmdTopPlayers = &cobra.Command{
		Use:   "topPlayers ",
		Short: "List Top Players By Year",
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Top %d %s Players in %d\n", limit, gender, year)

			db := sqlite.NewFileDB(dbPath)
			playerRatingRepository := sqlite.NewPlayerRatingRepository(db)

			playerRatings := playerRatingRepository.GetTopPlayerRatings(year, vbratings.GenderFromString(gender), limit)
			showTable(playerRatings)
		},
	}
	cmdTopPlayers.Flags().IntVarP(&year, "year", "y", 2018, "year")
	cmdTopPlayers.Flags().StringVarP(&gender, "gender", "g", "male", "gender")
	cmdTopPlayers.Flags().IntVarP(&limit, "limit", "l", 10, "limit")

	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.Flags().StringVarP(&dbPath, "dbPath", "d", "./_data/vb.db", "sqlite db path")
	rootCmd.AddCommand(cmdTopPlayers)
	rootCmd.Execute()
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
