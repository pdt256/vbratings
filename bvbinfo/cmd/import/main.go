package main

import (
	"fmt"
	"os"

	"github.com/namsral/flag"
	"github.com/pdt256/vbratings/bvbinfo"
	"github.com/pdt256/vbratings/sqlite"
)

func main() {
	fmt.Println("BVBInfo Importer")
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	dbPath := flag.String("dbPath", "./_data/vb.db", "sqlite db path")
	shouldInitDb := flag.Bool("init", false, "init db")

	flag.Parse()

	db := sqlite.NewFileDB(*dbPath)
	matchRepository := sqlite.NewMatchRepository(db)
	playerRepository := sqlite.NewPlayerRepository(db)

	if *shouldInitDb {
		fmt.Println("Initializing DB")
		matchRepository.InitDB()
		playerRepository.InitDB()
		return
	}

	importer := bvbinfo.NewImporter(matchRepository, playerRepository)

	fmt.Println("Importing Matches")
	totalMatchesImported := importer.ImportAllSeasons()
	fmt.Printf("\n%d matches imported\n", totalMatchesImported)

	fmt.Println("Importing Players")
	totalPlayersImported := importer.ImportAllPlayers()
	fmt.Printf("\n%d players imported\n", totalPlayersImported)
}
