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

	matchUrl := flag.String(
		"matchUrl",
		"",
		"match url",
	)
	flag.Parse()

	matchRepository := vbscraper.NewSqliteMatchRepository(*dbPath)

	if *shouldInitDb {
		fmt.Println("Initializing database")
		matchRepository.InitDB()
		return
	}

	importer := vbscraper.NewBvbInfoImporter(matchRepository)

	if *matchUrl != "" {
		matchesResponse, err := http.Get(*matchUrl)
		if err != nil {
			log.Fatal(err)
		}
		defer matchesResponse.Body.Close()
		importer.ImportMatches(matchesResponse.Body)
	}
}
