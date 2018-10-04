package graphql

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/pdt256/vbratings"
	"github.com/pdt256/vbratings/app"
)

func NewHandler(app *app.App) *relay.Handler {
	resolver := NewResolver(app)
	schema := graphql.MustParseSchema(getSchemaString(), resolver)
	return &relay.Handler{Schema: schema}
}

type resolver struct {
	app *app.App
}

func NewResolver(app *app.App) *resolver {
	return &resolver{app}
}

func (r *resolver) PlayerQueries() *PlayerQueries {
	return &PlayerQueries{r.app}
}

func (r *resolver) PlayerRatingQueries() *PlayerRatingQueries {
	return &PlayerRatingQueries{r.app}
}

func (r *resolver) PlayerCommands() *PlayerCommands {
	return &PlayerCommands{r.app}
}

func (r *resolver) PlayerRatingCommands() *PlayerRatingCommands {
	return &PlayerRatingCommands{r.app}
}

// Domain Queries

type PlayerQueries struct {
	app *app.App
}

func (q *PlayerQueries) GetPlayer(args struct{ Id string }) (*PlayerResolver, error) {
	player, err := q.app.Player.GetPlayer(args.Id)

	return &PlayerResolver{*player}, err
}

type PlayerRatingQueries struct {
	app *app.App
}

func (q *PlayerRatingQueries) GetTopPlayerRatings(args struct {
	Year   int32
	Gender string
	Limit  int32
}) []*PlayerAndRatingResolver {
	playerAndRatings := q.app.PlayerRating.GetTopPlayerRatings(
		int(args.Year),
		args.Gender,
		int(args.Limit))

	var r []*PlayerAndRatingResolver
	for _, value := range playerAndRatings {
		r = append(r, &PlayerAndRatingResolver{value})
	}

	return r
}

// Domain Commands

type PlayerCommands struct {
	app *app.App
}

type playerCreateInput struct {
	Id     string
	Name   string
	Gender string
}

func (c *PlayerCommands) Create(args playerCreateInput) (bool, error) {
	return true, c.app.Player.Create(
		args.Id,
		args.Name,
		args.Gender)
}

type PlayerRatingCommands struct {
	app *app.App
}

func (c *PlayerRatingCommands) Create(args struct {
	PlayerId     string
	Year         int32
	SeedRating   int32
	Rating       int32
	TotalMatches int32
}) (bool, error) {
	return true, c.app.PlayerRating.Create(
		args.PlayerId,
		int(args.Year),
		int(args.SeedRating),
		int(args.Rating),
		int(args.TotalMatches))
}

// Entity Resolvers

type PlayerAndRatingResolver struct {
	playerAndRating vbratings.PlayerAndRating
}

func (p *PlayerAndRatingResolver) Player() *PlayerResolver {
	return &PlayerResolver{p.playerAndRating.Player}
}

func (p *PlayerAndRatingResolver) PlayerRating() *PlayerRatingResolver {
	return &PlayerRatingResolver{p.playerAndRating.PlayerRating}
}

type PlayerResolver struct {
	player vbratings.Player
}

func (p *PlayerResolver) Id() string {
	return p.player.Id
}

func (p *PlayerResolver) Name() string {
	return p.player.Name
}

func (p *PlayerResolver) Gender() string {
	return p.player.Gender
}

type PlayerRatingResolver struct {
	playerRating vbratings.PlayerRating
}

func (pr *PlayerRatingResolver) PlayerId() string {
	return pr.playerRating.PlayerId
}

func (pr *PlayerRatingResolver) Year() int32 {
	return int32(pr.playerRating.Year)
}

func (pr *PlayerRatingResolver) SeedRating() int32 {
	return int32(pr.playerRating.SeedRating)
}

func (pr *PlayerRatingResolver) Rating() int32 {
	return int32(pr.playerRating.Rating)
}

func (pr *PlayerRatingResolver) TotalMatches() int32 {
	return int32(pr.playerRating.TotalMatches)
}

func getSchemaString() string {
	return `
		schema {
	    	query: Query
			mutation: Mutation
		}

		type Query {
			playerQueries: PlayerQueries
			playerRatingQueries: PlayerRatingQueries
		}

		type PlayerQueries {
			# Get Player by id
			getPlayer(
				id: String!
			): Player!
		}

		type PlayerRatingQueries {
			# Top player ratings by year and gender
    		getTopPlayerRatings(
				year: Int!
				gender: String!
				limit: Int!
			): [PlayerAndRating]!
		}

		type Mutation {
			playerCommands: PlayerCommands
			playerRatingCommands: PlayerRatingCommands
		}

		type PlayerCommands {
			create(
				id: String!
				name: String!
				gender: String!
			): Boolean!
		}

		type PlayerRatingCommands {
			# Create player rating
			create(
				playerId: String!
				year: Int!
				seedRating: Int!
				rating: Int!
				totalMatches: Int!
			): Boolean!
		}

		type Player {
			Id: String!
			Name: String!
			Gender: String!
		}

		type PlayerRating {
			PlayerId: String!
			Year: Int!
			SeedRating: Int!
			Rating: Int!
			TotalMatches: Int!			
		}

		type PlayerAndRating {
			player: Player!
			playerRating: PlayerRating
		}
	`
}
