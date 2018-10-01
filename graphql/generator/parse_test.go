package generator_test

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"github.com/graph-gophers/graphql-go"
	"github.com/pdt256/vbratings/graphql/generator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Parse_QueryWithNoParams(t *testing.T) {
	// Given
	code := `func (d *Domain) Query() bool { return true }`
	node := getNodeForFunction(t, code)

	// When
	domainRoot := generator.ParseDomain(node)

	// Then
	domain := domainRoot.Domains[0]
	query := domain.Queries[0]
	assert.Equal(t, "Query", query.Name)
	assert.Equal(t, 0, len(query.Params))
	assert.Equal(t, "bool", query.ReturnTypes[0].Type)
}

func Test_Parse_QueryWithDoc(t *testing.T) {
	// Given
	code := `// Line 1
// Line 2
func (d *Domain) Query() bool { return true }`
	node := getNodeForFunction(t, code)

	// When
	domainRoot := generator.ParseDomain(node)

	// Then
	domain := domainRoot.Domains[0]
	query := domain.Queries[0]
	assert.Equal(t, "// Line 1", query.Docs[0])
	assert.Equal(t, "// Line 2", query.Docs[1])
}

func Test_Parse_QueryWithParams(t *testing.T) {
	// Given
	code := `func (d *Domain) Query(one int, two string, three bool) bool { return true }`
	node := getNodeForFunction(t, code)

	// When
	domainRoot := generator.ParseDomain(node)

	// Then
	domain := domainRoot.Domains[0]
	query := domain.Queries[0]
	assert.Equal(t, "one", query.Params[0].Name)
	assert.Equal(t, "int", query.Params[0].Type)
	assert.Equal(t, "two", query.Params[1].Name)
	assert.Equal(t, "string", query.Params[1].Type)
	assert.Equal(t, "three", query.Params[2].Name)
	assert.Equal(t, "bool", query.Params[2].Type)
}

func Test_Parse_QueryWithArrayIdentReturn(t *testing.T) {
	// Given
	code := `func (d *Domain) Query() []bool { return []bool{true} }`
	node := getNodeForFunction(t, code)

	// When
	domainRoot := generator.ParseDomain(node)

	// Then
	domain := domainRoot.Domains[0]
	query := domain.Queries[0]
	assert.Equal(t, "[]bool", query.ReturnTypes[0].Type)
}

func Test_Parse_QueryWithStructReturn(t *testing.T) {
	// Given
	code := `func (d *Domain) Query() SimpleStruct { return SimpleStruct{} }`
	node := getNodeForFunction(t, code)

	// When
	domainRoot := generator.ParseDomain(node)

	// Then
	domain := domainRoot.Domains[0]
	query := domain.Queries[0]
	assert.Equal(t, "", query.ReturnTypes[0].Package)
	assert.Equal(t, "SimpleStruct", query.ReturnTypes[0].Type)
}

func Test_Parse_QueryWithPackageStructReturn(t *testing.T) {
	// Given
	code := `func (d *Domain) Query() testing.discard { return testing.discard{} }`
	node := getNodeForFunction(t, code)

	// When
	domainRoot := generator.ParseDomain(node)

	// Then
	domain := domainRoot.Domains[0]
	query := domain.Queries[0]
	assert.Equal(t, "testing", query.ReturnTypes[0].Package)
	assert.Equal(t, "discard", query.ReturnTypes[0].Type)
}

func Test_Parse_QueryWithArrayStructReturn(t *testing.T) {
	// Given
	code := `func (d *Domain) Query() SimpleStruct { return SimpleStruct{} }`
	node := getNodeForFunction(t, code)

	// When
	domainRoot := generator.ParseDomain(node)

	// Then
	domain := domainRoot.Domains[0]
	query := domain.Queries[0]
	assert.Equal(t, "", query.ReturnTypes[0].Package)
	assert.Equal(t, "SimpleStruct", query.ReturnTypes[0].Type)
}

func Test_Parse_QueryWithArrayPackageStructReturn(t *testing.T) {
	// Given
	code := `func (d *Domain) Query() []testing.discard { return []testing.discard{} }`
	node := getNodeForFunction(t, code)

	// When
	domainRoot := generator.ParseDomain(node)

	// Then
	domain := domainRoot.Domains[0]
	query := domain.Queries[0]
	assert.Equal(t, "testing", query.ReturnTypes[0].Package)
	assert.Equal(t, "[]discard", query.ReturnTypes[0].Type)
}

func Test_Parse_CommandWithNoReturn(t *testing.T) {
	// Given
	code := `func (d *Domain) Command() {}`
	node := getNodeForFunction(t, code)

	// When
	domainRoot := generator.ParseDomain(node)

	// Then
	domain := domainRoot.Domains[0]
	command := domain.Commands[0]
	assert.Equal(t, "Command", command.Name)
	assert.Equal(t, 0, len(command.Params))
	assert.Equal(t, 0, len(command.ReturnTypes))
}

func Test_Parse_CommandReturnsError(t *testing.T) {
	// Given
	code := `func (d *Domain) Command() error { return nil }`
	node := getNodeForFunction(t, code)

	// When
	domainRoot := generator.ParseDomain(node)

	// Then
	domain := domainRoot.Domains[0]
	command := domain.Commands[0]
	assert.Equal(t, "Command", command.Name)
	assert.Equal(t, 0, len(command.Params))
	assert.Equal(t, "error", command.ReturnTypes[0].Type)
}

func Test_Parse_CommandWithDoc(t *testing.T) {
	// Given
	code := `// Line 1
// Line 2
func (d *Domain) Command() {}`
	node := getNodeForFunction(t, code)

	// When
	domainRoot := generator.ParseDomain(node)

	// Then
	domain := domainRoot.Domains[0]
	command := domain.Commands[0]
	assert.Equal(t, "// Line 1", command.Docs[0])
	assert.Equal(t, "// Line 2", command.Docs[1])
}

func Test_Parse_CommandWithParams(t *testing.T) {
	// Given
	code := `func (d *Domain) Command(one int, two string, three bool) {}`
	node := getNodeForFunction(t, code)

	// When
	domainRoot := generator.ParseDomain(node)

	// Then
	domain := domainRoot.Domains[0]
	command := domain.Commands[0]
	assert.Equal(t, "one", command.Params[0].Name)
	assert.Equal(t, "int", command.Params[0].Type)
	assert.Equal(t, "two", command.Params[1].Name)
	assert.Equal(t, "string", command.Params[1].Type)
	assert.Equal(t, "three", command.Params[2].Name)
	assert.Equal(t, "bool", command.Params[2].Type)
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

func getNodeForFunction(t *testing.T, function string) *ast.File {
	code := `package test
type Domain struct{}
` + function

	fs := token.NewFileSet()
	node, err := parser.ParseFile(fs, "", code, parser.ParseComments)
	require.NoError(t, err)
	return node
}
