package main

import (
	"bufio"
	"go/ast"
	"go/token"
	"log"
	"os"
	"strings"
)
type visitor struct {
	LLOC int
	numOfComments int
	numOfEmptyStatements int
	fset *token.FileSet
}

type Method struct {
	PkgName           string
	StructName        string
	FuncName          string
	Complexity        int
	Comments		  int
	BlankLines		  int
	LLOC			  int
	Receiver          variable
	Parameters        []variable
	Selectors         []Selector
	SelfVarAccessed   []Selector
	OthersVarAccessed []Selector
	Pos               token.Position
}

func findMethodsfromFile(fset *token.FileSet, node *ast.File, fileName string)([]Method) {

	var methods [] Method
	/*fset := token.NewFileSet();
	node, err := parser.ParseFile(fset, fileName, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}*/
	for _, decl := range node.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}

		structName, funcName := funcName(fn)
		var params []variable
		for _, l := range fn.Type.Params.List {
			if l.Names == nil {
				continue
			}
			temp := variable{
				name:    l.Names[0].Name,
				varType: recvString(l.Type),

			}
			params = append(params, temp)
		}
		func_visitor := visitor{}
		//selectors :=Selector_Visitor{}
		varAll :=findSelectorsFromMethod(fn)
		for i, n := range varAll.selectors {
			varAll.selectors[i].line = findLine(fileName, fset.Position(n.pos).Line)
		}
		for i:= range varAll.selectors {
			println(varAll.selectors[i].line)
		}
		switch node_type := decl.(type) {
		case *ast.FuncDecl:
			//println("Function Name: ",node_type.Name.Name)
			/*for _,r := range node_type.Body.List{
				switch t := r.(type) {

				case *ast.AssignStmt:
					nloc+= (fset.Position(t.End()).Line)- (fset.Position(t.Pos()).Line)+1

				case *ast.BlockStmt:
					nloc+= (fset.Position(t.End()).Line)- (fset.Position(t.Pos()).Line)+1

				case *ast.BranchStmt:
					nloc+= (fset.Position(t.End()).Line)- (fset.Position(t.Pos()).Line)+1

				case *ast.SwitchStmt:
					nloc+= (fset.Position(t.End()).Line)- (fset.Position(t.Pos()).Line)+1

				case *ast.LabeledStmt:
					nloc+= (fset.Position(t.End()).Line)- (fset.Position(t.Pos()).Line)+1

				case *ast.ExprStmt:
					nloc+= (fset.Position(t.End()).Line)- (fset.Position(t.Pos()).Line)+1

				case *ast.SendStmt:
					nloc+= (fset.Position(t.End()).Line)- (fset.Position(t.Pos()).Line)+1

				case *ast.IncDecStmt:
					nloc+= (fset.Position(t.End()).Line)- (fset.Position(t.Pos()).Line)+1

				case *ast.GoStmt:
					nloc+= (fset.Position(t.End()).Line)- (fset.Position(t.Pos()).Line)+1


				case *ast.DeferStmt:
					nloc+= (fset.Position(t.End()).Line)- (fset.Position(t.Pos()).Line)+1

				case *ast.ReturnStmt:
					nloc+= (fset.Position(t.End()).Line)- (fset.Position(t.Pos()).Line)+1

				case *ast.IfStmt:
					nloc+= (fset.Position(t.End()).Line)- (fset.Position(t.Pos()).Line)+1

				case *ast.CaseClause:
					nloc+= (fset.Position(t.End()).Line)- (fset.Position(t.Pos()).Line)+1



				case *ast.TypeSwitchStmt:
					nloc+= (fset.Position(t.End()).Line)- (fset.Position(t.Pos()).Line)+1

				case *ast.CommClause:
					nloc+= (fset.Position(t.End()).Line)- (fset.Position(t.Pos()).Line)+1

				case *ast.SelectStmt:
					nloc+= (fset.Position(t.End()).Line)- (fset.Position(t.Pos()).Line)+1

				case *ast.ForStmt:
					nloc+= (fset.Position(t.End()).Line)- (fset.Position(t.Pos()).Line)+1

				case *ast.RangeStmt:
					nloc+= (fset.Position(t.End()).Line)- (fset.Position(t.Pos()).Line)+1

				}
			}*/


          func_visitor.numOfComments = countComments(node,node_type)

			//func_visitor.numOfEmptyStatements = len(emptystatements)
            line,err := countEmptyStatements(fileName,fset.Position(fn.Body.Pos()).Line,fset.Position(fn.Body.End()).Line)
            if err !=nil{
            	log.Fatal(err)
			}
			func_visitor.numOfEmptyStatements = line
			func_visitor.LLOC =(fset.Position(fn.End()).Line)-(fset.Position(fn.Pos()).Line)-1-func_visitor.numOfEmptyStatements-func_visitor.numOfComments

		}
		method := Method{
			PkgName:   node.Name.Name,
			StructName: structName,
			FuncName:   funcName,
			//Receiver:   rcv,
			Parameters: params,
			//Selectors:  varAll.selectors,
			Complexity: method_complexity(fn),
			//LLOC: logical_Line_of_Code(fn),
			Comments:func_visitor.numOfComments,
			BlankLines:func_visitor.numOfEmptyStatements,
			LLOC :func_visitor.LLOC,
			Pos:  fset.Position(fn.Pos()),
		}
		methods =append(methods,method)
		/*println(methods[0].FuncName)
		println(len(methods[0].Parameters))*/
	}
  return methods
}
func funcName(fn *ast.FuncDecl) (string, string) {
	if fn.Recv != nil {
		if fn.Recv.NumFields() > 0 {
			typ := fn.Recv.List[0].Type

			class := recvOnlyNameString(typ)

			return class, fn.Name.Name
			// return fmt.Sprintf("(%s).%s", recvString(typ), fn.Name)
		}
	}
	return "", fn.Name.Name
}
func countComments(file *ast.File,fn *ast.FuncDecl)(int)  {
	var comments[] *ast.CommentGroup
for _,cmnt:= range file.Comments{

	   if fn.Pos()<=cmnt.Pos() && cmnt.End()<=fn.End(){
	   	comments =append(comments,cmnt)
	   }

	}
   return len(comments)
}

func countEmptyStatements(fileName string,start int,end int )(int,error)  {
	emptylines:=0
	//println(start," ",end)
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	for  l:=0; l<start;l++{
		scanner.Scan()
	}
	for t:=start+1;t<=end-1;t++{

		scanner.Scan()
		line :=scanner.Text()
		line = strings.Trim(line," ")
		line = strings.TrimSpace(line)
		//count blanklines annd only bracket
		if line=="" || line=="{" || line=="}"{
			emptylines++;
		}

	}
	return emptylines,nil

}


