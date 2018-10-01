package generator

import (
	"fmt"
	"go/ast"
	"io"
	"strings"
	"unicode"
)

type Root struct {
	Domains []*Domain
}

func (r Root) GraphQLSchema(schema io.Writer) {
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
	}

	return s
}

func lowerInitial(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

func (r Root) hasQueries() bool {
	return true
}

func (r Root) hasCommands() bool {
	return true
}

type Domain struct {
	Name     string
	Commands []*UseCase
	Queries  []*UseCase
}

type UseCase struct {
	Name        string
	Docs        []string
	Params      []Param
	ReturnTypes []ReturnType
}

type Param struct {
	Name string
	Type string
}

type ReturnType struct {
	Type string
}

func ParseDomain(node *ast.File) *Root {
	visitor := Visitor{}
	ast.Walk(&visitor, node)

	return &visitor.root
}

type Visitor struct {
	root          Root
	currentDomain *Domain
}

func (v *Visitor) Visit(node ast.Node) (w ast.Visitor) {
	//fmt.Printf("%T\n", node)

	switch n := node.(type) {
	case *ast.TypeSpec:
		v.currentDomain = &Domain{Name: n.Name.Name}
		v.root.Domains = append(v.root.Domains, v.currentDomain)
	case *ast.FuncDecl:
		useCaseName := n.Name.String()
		currentUseCase := &UseCase{
			Name: useCaseName,
		}

		if isCommand(n) {
			v.currentDomain.Commands = append(
				v.currentDomain.Commands,
				currentUseCase,
			)
		} else {
			v.currentDomain.Queries = append(
				v.currentDomain.Queries,
				currentUseCase,
			)
		}

		if n.Doc != nil {
			for _, comment := range n.Doc.List {
				currentUseCase.Docs = append(currentUseCase.Docs, comment.Text)
			}
		}

		if n.Type.Params.List != nil {
			for _, param := range n.Type.Params.List {
				currentUseCase.Params = append(
					currentUseCase.Params,
					Param{
						Name: param.Names[0].String(),
						Type: fmt.Sprintf("%s", param.Type),
					})
			}
		}

		if n.Type.Results != nil {
			for _, result := range n.Type.Results.List {
				switch r := result.Type.(type) {
				case *ast.ArrayType:
					switch s := r.Elt.(type) {
					//case *ast.SelectorExpr:
					//	fmt.Printf("      []%s\n", s.Sel.Name)
					case *ast.Ident:
						currentUseCase.ReturnTypes = append(
							currentUseCase.ReturnTypes,
							ReturnType{
								Type: fmt.Sprintf("[]%s", s.Name),
							},
						)
					}
				case *ast.Ident:
					currentUseCase.ReturnTypes = append(
						currentUseCase.ReturnTypes,
						ReturnType{
							Type: fmt.Sprintf("%s", r.Name),
						},
					)
				}
			}
		}
	}

	return v
}

func isCommand(n *ast.FuncDecl) bool {
	if n.Type.Results == nil {
		return true
	}

	if len(n.Type.Results.List) == 1 {
		if v, ok := n.Type.Results.List[0].Type.(*ast.Ident); ok {
			if v.Name == "error" {
				return true
			}
		}
	}

	return false
}
