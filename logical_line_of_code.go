/*package main

import (
	"go/ast"
	"go/token"
)


type p = token.Pos
type visitor struct {
	LLOC int
	fset *token.FileSet
}

func logical_Line_of_Code( fn *ast.FuncDecl) (int)  {
	func_visitor := visitor{}
	ast.Walk(&func_visitor, fn)

	return func_visitor.LLOC
}
func (v *visitor) Visit(node ast.Node)  ast.Visitor  {

	switch node_type := node.(type) {
	case *ast.FuncDecl:
		println("Function Name: ",node_type.Name.Name)
		for _,r := range node_type.Body.List{
			switch t := r.(type) {

			case *ast.AssignStmt:
				println("assign")
			case *ast.BlockStmt:
				println("block")
			case *ast.BranchStmt:
				println("branch")
			case *ast.SwitchStmt:
				println("switch",)

			case *ast.LabeledStmt:
				println("labeled")

			case *ast.ExprStmt:
				println("Expression")
			case *ast.SendStmt:
				println("send")
			case *ast.IncDecStmt:
				println("Indec")
			case *ast.GoStmt:
				println("Gostmnt")

			case *ast.DeferStmt:
				println("Defer")
			case *ast.ReturnStmt:
				println("Return")


			case *ast.IfStmt:
				println("if")
			case *ast.CaseClause:
				println("case")


			case *ast.TypeSwitchStmt:
				/*func (s *ast.TypeSwitchStmt) End() token.Pos { return s.Body.End() }*/
				/*println("Typeswitch",t.Pos())
			case *ast.CommClause:
				println("Commclause")
			case *ast.SelectStmt:
				println("select")
			case *ast.ForStmt:
				println("For")
			case *ast.RangeStmt:
				println("Range")
			}
		}*/
		 //fn := node.(*ast.FuncDecl)
		 //walkStmtList(*v,fn.Body.List)
		//fn := node.(*ast.FuncDecl)
        //println(fset.Position(fn.Pos()).Line)
		//mutex.Unlock()
		//println("Function name: ",node_type.Name.Name,node_type.Doc.Text()," ",node_type.End())
		//println(len(node_type.Body.List) )
		//println("Function Name: ",node_type.Name.Name, "Body: ",node_type.Body.End())

		 //nloc :=  (v.fset.Position(node_type.End()).Line) - (v.fset.Position(node_type.Pos()).Line + 1)
		  //mutex.Lock()
		//nloc :=v.fset.Position(node_type.End()).Line- v.fset.Position(node_type.Pos()).Line +1



   /* v.LLOC =len(node_type.Body.List)
//v.LLOC = nloc
}

return v
}
*/