package main
import "path/filepath"

func checklongParameterInMethod(method []Method, file string)  {
   println(" FileName:", filepath.Base(file),"\n","package  Name:",method[0].PkgName,)
	for _,l := range method{
      println(" Function Name: ",l.FuncName," |","Parameters: ",len(l.Parameters),"| ","LLOC : ",l.LLOC," | ","Comments: ",l.Comments," | ","Blank Lines: ",l.BlankLines,"|"," complexity:", l.Complexity)
	}
}
