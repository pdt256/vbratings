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

func Test_Parse_Queries(t *testing.T) {
	// Given
	filePath := `./testdata/simple_domain_queries.go`
	fs := token.NewFileSet()
	node, err := parser.ParseFile(fs, filePath, nil, parser.ParseComments)
	require.NoError(t, err)

	// When
	domainRoot := generator.ParseDomain(node)

	// Then
	domain := domainRoot.Domains[0]
	assert.Equal(t, "SimpleDomainQueries", domain.Name)

	queryWithNoParams := domain.Queries[0]
	assert.Equal(t, "QueryWithNoParams", queryWithNoParams.Name)
	assert.Equal(t, 0, len(queryWithNoParams.Params))
	assert.Equal(t, "bool", queryWithNoParams.ReturnTypes[0].Type)

	queryWithDoc := domain.Queries[1]
	assert.Equal(t, "QueryWithDoc", queryWithDoc.Name)
	assert.Equal(t, "// Line 1", queryWithDoc.Docs[0])
	assert.Equal(t, "// Line 2", queryWithDoc.Docs[1])

	queryWithParams := domain.Queries[2]
	assert.Equal(t, "QueryWithParams", queryWithParams.Name)
	assert.Equal(t, "oneInt", queryWithParams.Params[0].Name)
	assert.Equal(t, "int", queryWithParams.Params[0].Type)
	assert.Equal(t, "twoString", queryWithParams.Params[1].Name)
	assert.Equal(t, "string", queryWithParams.Params[1].Type)
	assert.Equal(t, "threeBool", queryWithParams.Params[2].Name)
	assert.Equal(t, "bool", queryWithParams.Params[2].Type)
	assert.Equal(t, "bool", queryWithParams.ReturnTypes[0].Type)

	queryWithArrayIdentReturn := domain.Queries[3]
	assert.Equal(t, "QueryWithArrayIdentReturn", queryWithArrayIdentReturn.Name)
	assert.Equal(t, "[]bool", queryWithArrayIdentReturn.ReturnTypes[0].Type)

	queryWithStructReturn := domain.Queries[4]
	assert.Equal(t, "QueryWithStructReturn", queryWithStructReturn.Name)
	assert.Equal(t, "", queryWithStructReturn.ReturnTypes[0].Package)
	assert.Equal(t, "SimpleStruct", queryWithStructReturn.ReturnTypes[0].Type)

	queryWithSelectorStructReturn := domain.Queries[5]
	assert.Equal(t, "QueryWithSelectorStructReturn", queryWithSelectorStructReturn.Name)
	assert.Equal(t, "sample", queryWithSelectorStructReturn.ReturnTypes[0].Package)
	assert.Equal(t, "SimpleStruct", queryWithSelectorStructReturn.ReturnTypes[0].Type)

	queryWithArrayStructReturn := domain.Queries[6]
	assert.Equal(t, "QueryWithArrayStructReturn", queryWithArrayStructReturn.Name)
	assert.Equal(t, "", queryWithArrayStructReturn.ReturnTypes[0].Package)
	assert.Equal(t, "[]SimpleStruct", queryWithArrayStructReturn.ReturnTypes[0].Type)

	queryWithArraySelectorStructReturn := domain.Queries[7]
	assert.Equal(t, "QueryWithArraySelectorStructReturn", queryWithArraySelectorStructReturn.Name)
	assert.Equal(t, "sample", queryWithArraySelectorStructReturn.ReturnTypes[0].Package)
	assert.Equal(t, "[]SimpleStruct", queryWithArraySelectorStructReturn.ReturnTypes[0].Type)
}

func Test_Parse_Commands(t *testing.T) {
	// Given
	filePath := `./testdata/simple_domain_commands.go`
	fs := token.NewFileSet()
	node, err := parser.ParseFile(fs, filePath, nil, parser.ParseComments)
	require.NoError(t, err)

	// When
	domainRoot := generator.ParseDomain(node)

	// Then
	domain := domainRoot.Domains[0]
	assert.Equal(t, "SimpleDomainCommands", domain.Name)

	commandWithNoReturn := domain.Commands[0]
	assert.Equal(t, "CommandWithNoReturn", commandWithNoReturn.Name)
	assert.Equal(t, 0, len(commandWithNoReturn.Params))
	assert.Equal(t, 0, len(commandWithNoReturn.ReturnTypes))

	commandReturnsError := domain.Commands[1]
	assert.Equal(t, "CommandReturnsError", commandReturnsError.Name)
	assert.Equal(t, 0, len(commandReturnsError.Params))
	assert.Equal(t, "error", commandReturnsError.ReturnTypes[0].Type)

	commandWithParams := domain.Commands[2]
	assert.Equal(t, "oneInt", commandWithParams.Params[0].Name)
	assert.Equal(t, "int", commandWithParams.Params[0].Type)
	assert.Equal(t, "twoString", commandWithParams.Params[1].Name)
	assert.Equal(t, "string", commandWithParams.Params[1].Type)
	assert.Equal(t, "threeBool", commandWithParams.Params[2].Name)
}

func Test_GraphQLSchema(t *testing.T) {
	// Given
	filePath := `./testdata/simple_domain.go`
	fs := token.NewFileSet()
	node, _ := parser.ParseFile(fs, filePath, nil, parser.ParseComments)
	domainRoot := generator.ParseDomain(node)
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
`
	assert.Equal(t, expectedSchema, schema.String())
	_, err := graphql.ParseSchema(schema.String(), nil)
	require.NoError(t, err)
}
