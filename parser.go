package main

import (
	"go/parser"
	"go/token"
	"log"
)
func (st *Struct) addUsedMethod(method Method) {
	st.Methods = append(st.Methods, method)
}
func (st *Struct) addAllMethods(method []Method) {
	for _,m := range method {
		st.Totalmethods = append(st.Totalmethods, m)
	}
}

func parseFile(fileName string)([]Struct,[]Method)  {
	var methods [] Method
	var structs[] Struct
	fset := token.NewFileSet();
	node, err := parser.ParseFile(fset, fileName, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}
	structs = findStructsfromFile(fset, node,fileName)
	methods = findMethodsfromFile(fset, node, fileName,structs)

	return structs,methods
}
