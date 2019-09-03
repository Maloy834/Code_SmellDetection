package main

import (
	"go/ast"
	"go/token"
)

type Selector struct {
	left  string
	right string
	pos   token.Pos
	line  string
}

func (s *Selector) toString() string {
	return s.left + "." + s.right
}

type Selector_Visitor struct {
	selectors []Selector
}

func findSelectorsFromMethod(fn *ast.FuncDecl)(Selector_Visitor)  {
 variables :=Selector_Visitor{}
 ast.Walk(&variables,fn.Body)
 return variables
}
func (v *Selector_Visitor) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return v
	}

	if selectorExp, ok := n.(*ast.SelectorExpr); ok {
		if va, ok := selectorExp.X.(*ast.Ident); ok {
			if va.Obj == nil {
				return v
			}

			if va.Obj.Kind.String() == "var" {

				newSelector := Selector{
					left:  va.Name,
					right: selectorExp.Sel.Name,
					pos:   va.Pos(),
				}
				v.add(newSelector)
			}
		}
	}
	return v
}

func (v *Selector_Visitor) add(s Selector) {
	if !v.exists(s) {
		v.selectors = append(v.selectors, s)
	}
}

func (v *Selector_Visitor) exists(s Selector) bool {
	for _, n := range v.selectors {
		if n.left == s.left && n.right == s.right {
			return true
		}
	}
	return false
}
