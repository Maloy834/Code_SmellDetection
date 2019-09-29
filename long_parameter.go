package main

import "path/filepath"

func checkMethodExtract(method []Method, file string) {
	println(" FileName:", filepath.Base(file), "\n", "package  Name:", method[0].PkgName)
	for _, l := range method {
		println(" Function Name: ", l.FuncName, "|", "Struct Used:", l.StructName, "|", "Parameters: ", len(l.Parameters), "| ", "LLOC : ", l.LLOC, " | ", "Comments: ", l.Comments, " | ", "Blank Lines: ", l.BlankLines, "|", " complexity:", l.Complexity, "|", " SelfVarAccessed:", len(l.SelfVarAccessed), "|", " OthervarAccessed: ", len(l.OthersVarAccessed))
	}
}
func CheckStructMatrix(structs []Struct) {
	if len(structs) != 0 {
		println("--------------", "Struct Analysis", "---------------")
	}
	for _, st := range structs {
		println("Stcuct Name: ", st.StructName, "|", " WMC:", st.WMC, "|", " ATFD: ", st.ATFD, "|", " TCC: ", st.TCC, "|", " NDC: ", st.NDC, "|", " NP: ", st.NP)
	}
	//
}
