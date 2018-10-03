package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/pdt256/vbratings/cbva"
	"github.com/pdt256/vbratings/pkg/uuid"
	"github.com/pdt256/vbratings/sqlite"
)

func main() {
	fmt.Println("CBVA Importer")
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	dbPath := flag.String("dbPath", "./_data/vb.db", "sqlite db path")

	flag.Parse()

	db := sqlite.NewFileDB(*dbPath)
	tournamentRepository := sqlite.NewTournamentRepository(db)
	playerRepository := sqlite.NewPlayerRepository(db)
	cbvaRepository := cbva.NewSqliteRepository(db)
	uuidGenerator := uuid.NewService()

	importer := cbva.NewImporter(
		tournamentRepository,
		playerRepository,
		cbvaRepository,
		uuidGenerator,
	)

	fmt.Println("Importing Tournaments")
	totalResults, totalPlayers := importer.ImportAllTournaments()
	fmt.Printf("\n%d tournament results imported\n", totalResults)
	fmt.Printf("%d players imported\n", totalPlayers)
}
