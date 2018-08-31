package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/namsral/flag"
	"github.com/pdt256/vbscraper"
)

func main() {
	fmt.Println("BVBInfo Importer")
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.PanicOnError)

	dbPath := flag.String(
		"dbPath",
		"./_data/vb.db",
		"sqlite db path",
	)

	shouldInitDb := flag.Bool(
		"init",
		false,
		"init db",
	)

	tournamentUrl := flag.String(
		"tournamentUrl",
		"",
		"tournament url",
	)

	seasonUrl := flag.String(
		"seasonUrl",
		"",
		"season url",
	)
	flag.Parse()

	matchRepository := vbscraper.NewSqliteMatchRepository(*dbPath)

	if *shouldInitDb {
		fmt.Println("Initializing database")
		matchRepository.InitDB()
		return
	}

	importer := vbscraper.NewBvbInfoImporter(matchRepository)

	if *tournamentUrl != "" {
		importTournament(*tournamentUrl, importer)
		return
	}

	if *seasonUrl != "" {
		seasonResponse, err := http.Get(*seasonUrl)
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
		fmt.Printf("Done! (%d) matches imported\n", totalMatchesImported)
	}
}

func importTournament(tournamentUrl string, importer *vbscraper.BvbinfoImporter) int {
	fmt.Printf("Importing Tournament: %s ", tournamentUrl)
	tournamentResponse, err := http.Get(tournamentUrl)
	if err != nil {
		log.Fatal(err)
	}

	defer tournamentResponse.Body.Close()

	totalMatchesImported := importer.ImportMatches(tournamentResponse.Body)
	fmt.Printf("(%d) matches imported\n", totalMatchesImported)
	return totalMatchesImported
}
