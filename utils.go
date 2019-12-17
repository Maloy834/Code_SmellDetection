package main

import (
	"bufio"
	"go/ast"
	"go/token"
	"os"
	"strings"
)

const TCC_Null = 99999

type variable struct {
	name    string
	varType string
	count   int
}
type method_visitor struct{
	methodName string

}


func isDir(filename string) bool {
	fi, err := os.Stat(filename)
	return err == nil && fi.IsDir()
}

func recvString(recv ast.Expr) string {
	switch t := recv.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.StarExpr:
		return "*" + recvString(t.X)
	}
	return "BADRECV"
}

func recvOnlyNameString(recv ast.Expr) string {
	switch t := recv.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.StarExpr:
		return recvOnlyNameString(t.X)
	}
	return "BADRECV"
}

func isAlphaNumeric(c byte) bool {
	s := "123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := []byte(s)

	for _, x := range b {
		if x == c {
			return true
		}
	}

	return false
}

func findLine(file string, pos int) string {
	f, err := os.Open(file)
	defer f.Close()

	if err != nil {
		return "ERROR"
	}

	bf := bufio.NewReader(f)
	var line string
	for lnum := 0; lnum < pos; lnum++ {
		line, err = bf.ReadString('\n')
		if err != nil {
			return "ERROR"
		}
	}

	return strings.TrimSpace(line)
}

func isVariable(line string, leftVar string, rightVar string) bool {
	for len(line) > len(leftVar)+len(rightVar) {
		pos := strings.Index(line, leftVar)

		if pos == -1 {
			break
		}

		if line[pos+len(leftVar)] == '.' {
			line2 := line[pos+len(leftVar)+1:]

			if pos2 := strings.Index(line2, rightVar); pos2 == 0 {
				if len(line2) == len(rightVar) {
					return true
				}

				line3 := line2[len(rightVar):]

				if line3[0] == '(' {
					return false
				} else if !isAlphaNumeric(line3[0]) {
					return true
				} else {
					return true
				}
			}
		}

		/*if pos = strings.Index(line, " "); pos == -1 {
			break
		}*/

		line = line[pos+1:]
	}

	return false
}

type uniqeSelectors struct {
	selectors []Selector
}
type uniqueFileName struct{
	fileName [] string
}
type uniqueStruct struct {
	structs [] Struct
	//stName [] string
}
type VariableDeclare struct {
	var_type string
	pos token.Pos
	name  string

}
type variable_visitor struct {
	vardeclare [] VariableDeclare
}
func findVariableDeclaration(fn *ast.FuncDecl) variable_visitor{
	node_visitor := variable_visitor{}
	ast.Walk(&node_visitor,fn.Body)
	return node_visitor
}
func (node_visitor *variable_visitor) Visit(node ast.Node) ast.Visitor {
	if node == nil {
		return node_visitor
	}
	switch node:=node.(type) {
	case *ast.GenDecl:
		for _, spec := range node.Specs {
			switch spec := spec.(type) {

			case *ast.ValueSpec:
				for _, id := range spec.Names {
					newVariable:= VariableDeclare{
						name:id.Name,
						pos:id.Pos(),
					}
					node_visitor.addVariableDeclare(newVariable)
				}

			}
		}

	}
	return node_visitor
}
func (sv * variable_visitor)addVariableDeclare(s VariableDeclare)  {
	sv.vardeclare=append(sv.vardeclare,s)
}
 func containsVariable(value string, attributes[] variable) bool{
    for _,var_name:= range attributes{
    	if var_name.name==value{
    		return true
		}
	}
 	return false
 }
 func (u *uniqueStruct) existStruct(st Struct) bool{
 	for _,v := range u.structs{
 		if v.StructName== st.StructName{
			return true
		}
	}
	return false
 }
func (u *uniqeSelectors) exists(s Selector) bool {
	for _, v := range u.selectors {
		// println("WKAALAKLKA", v.toString(), s.toString)
		if v.toString() == s.toString() {
			return true
		}
	}
	return false
}
func (u *uniqueStruct) addStruct(st Struct){
	u.structs=append(u.structs,st)
}
func (u *uniqueFileName)addFileName(name string){
	u.fileName=append(u.fileName,name)
}
func (u *uniqueFileName) existFilename(name string) bool{
	for _,v := range u.fileName{
		if v == name{
			return true
		}
	}
	return false
}

func (u *uniqeSelectors) add(s Selector) {
	u.selectors = append(u.selectors, s)
}
func checkMethodCalling(stmts []ast.Stmt) [] method_visitor{
	var mVisitor [] method_visitor
	for _, stmt := range stmts {
		if exprStmt, ok := stmt.(*ast.ExprStmt); ok {
			if call, ok := exprStmt.X.(*ast.CallExpr); ok {
				if fun, ok := call.Fun.(*ast.SelectorExpr); ok {
					//println( fun.Sel.Name)
					newMethodVisitor:= method_visitor{
						methodName:   fun.Sel.Name,

					}
					mVisitor=append(mVisitor,newMethodVisitor)

				}else{
					if tr, ok := call.Fun.(*ast.Ident); ok {
						new_MethodVisitor:= method_visitor{
							methodName:   tr.Name,

						}
						mVisitor=append(mVisitor,new_MethodVisitor)
					}
				}
			}
		}
	}
	return mVisitor
}
func findMethodStruct(structs[] Struct,value string) string {
	for _,st:= range structs{
		if value==st.StructName{
			return st.StructName
		}

	}
	return ""
}
func (m * Method)calculateCallerMethod(methods [] Method)  {
	cm:=0
	//ustruct:= uniqueStruct{}
	ustruct:=uniqueFileName{}
	for _,mt:= range methods{
		for _,callerMt:= range mt.funcCalling{
			//println(callerMt.methodName,"   trr ",m.FuncName)
			if callerMt.methodName== m.FuncName{
				cm++
				/*if m.StructName!="" && m.StructName!= mt.StructName && !ustruct.existStructame(m.StructName){
					ustruct.addStructName(m.StructName)
				}*/
				if m.FileName!= mt.FileName && !ustruct.existFilename(m.FileName){
					//println(m.FileName,"  debug  ", mt.FileName)
					ustruct.addFileName(m.FileName)
				}
			}
		}
	}
	cc:= len(ustruct.fileName)
	m.CM=cm
	m.CC=cc
}
