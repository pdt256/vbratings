package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/pdt256/vbratings/bvbinfo"
	"github.com/pdt256/vbratings/pkg/uuid"
	"github.com/pdt256/vbratings/sqlite"
)

func main() {
	fmt.Println("BVBInfo Importer")
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	dbPath := flag.String("dbPath", "./_data/vb.db", "sqlite db path")

	flag.Parse()

	db := sqlite.NewFileDB(*dbPath)
	tournamentRepository := sqlite.NewTournamentRepository(db)
	matchRepository := sqlite.NewMatchRepository(db)
	playerRepository := sqlite.NewPlayerRepository(db)
	bvbInfoRepository := bvbinfo.NewRepositoryWithCaching(db)

	uuidGenerator := uuid.NewService()

	importer := bvbinfo.NewImporter(
		tournamentRepository,
		matchRepository,
		playerRepository,
		bvbInfoRepository,
		uuidGenerator,
	)

	fmt.Println("Importing Matches")
	totalTournaments, totalMatches, totalPlayers := importer.ImportAllSeasons()
	fmt.Printf("\n%d total tournaments imported\n", totalTournaments)
	fmt.Printf("%d total matches imported\n", totalMatches)
	fmt.Printf("%d total players imported\n", totalPlayers)
}
