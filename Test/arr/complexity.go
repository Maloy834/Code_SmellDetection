package main

import (
	"go/ast"
	"go/token"
)

type complexityVisitor struct {
	// Complexity is the cyclomatic complexity
	Complexity int
	Count int
}

func complexity(fn *ast.FuncDecl) int {
	v := complexityVisitor{}
	for i:=0;i<10;i++{
		print(i)
	}
	ast.Walk(&v, fn)
	return v.Complexity
}

func (v *complexityVisitor) Visit(n ast.Node) ast.Visitor {
	s := sumation{}
	for i:=0;i<10;i++{
		s.sum(i,i+1,i+2)
	}
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
