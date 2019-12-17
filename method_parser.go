package main

import (
	"bufio"
	"go/ast"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type visitor struct {
	LLOC                 int
	numOfComments        int
	numOfEmptyStatements int
	fset                 *token.FileSet
}

type Method struct {
	PkgName           		string
	StructName        		string
	FuncName          		string
	FileName                string
	Complexity        		int
	Comments         		int
	BlankLines        		int
	LLOC              		int
	MAXNESTING				int
	ATFD			  		int
	RatioLOCAndComplexity 	float64
	FDP				  		int
	NOAV			  		int
	LAA				  		float64
	funcCalling				[]method_visitor
	CM                      int
	CC						int
	isExported				bool
	FeatureEnvy				bool
	BrainMethod             bool
	LongParameter           bool
	ShortGunSurgery         bool
	Receiver          		variable
	Parameters        		[]variable
	Selectors         		[]Selector
	SelfVarAccessed   		[]Selector
	OthersVarAccessed 		[]Selector
	Pos               		token.Position
}

func findMethodsfromFile(fset *token.FileSet, node *ast.File, fileName string,structs [] Struct) []Method {

	var methods []Method
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
		/*if fn.Recv == nil || fn.Recv.List[0].Names == nil {
			continue
		}*/

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
		receivers := variable{}
		if fn.Recv == nil || fn.Recv.List[0].Names == nil {
			receivers.name = ""
			receivers.varType = ""
			/*receivers := variable{
				name:    fn.Recv.List[0].Names[0].Name,
				varType: recvString(fn.Recv.List[0].Type),
			}*/
		} else {

			receivers.name = fn.Recv.List[0].Names[0].Name
			receivers.varType = recvString(fn.Recv.List[0].Type)
		}

		func_visitor := visitor{}
		//selectors :=Selector_Visitor{}
		//ast.Walk(&selectors, fn.Body)
		varAll := findSelectorsFromMethod(fn)
		varType := findVariableDeclaration(fn)
		funcCallingAll:=checkMethodCalling(fn.Body.List)
		for i,n:=range varType.vardeclare{

			varType.vardeclare[i].var_type=findVarType(fileName,fset.Position(n.pos).Line,fset.Position(n.pos).Line,varType.vardeclare[i].name)
			if structName==""{
				structName=findMethodStruct(structs,varType.vardeclare[i].var_type)
			}
		}
		for i, n := range varAll.selectors {
			varAll.selectors[i].line = findLine(fileName, fset.Position(n.pos).Line)
		}
		/*for i := range varAll.selectors {
			println("----------Selecotrs")
			println(funcName)
			println("Selecotrs left: " + varAll.selectors[i].left + " Selecotrs right: " + varAll.selectors[i].right)
		}*/
		switch node_type := decl.(type) {
		case *ast.FuncDecl:

			func_visitor.numOfComments = countComments(node, node_type)

			//func_visitor.numOfEmptyStatements = len(emptystatements)
			line, err := countEmptyStatements(fileName, fset.Position(fn.Body.Pos()).Line, fset.Position(fn.Body.End()).Line)
			if err != nil {
				log.Fatal(err)
			}
			func_visitor.numOfEmptyStatements = line
			func_visitor.LLOC = (fset.Position(fn.End()).Line) - (fset.Position(fn.Pos()).Line) - 1-func_visitor.numOfEmptyStatements - func_visitor.numOfComments
			if func_visitor.LLOC<=0{
				func_visitor.LLOC=1
			}

		}
		method := Method{
			PkgName:    node.Name.Name,
			StructName: structName,
			FuncName:   funcName,
			FileName: filepath.Base(fileName),
			Receiver:   receivers,
			funcCalling:funcCallingAll,
			Parameters: params,
			Selectors:  varAll.selectors,
			Complexity: method_complexity(fn),
			MAXNESTING:calculateMethodNesting(node,fn),
			Comments:   func_visitor.numOfComments,
			BlankLines: func_visitor.numOfEmptyStatements,
			LLOC:       func_visitor.LLOC,
			Pos:        fset.Position(fn.Pos()),
		}
		if fn.Name.IsExported() == true {
			method.isExported = true
		} else {
			method.isExported = false
		}
		method.separateAccessedVariable()
		method.claculateMethodATFD()
		method.calculateMethodLAA()
		method.NOAV=len(method.SelfVarAccessed)+len(method.OthersVarAccessed)
		//method.calculateFDP()
		//println(method.Complexity/method.LLOC)
		//var t= method.Complexity
		//var m= method.LLOC
		//var n float64
		//n =float64(t)/float64(m)
		//r:=strconv.FormatFloat(n, 'f', -1, 64)
		/*println(t," ", m)
		println(r)*/
		method.RatioLOCAndComplexity=float64(method.Complexity)/float64(method.LLOC)
		// s:=fmt.Sprintf("%.12f",method.RatioLOCAndComplexity)
		// println(s)
		//io.WriteString(os.Stdout, s)
		methods = append(methods, method)
		/*println(methods[0].FuncName)
		println(len(methods[0].Parameters))*/
	}
	return methods
}
func (m *Method) separateAccessedVariable() {
	for _, s := range m.Selectors {
		//println(m.FuncName, " ",s.left,"  ",s.exported)

		if !isVariable(s.line, s.left, s.right) {
			continue
		}

		if s.left == m.Receiver.name {
			m.SelfVarAccessed = append(m.SelfVarAccessed, s)
		} else {
			m.OthersVarAccessed = append(m.OthersVarAccessed, s)
		}
	}
}
func funcName(fn *ast.FuncDecl) (string, string) {
	if fn.Recv != nil {
		if fn.Recv.NumFields() > 0 {
			typ := fn.Recv.List[0].Type

			class := recvOnlyNameString(typ)
			//println("structName: ",class)
			return class, fn.Name.Name
			// return fmt.Sprintf("(%s).%s", recvString(typ), fn.Name)
		}
	}
	return " ", fn.Name.Name
}
func countComments(file *ast.File, fn *ast.FuncDecl) int {
	var comments []*ast.CommentGroup
	for _, cmnt := range file.Comments {

		if fn.Pos() <= cmnt.Pos() && cmnt.End() <= fn.End() {
			comments = append(comments, cmnt)
		}

	}
	//println(fr)
	return len(comments)
}

func countEmptyStatements(fileName string, start int, end int) (int, error) {
	emptylines := 0
	//println(start," ",end)
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	for l := 0; l < start; l++ {
		scanner.Scan()
	}
	for t := start + 1; t <= end-1; t++ {

		scanner.Scan()
		line := scanner.Text()
		line = strings.Trim(line, " ")
		line = strings.TrimSpace(line)
		//count blanklines annd only bracket
		if line == "" || line == "{" || line == "}" {
			emptylines++
		}

	}
	return emptylines, nil

}
func findVarType(fileName string, start int, end int,value string) string {
	varType := ""
	//println(start," ",end)
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	for l := 0; l < start-1; l++ {
		scanner.Scan()
	}
	for t := start; t <= end; t++ {
		scanner.Scan()
		line := scanner.Text()
		line = strings.Trim(line, " ")
		line = strings.TrimSpace(line)
		//count blanklines annd only bracket
		indx:= strings.Index(line,value)
		varType= line[indx+1:]
		varType = strings.TrimSpace(varType)



	}

	return varType

}
