package main

import (
	"bytes"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/pdt256/vbratings/cbva"
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

	fmt.Println("Importing Tournaments")

	postData := `{"id":"A14CC0CB1B90719A"}`
	req, _ := http.NewRequest("POST", "https://cbva.com/Results/GetTournamentTeamResult", bytes.NewReader([]byte(postData)))
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	tournamentResponse, _ := client.Do(req)
	defer tournamentResponse.Body.Close()

	importer := cbva.NewImporter(tournamentRepository, playerRepository)
	totalTournamentResultsImported := importer.ImportTournamentResults(tournamentResponse.Body)
	fmt.Printf("\n%d tournament results imported\n", totalTournamentResultsImported)
}
