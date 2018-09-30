package generator_test

import (
	"go/parser"
	"go/token"
	"testing"

	"github.com/pdt256/vbratings/graphql/generator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Parse_SingleDomainSingleQuery(t *testing.T) {
	// Given
	filePath := `./testdata/single_domain_single_query.go`
	fs := token.NewFileSet()
	node, err := parser.ParseFile(fs, filePath, nil, parser.ParseComments)
	require.NoError(t, err)

	// When
	domainRoot := generator.ParseDomain(node)

	// Then
	domain := domainRoot.Domains[0]
	firstQuery := domain.Queries[0]
	assert.Equal(t, "SingleDomainSingleQuery", domain.Name)
	assert.Equal(t, "GetQuery", firstQuery.Name)
	assert.Equal(t, "// Line 1", firstQuery.Docs[0])
	assert.Equal(t, "// Line 2", firstQuery.Docs[1])
	assert.Equal(t, "oneInt", firstQuery.Params[0].Name)
	assert.Equal(t, "int", firstQuery.Params[0].Type)
	assert.Equal(t, "twoString", firstQuery.Params[1].Name)
	assert.Equal(t, "string", firstQuery.Params[1].Type)
	assert.Equal(t, "threeBool", firstQuery.Params[2].Name)
	assert.Equal(t, "bool", firstQuery.Params[2].Type)
	assert.Equal(t, "bool", firstQuery.ReturnTypes[0].Type)
}
