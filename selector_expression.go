package main

import (
	"go/ast"
	"go/token"
)

type Selector struct {
	left  string
	right string
	exported bool
	st_name interface{}
	pos   token.Pos
	line  string
}

func (s *Selector) toString() string {
	return s.left + "." + s.right

}

type Selector_Visitor struct {
	selectors []Selector
}

func findSelectorsFromMethod(fn *ast.FuncDecl) Selector_Visitor {
	variables := Selector_Visitor{}
	ast.Walk(&variables, fn.Body)
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
					exported:va.IsExported(),
					//st_name:selectorExp.Sel.Obj.Name,
					pos:   va.Pos(),
				}
				//println(newSelector.st_name)
				v.add(newSelector)
			}
		}
	}
	return v
}

func (v *Selector_Visitor) add(s Selector) {
	//println("Before Check: "+s.left+" "+"  "+s.right)
	if !v.exists(s) {
		//println("After Check: "+s.left+" "+"  "+s.right)
		v.selectors = append(v.selectors, s)
		//println(len(v.selectors))
	}

}

func (v *Selector_Visitor) exists(s Selector) bool {
	//println(len(v.selectors))
	for _, n := range v.selectors {
		//println(len(v.selectors))
		/*println(n.left+" "+s.left+" "+n.right+"  "+s.right+ "")
		println(len(n.left))
		println(len(s.left))
		println(len(n.right))
		println(len(s.right))*/
		if n.left == s.left && n.right == s.right {
			//println("check")
			//println(s.left+" "+"  "+s.right)
			return true
		}
	}
	//println(s.left+" "+"  "+s.right)
	return false
}
