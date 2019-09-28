package main

import (
	"go/ast"
	"go/token"
)

type complexityVisitor struct {
	// Complexity is the cyclomatic complexity
	Complexity int
}

func complexity(fn *ast.FuncDecl) int {
	v := complexityVisitor{}
	ast.Walk(&v, fn)
	return v.Complexity
}

func (v *complexityVisitor) Visit(n ast.Node) ast.Visitor {

	switch n := n.(type) {
	case *ast.FuncDecl, *ast.IfStmt, *ast.ForStmt, *ast.RangeStmt, *ast.CaseClause, *ast.CommClause:
		v.Complexity++
		// its a check
	case *ast.BinaryExpr:
		if n.Op == token.LAND || n.Op == token.LOR {
			v.Complexity++
			// its a check

		}
	}
	return v
}
