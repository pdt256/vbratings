package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/namsral/flag"
	"github.com/pdt256/vbratings/bvbinfo"
	"github.com/pdt256/vbratings/sqlite"
)

func main() {
	fmt.Println("BVBInfo Player Importer")
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	dbPath := flag.String("dbPath", "./_data/vb.db", "sqlite db path")
	shouldInitDb := flag.Bool("init", false, "init db")

	flag.Parse()

	db, err := sql.Open("sqlite3", *dbPath)
	checkError(err)

	matchRepository := sqlite.NewMatchRepository(db)
	playerRepository := sqlite.NewPlayerRepository(db)

	if *shouldInitDb {
		fmt.Println("Initializing player database")
		playerRepository.InitDB()
		return
	}

	totalImported := 0
	for _, playerId := range matchRepository.GetAllPlayerIds() {
		playerUrl := fmt.Sprintf("http://bvbinfo.com/player.asp?ID=%d", playerId)

		playerResponse, err := http.Get(playerUrl)
		checkError(err)

		player := bvbinfo.GetPlayer(playerResponse.Body, playerId)
		playerResponse.Body.Close()

		playerRepository.Create(player)
		totalImported++
		fmt.Print(".")
	}

	fmt.Printf("\n%d players imported\n", totalImported)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
