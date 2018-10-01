package generator

import (
	"fmt"
	"go/ast"
)

type EntityRoot struct {
	Entities []*Entity
}

type Entity struct {
	Name   string
	Fields []Field
}

type Field struct {
	Name string
	Type string
}

func ParseEntities(node ast.Node) *EntityRoot {
	visitor := EntityVisitor{}
	ast.Walk(&visitor, node)

	return &visitor.entityRoot
}

type EntityVisitor struct {
	entityRoot EntityRoot
}

func (v *EntityVisitor) Visit(node ast.Node) ast.Visitor {
	//fmt.Printf("%T\n", node)

	switch n := node.(type) {
	case *ast.TypeSpec:
		if s, ok := n.Type.(*ast.StructType); ok {
			currentEntity := &Entity{
				Name: n.Name.Name,
			}
			v.entityRoot.Entities = append(v.entityRoot.Entities, currentEntity)
			for _, field := range s.Fields.List {
				typeName := fmt.Sprintf("%s", field.Type)

				var name string
				if len(field.Names) > 0 {
					name = field.Names[0].Name
				} else {
					name = lowerInitial(typeName)
				}

				currentEntity.Fields = append(
					currentEntity.Fields,
					Field{
						Name: name,
						Type: typeName,
					},
				)
			}
		}
	}

	return v
}
