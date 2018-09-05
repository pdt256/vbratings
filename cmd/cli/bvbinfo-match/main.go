package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/namsral/flag"
	"github.com/pdt256/vbscraper"
)

func main() {
	fmt.Println("BVBInfo Match Importer")
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	dbPath := flag.String("dbPath", "./_data/vb.db", "sqlite db path")
	shouldInitDb := flag.Bool("init", false, "init db")
	tournamentUrl := flag.String("tournamentUrl", "", "tournament url")
	seasonUrl := flag.String("seasonUrl", "", "season url")
	allSeasons := flag.Bool("allSeasons", false, "load all seasons")

	flag.Parse()

	db, err := sql.Open("sqlite3", *dbPath)
	if err != nil {
		log.Fatal(err)
	}
	matchRepository := vbscraper.NewSqliteMatchRepository(db)
	playerRepository := vbscraper.NewSqlitePlayerRepository(db)

	if *shouldInitDb {
		fmt.Println("Initializing database")
		matchRepository.InitDB()
		return
	}

	importer := vbscraper.NewBvbInfoImporter(matchRepository, playerRepository)

	if *tournamentUrl != "" {
		total := importTournament(*tournamentUrl, importer)
		printTotal(total)
		return
	}

	if *seasonUrl != "" {
		total := importSeason(*seasonUrl, importer)
		printTotal(total)
		return
	}

	if *allSeasons {
		total := importAllSeasons(importer)
		printTotal(total)
		return
	}
}

func printTotal(total int) {
	fmt.Printf("%d matches imported\n", total)
}

func importTournament(tournamentUrl string, importer *vbscraper.BvbinfoImporter) int {
	fmt.Printf("Importing Tournament: %s\n", tournamentUrl)
	tournamentResponse, err := http.Get(tournamentUrl)
	if err != nil {
		log.Fatal(err)
	}

	defer tournamentResponse.Body.Close()

	totalMatchesImported := importer.ImportMatches(tournamentResponse.Body)
	return totalMatchesImported
}

func importSeason(seasonUrl string, importer *vbscraper.BvbinfoImporter) int {
	fmt.Printf("Importing Season: %s\n", seasonUrl)
	seasonResponse, err := http.Get(seasonUrl)
	if err != nil {
		log.Fatal(err)
	}
	tournaments := vbscraper.GetTournaments(seasonResponse.Body)
	seasonResponse.Body.Close()

	totalMatchesImported := 0
	for _, tournament := range tournaments {
		tournamentUrl := fmt.Sprintf("http://bvbinfo.com/Tournament.asp?ID=%s&Process=Matches", tournament.BvbId)
		totalMatchesImported += importTournament(tournamentUrl, importer)
	}
	return totalMatchesImported
}

func importAllSeasons(importer *vbscraper.BvbinfoImporter) int {
	allSeasonsUrl := "http://bvbinfo.com/season.asp"
	allSeasonsResponse, err := http.Get(allSeasonsUrl)
	if err != nil {
		log.Fatal(err)
	}
	seasons := vbscraper.GetSeasons(allSeasonsResponse.Body)
	allSeasonsResponse.Body.Close()

	totalMatchesImported := 0
	for _, season := range seasons {
		seasonUrl := fmt.Sprintf("http://bvbinfo.com/Season.asp?AssocID=%s&Year=%s", season.AssocID, season.Year)
		totalMatchesImported += importSeason(seasonUrl, importer)
		fmt.Printf("%d matches imported\n", totalMatchesImported)
	}
	return totalMatchesImported
}
