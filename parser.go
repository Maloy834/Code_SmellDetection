package main

import (
	"go/parser"
	"go/token"
	"log"
)
func (st *Struct) addMethod(method Method) {
	st.Methods = append(st.Methods, method)
}

func parseFile(fileName string)([]Struct,[]Method)  {
	var methods [] Method
	var structs[] Struct
	fset := token.NewFileSet();
	node, err := parser.ParseFile(fset, fileName, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}
	structs = findStructsfromFile(fset, node)
	methods = findMethodsfromFile(fset, node, fileName)

	return structs,methods
}
