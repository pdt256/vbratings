package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	vbgraphql "github.com/pdt256/vbratings/graphql"
	"github.com/pdt256/vbratings/sqlite"
)

func main() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	dbPath := flag.String("dbPath", "./_data/vb.db", "sqlite db path")
	port := flag.Int("port", 8080, "port")
	flag.Parse()

	fmt.Println("Volleyball Ratings GraphQL")
	fmt.Printf("Starting on port %d\n", *port)

	db := sqlite.NewFileDB(*dbPath)
	playerRatingRepository := sqlite.NewPlayerRatingRepository(db)

	query := vbgraphql.NewQuery(playerRatingRepository)
	schema := graphql.MustParseSchema(getSchemaString(), query)
	http.Handle("/query", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}

func getSchemaString() string {
	b, err := ioutil.ReadFile("./graphql/schema.graphql")
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}
