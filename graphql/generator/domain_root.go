package generator

import (
	"fmt"
	"io"
	"strings"
	"unicode"
)

type DomainRoot struct {
	Domains  []*Domain
	Entities []*Entity
}

func (r DomainRoot) GraphQLSchema(schema io.Writer) {
	fmt.Fprintf(schema, "schema {\n")

	if r.hasQueries() {
		fmt.Fprintf(schema, "\tquery: Query\n")
	}

	if r.hasCommands() {
		fmt.Fprintf(schema, "\tmutation: Mutation\n")
	}

	fmt.Fprintf(schema, "}\n")

	if r.hasQueries() {
		fmt.Fprintf(schema, "\ntype Query {\n")

		for _, domain := range r.Domains {
			fmt.Fprintf(schema, "\t%sQueries: %sQueries\n", lowerInitial(domain.Name), domain.Name)
		}
		fmt.Fprintf(schema, "}\n")

		for _, domain := range r.Domains {
			fmt.Fprintf(schema, "\ntype %sQueries {\n", domain.Name)
			for _, query := range domain.Queries {
				addDocs(query, schema)
				fmt.Fprintf(schema, "\t%s", lowerInitial(query.Name))

				addParams(query, schema)

				fmt.Fprintf(schema, ": %s\n", getGraphQLType(query.ReturnTypes[0].Type))
			}
			fmt.Fprintf(schema, "}\n")
		}
	}

	if r.hasCommands() {
		fmt.Fprintf(schema, "\ntype Mutation {\n")

		for _, domain := range r.Domains {
			fmt.Fprintf(schema, "\t%sCommands: %sCommands\n", lowerInitial(domain.Name), domain.Name)
		}
		fmt.Fprintf(schema, "}\n")

		for _, domain := range r.Domains {
			fmt.Fprintf(schema, "\ntype %sCommands {\n", domain.Name)
			for _, command := range domain.Commands {
				addDocs(command, schema)
				fmt.Fprintf(schema, "\t%s", lowerInitial(command.Name))

				addParams(command, schema)

				fmt.Fprintf(schema, ": %s\n", "Boolean!")

			}
			fmt.Fprintf(schema, "}\n")
		}
	}

	for _, entity := range r.Entities {
		fmt.Fprintf(schema, "\ntype %s {", entity.Name)

		for _, field := range entity.Fields {
			fmt.Fprintf(schema, "\n\t%s: %s", field.Name, getGraphQLType(field.Type))
		}

		fmt.Fprintf(schema, "\n}\n")
	}
}

func addDocs(useCase *UseCase, schema io.Writer) {
	for _, doc := range useCase.Docs {
		fmt.Fprintf(schema, "\t# %s\n", strings.TrimPrefix(doc, "// "))
	}
}

func addParams(useCase *UseCase, schema io.Writer) {
	if len(useCase.Params) > 0 {
		fmt.Fprintf(schema, "(")
		for _, param := range useCase.Params {
			fmt.Fprintf(schema, "\n\t\t%s: %s", param.Name, getGraphQLType(param.Type))
		}
		fmt.Fprintf(schema, "\n\t)")
	}
}

func getGraphQLType(s string) string {
	switch s {
	case "bool":
		return "Boolean!"
	case "int":
		return "Int!"
	case "string":
		return "String!"
	default:
		return s + "!"
	}
}

func lowerInitial(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

func (r DomainRoot) hasQueries() bool {
	return true
}

func (r DomainRoot) hasCommands() bool {
	return true
}
