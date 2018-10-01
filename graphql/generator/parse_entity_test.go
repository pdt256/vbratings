package generator_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"github.com/pdt256/vbratings/graphql/generator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ParseEntities_SimpleEntity(t *testing.T) {
	// Given
	code := `type Entity struct { one int; two string; three bool }`
	node := getNodeForStruct(t, code)

	// When
	entityRoot := generator.ParseEntities(node)

	// Then
	entity := entityRoot.Entities[0]
	assert.Equal(t, "Entity", entity.Name)
	assert.Equal(t, "one", entity.Fields[0].Name)
	assert.Equal(t, "int", entity.Fields[0].Type)
	assert.Equal(t, "two", entity.Fields[1].Name)
	assert.Equal(t, "string", entity.Fields[1].Type)
	assert.Equal(t, "three", entity.Fields[2].Name)
	assert.Equal(t, "bool", entity.Fields[2].Type)
}

func getNodeForStruct(t *testing.T, structCode string) *ast.File {
	code := `package test
` + structCode

	fs := token.NewFileSet()
	node, err := parser.ParseFile(fs, "", code, parser.ParseComments)
	require.NoError(t, err)
	return node
}
