package generator_test

import (
	"bytes"
	"go/parser"
	"go/token"
	"testing"

	"github.com/graph-gophers/graphql-go"
	"github.com/pdt256/vbratings/graphql/generator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_GraphQLSchema(t *testing.T) {
	// Given
	domainFilePath := `./testdata/simple_domain.go`
	domainFs := token.NewFileSet()
	domainNode, _ := parser.ParseFile(domainFs, domainFilePath, nil, parser.ParseComments)
	domainRoot := generator.ParseDomain(domainNode)

	entitiesFilePath := `./testdata/simple_entities.go`
	entitiesFs := token.NewFileSet()
	entitiesNode, _ := parser.ParseFile(entitiesFs, entitiesFilePath, nil, parser.ParseComments)
	domainRoot.Entities = append(
		domainRoot.Entities,
		generator.ParseEntities(entitiesNode).Entities...,
	)

	var schema bytes.Buffer

	// When
	domainRoot.GraphQLSchema(&schema)

	// Then
	expectedSchema := `schema {
	query: Query
	mutation: Mutation
}

type Query {
	simpleDomainQueries: SimpleDomainQueries
}

type SimpleDomainQueries {
	# Query 1 Doc
	query1: Boolean!
	# Query 2 Doc
	# Second Line
	query2: String!
	query3(
		one: Int!
		two: String!
		three: Boolean!
	): Int!
	query4: Struct1!
}

type Mutation {
	simpleDomainCommands: SimpleDomainCommands
}

type SimpleDomainCommands {
	# Command 1 Doc
	command1: Boolean!
	# Command 2 Doc
	command2: Boolean!
	command3(
		one: Int!
		two: String!
		three: Boolean!
	): Boolean!
}

type Struct1 {
	one: Int!
	two: String!
	three: Boolean!
}

type ParentStruct {
	childStruct: ChildStruct!
}

type ChildStruct {
	name: String!
}
`
	_, err := graphql.ParseSchema(schema.String(), nil)
	require.NoError(t, err)
	assert.Equal(t, expectedSchema, schema.String())
}
