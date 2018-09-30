package generator

import (
	"fmt"
	"go/ast"
)

type Root struct {
	Domains []*Domain
}

type Domain struct {
	Name string
	//Commands []*UseCase
	Queries []*UseCase
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
		v.currentDomain.Queries = append(
			v.currentDomain.Queries,
			currentUseCase,
		)

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
				//case *ast.ArrayType:
				//	switch s := r.Elt.(type) {
				//	case *ast.SelectorExpr:
				//		fmt.Printf("      []%s\n", s.Sel.Name)
				//	case *ast.Ident:
				//		fmt.Printf("      []%s\n", s.Name)
				//	default:
				//		fmt.Printf(" UNKNOWN %#v\n", s)
				//	}
				case *ast.Ident:
					currentUseCase.ReturnTypes = append(
						currentUseCase.ReturnTypes,
						ReturnType{
							Type: fmt.Sprintf("%s", r.Name),
						},
					)
					//default:
					//	fmt.Printf(" UNKNOWN %#v\n", r)
				}

			}
		}
	}

	return v
}
