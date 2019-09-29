package main

import (
	"go/ast"
	"go/token"
)

type Struct struct {
	PkgName    string
	StructName string
	Attributes []variable
	Methods    []Method // How many method use this struct
	Pos        token.Position
	WMC        int
	NDC        int
	NP         int
	ATFD       int
	TCC        float64
	GodStruct  bool
	DemiGod    bool
}

func findStructsfromFile(fset *token.FileSet, node *ast.File) []Struct {
	var structs []Struct
	packageName := node.Name.Name
	findStructs := func(node ast.Node) bool {
		st, ok := node.(*ast.TypeSpec)
		if !ok || st.Type == nil {
			return true
		}
		structName := st.Name.Name
		var attributes []variable
		p, ok := st.Type.(*ast.StructType)
		if !ok {
			return true
		}
		for _, field_list := range p.Fields.List {

			if field_list.Names == nil {
				continue
			}
			for i := 0; i < len(field_list.Names); i++ {
				list := variable{
					name:    field_list.Names[i].Name,
					varType: recvString(field_list.Type),
				}
				attributes = append(attributes, list)
			}

		}
		c := Struct{
			PkgName:    packageName,
			StructName: structName,
			Attributes: attributes,
			Pos:        fset.Position(st.Pos()),
		}

		structs = append(structs, c)
		return true

	}

	ast.Inspect(node, findStructs)

	return structs
}
