package generator

import (
	"fmt"
	"go/ast"
)

type Root struct {
	Domains []*Domain
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
