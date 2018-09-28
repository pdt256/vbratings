package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/pdt256/vbratings/app"
	vbgraphql "github.com/pdt256/vbratings/graphql"
)

func main() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	dbPath := flag.String("dbPath", "./_data/vb.db", "sqlite db path")
	port := flag.Int("port", 8080, "port")
	flag.Parse()

	fmt.Println("Volleyball Ratings GraphQL")
	fmt.Printf("Starting on port %d\n", *port)

	configuration := app.NewConfiguration(*dbPath)
	application := app.New(configuration)
	graphqlHandler := vbgraphql.NewGraphQLHandler(application)
	http.Handle("/query", graphqlHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
