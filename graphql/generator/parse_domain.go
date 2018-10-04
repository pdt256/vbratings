package generator

import (
	"fmt"
	"go/ast"
)

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
	Package string
	Type    string
}

func ParseDomain(node ast.Node) *DomainRoot {
	visitor := DomainVisitor{}
	ast.Walk(&visitor, node)

	return &visitor.root
}

type DomainVisitor struct {
	root          DomainRoot
	currentDomain *Domain
}

func (v *DomainVisitor) Visit(node ast.Node) ast.Visitor {
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
				case *ast.SelectorExpr:
					currentUseCase.ReturnTypes = append(
						currentUseCase.ReturnTypes,
						ReturnType{
							Package: r.X.(*ast.Ident).Name,
							Type:    r.Sel.Name,
						},
					)
				case *ast.ArrayType:
					switch s := r.Elt.(type) {
					case *ast.SelectorExpr:
						currentUseCase.ReturnTypes = append(
							currentUseCase.ReturnTypes,
							ReturnType{
								Package: s.X.(*ast.Ident).Name,
								Type:    fmt.Sprintf("[%s]", s.Sel.Name),
							},
						)
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
