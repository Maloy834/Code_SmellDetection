package main

import (
	"go/ast"

)

type method_nesting struct {
	MAXNESTING int
	OFFSET int
}


func calculateMethodNesting(file *ast.File,fn *ast.FuncDecl) int {
	nesting_visitor := method_nesting{
		OFFSET:int(file.Pos()),
	}
	ast.Walk(&nesting_visitor, fn)
	//println(fn.Name.Name)
	return nesting_visitor.MAXNESTING
}
func (node_visitor *method_nesting) Visit(node ast.Node) ast.Visitor {
	//println(node_visitor)
	if node == nil {
		return node_visitor
	}
	switch node := node.(type) {
		case *ast.FuncDecl:
			fun := node
			node_visitor.MAXNESTING=calculateNesting(node_visitor.OFFSET,fun)

	}
	return node_visitor
}


