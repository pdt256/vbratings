package graphql

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/pdt256/vbratings"
	"github.com/pdt256/vbratings/app"
)

type PlayerRatingResolver struct {
	pr vbratings.PlayerAndRating
}

func (p *PlayerRatingResolver) Rating() int32 {
	return int32(p.pr.Rating)
}

func (p *PlayerRatingResolver) PlayerName() string {
	return p.pr.Name
}

func (p *PlayerRatingResolver) TotalMatches() int32 {
	return int32(p.pr.TotalMatches)
}

type query struct {
	app *app.App
}

func NewQuery(app *app.App) *query {
	return &query{app}
}

func (q *query) GetTopPlayerRatings(args struct {
	Year   int32
	Gender string
	Limit  int32
}) []*PlayerRatingResolver {
	playerAndRatings := q.app.PlayerRating.GetTopPlayerRatings(
		int(args.Year),
		args.Gender,
		int(args.Limit))

	var r []*PlayerRatingResolver
	for _, value := range playerAndRatings {
		r = append(r, &PlayerRatingResolver{value})
	}

	return r
}

func NewGraphQLHandler(app *app.App) *relay.Handler {
	query := NewQuery(app)
	schema := graphql.MustParseSchema(getSchemaString(), query)
	return &relay.Handler{Schema: schema}
}

func getSchemaString() string {
	return `
		schema {
	    	query: Query
		}

		type Query {
			# Top player ratings by year and gender
    		getTopPlayerRatings(
				year: Int!
				gender: String!
				limit: Int!
			): [PlayerRating]!
		}

		type PlayerRating {
    		rating: Int!
    		playerName: String!
    		totalMatches: Int!
		}
	`
}
