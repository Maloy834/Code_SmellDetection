package main

import (
	"go/ast"
	"go/token"
)

type complexity_visitor struct {
	complexity int
}

func method_complexity(fn *ast.FuncDecl) (int) {
	node_visitor :=complexity_visitor{}
	ast.Walk(&node_visitor,fn)
	println(fn.Name.Name)
    return node_visitor.complexity
}
func (node_visitor *complexity_visitor)Visit(node ast.Node) ast.Visitor  {
	//println(node_visitor)
	switch node:=node.(type) {
	case *ast.FuncDecl, *ast.IfStmt, *ast.ForStmt, *ast.RangeStmt, *ast.CaseClause, *ast.CommClause:
		//println(node.())
		node_visitor.complexity++
	case *ast.BinaryExpr:
		if node.Op == token.LAND || node.Op == token.LOR {

			node_visitor.complexity++
		}
	}
	return node_visitor
}