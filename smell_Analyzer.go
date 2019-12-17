package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*func checkMethodExtract(method []Method, file string,structs []Struct) {
	var instanceMethod =0
	if(len(structs)!=0){
		instanceMethod = len(structs[0].Methods)
	}
	println(" FileName:", filepath.Base(file), "\n", "package  Name:", method[0].PkgName,"|","\n","Total Methods: ", len(method),"|","Instance Method: ",instanceMethod)
	if len(method) != 0 {
		println("--------------", "Method Analysis", "---------------")
	}
	for _, l := range method {
		RatioLOCAndComplexity:=fmt.Sprintf("%.12f",l.RatioLOCAndComplexity)
		println(" Function Name: ", l.FuncName, "|", "Struct Used:", l.StructName, "|",  "Parameters: ", len(l.Parameters), "| ", "LLOC : ", l.LLOC, " | ", "Comments: ", l.Comments, " | ", "Blank Lines: ", l.BlankLines, "|", " complexity:", l.Complexity, "|","MAXNESTING",l.MAXNESTING," |" ,"Ratio of LOC over Complexity:",RatioLOCAndComplexity," |", "NOAV: ",l.NOAV," |","FDP: ",l.FDP," |"," SelfVarAccessed:", len(l.SelfVarAccessed), "|", " OthervarAccessed: ", len(l.OthersVarAccessed))
	}
	if len(structs) != 0 {
		println("--------------", "Struct Analysis", "---------------")
	}
	for _, st := range structs {
		//lcom:=fmt.Sprintf("%.12f",st.LCOM)
		tcc:=fmt.Sprintf("%.12f",st.TCC)
		println("Stcuct Name: ", st.StructName, "|", " WMC:", st.WMC, "|", " ATFD: ", st.ATFD, "|", " TCC: ", tcc, "|", " NDC: ", st.NDC, "|", " NP: ", st.NP, "|"," LCOM: ",st.LCOM, "|"," Total Methods: ", len(method),"|","Instance Method: ",len(st.Methods))
	}

}*/
/*func analyseStructs(structs[] Struct)  {
	println("--------------", "Struct Analysis", "---------------")
	for _, st := range structs {
		instanceMethod := len(st.Methods)
		//lcom:=fmt.Sprintf("%.12f",st.LCOM)
		println(" FileName:", st.FileName, "\n", "package  Name:", st.PkgName,"|","\n","Total Methods: ", len(st.Totalmethods),"|","Instance Method: ",instanceMethod)
		tcc:=fmt.Sprintf("%.12f",st.TCC)
		println("Stcuct Name: ", st.StructName, "|", " WMC:", st.WMC, "|", " ATFD: ", st.ATFD, "|", " TCC: ", tcc, "|", " NDC: ", st.NDC, "|", " NP: ", st.NP, "|"," LCOM: ",st.LCOM, "|"," Total Methods: ", len(st.Totalmethods),"|","Instance Method: ",len(st.Methods))
	}

}*/
/*func analyseMethods(methods[] Method){
	println("--------------", "Method Analysis", "---------------")
	for _, l := range methods {
		println(" FileName:", l.FileName, "\n", "package  Name:", l.PkgName)
		RatioLOCAndComplexity:=fmt.Sprintf("%.12f",l.RatioLOCAndComplexity)
		println(" Function Name: ", l.FuncName, "|", "Struct Used:", l.StructName, "|",  "Parameters: ", len(l.Parameters), "| ", "LLOC : ", l.LLOC, " | ", "Comments: ", l.Comments, " | ", "Blank Lines: ", l.BlankLines, "|", " complexity:", l.Complexity, "|","MAXNESTING",l.MAXNESTING," |" ,"Ratio of LOC over Complexity:",RatioLOCAndComplexity," |", "NOAV: ",l.NOAV," |","FDP: ",l.FDP," |"," SelfVarAccessed:", len(l.SelfVarAccessed), "|", " OthervarAccessed: ", len(l.OthersVarAccessed))
	}
}*/
func analyseStruct(structs[] Struct, projectName string)  {
	pos:= strings.Index(projectName,".")
	str:= projectName[0:pos]
	projectName=str
	outputFile,err:= os.Create(projectName+"_StructSumarry.csv")
	checkError("Cannot create File ",err)
	defer outputFile.Close()
	writer := csv.NewWriter(outputFile)
	defer writer.Flush()
	var tcc string
	///instanceMethod := strconv.Itoa(len(st.Methods))
	var data = [] string{
		"FileName"," Struct Name", "Package Name","Total Methods","Instance Methods","WMC","ATFD","TCC","NDC","NP"}
	writer.Write(data)
	for _,st:= range structs{
		totalMethods:=strconv.Itoa(len(st.Totalmethods))
		instanceMethods:=strconv.Itoa(len(st.Methods))
		wmc:= strconv.Itoa(st.WMC)
		atfd:=strconv.Itoa(st.ATFD)
		//lcom:=strconv.Itoa(st.LCOM)
		ndc:=strconv.Itoa(st.NDC)
		np:=strconv.Itoa(st.NP)
		if st.TCC>1{
			tcc="null"
		}else {
			tcc = fmt.Sprintf("%.12f", st.TCC)
		}

		var data=[] string{
			st.FileName,st.StructName,st.PkgName,totalMethods,instanceMethods,wmc,atfd,tcc,ndc,np}
		writer.Write(data)
	}
	//writeIntoCSVFile()

}
func analyseMethods(methods[]Method,projectName string){
	pos:= strings.Index(projectName,".")
	str:= projectName[0:pos]
	projectName=str
	outputFile,err:= os.Create(projectName+"_MethodSumarry.csv")
	checkError("Cannot create File ",err)
	defer outputFile.Close()
	writer := csv.NewWriter(outputFile)
	defer writer.Flush()
	var laa string
	var data = []string{
		"Method Name","File Name","Struct Used","Package Name","Own Variable Accessed","Other Variable Accessed",
		"Cyclomatic Complexity","LLOC","Comments","Exported","Parameters","FDP","LAA","ATFD","MAXNESTING","NOAV","Ratio of LOC over Complexity","CM","CC"}
	writer.Write(data)
	for _,mt:= range methods{
		ownVar:= strconv.Itoa(len(mt.SelfVarAccessed))
		otherVar :=strconv.Itoa(len(mt.OthersVarAccessed))
		complexity:= strconv.Itoa(mt.Complexity)
		lloc:= strconv.Itoa(mt.LLOC)
		cm:=strconv.Itoa(mt.CM)
		cc:=strconv.Itoa(mt.CC)
		comments:=strconv.Itoa(mt.Comments)
		exported:=strconv.FormatBool(mt.isExported)
		parameters:= strconv.Itoa(len(mt.Parameters))
		fdp:= strconv.Itoa(mt.FDP)
		atfd:= strconv.Itoa(mt.ATFD)
		if mt.LAA>1{
			laa="null"
		}else {
			laa = fmt.Sprintf("%.12f", mt.LAA)
		}
		noav:=strconv.Itoa(mt.NOAV)
		maxnesting:=strconv.Itoa(mt.MAXNESTING)
		LOCOverComplexity:=fmt.Sprintf("%.12f",mt.RatioLOCAndComplexity)
		var data=[]string{
			mt.FuncName,mt.FileName,mt.StructName,mt.PkgName,ownVar,otherVar,complexity,lloc,comments,exported,parameters,fdp,laa,atfd,maxnesting,noav,LOCOverComplexity,cm,cc}
		writer.Write(data)
	}
}
func writeCodeSmellSumarry(structs[] Struct,methods[]Method, projectName string)  {
	GodStruct:=0
	//DataStruct:=0
	FeatureEnvy:=0
	BrainMethod:=0
	LongParameter:=0
	shortgun:=0
	for _,st:= range structs{
		if st.GodStruct{
			GodStruct++
		}
		/*if(st.DataStruct){
			DataStruct++
			//println(st.FileName,"    || ",st.StructName )
		}*/
	}
	for _,mt:= range methods{
		if mt.BrainMethod{
			BrainMethod++
		}
		if mt.FeatureEnvy{
			FeatureEnvy++
		}
		if mt.LongParameter{
			LongParameter++
		}
		if mt.ShortGunSurgery{
			shortgun++
		}

	}
	totalstruct:= strconv.Itoa(len(structs))
	godStruct:=strconv.Itoa(GodStruct)
	//dataStruct:=strconv.Itoa(DataStruct)
	featureEnvy:= strconv.Itoa(FeatureEnvy)
	brainMethod:=strconv.Itoa(BrainMethod)
	parameter:=strconv.Itoa(LongParameter)
	ShortGun:=strconv.Itoa(shortgun)
	pos:= strings.Index(projectName,".")
	str:= projectName[0:pos]
	projectName=str
	outputFile,err:= os.Create(projectName+"_CodeSmellSumarry.csv")
	checkError("Cannot create File ",err)
	defer outputFile.Close()
	writer := csv.NewWriter(outputFile)
	defer writer.Flush()
	var data = [][]string{
		{"Total Structs","God Struct", "Feature Envy", "Brain Method", "Long parameter","ShortGun Surgery"},
		{totalstruct,godStruct,featureEnvy,brainMethod,parameter,ShortGun},
	}
	for _, row := range data {
		_ = writer.Write(row)
	}


}
func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
func writeCodeSmellList(structs[] Struct,method[] Method,projectName string){
	pos:= strings.Index(projectName,".")
	str:= projectName[0:pos]
	projectName=str
	outputFile,err:= os.Create(projectName+"_CodeSmellList.csv")
	checkError("Cannot create File ",err)
	defer outputFile.Close()
	writer := csv.NewWriter(outputFile)
	defer writer.Flush()
	var godStruct [] string
	var brainMethod [] string
	var featureEnvy[] string
	var shortGun[] string
	var longparam [] string
	tt:=[][] string{}
	//godStruct.
	for _,st:= range structs{
		if st.GodStruct{
			godStruct = append(godStruct, st.StructName)
		}

	}
	for _,st:= range method{
		if st.BrainMethod{
			brainMethod = append(brainMethod, st.FuncName)
		}
		if st.FeatureEnvy{
			featureEnvy=append(featureEnvy,st.FuncName)
		}
		if st.ShortGunSurgery{
			shortGun=append(shortGun,st.FuncName)
		}
		if st.LongParameter{
			longparam=append(longparam,st.FuncName)
		}
	}
     if len(godStruct)>0{
         godStruct=append([]string{"God Struct"},godStruct...)
         tt=append(tt,godStruct)
	 }
	 if len(brainMethod)>0{
	 	brainMethod=append([]string{"Brain Method"},brainMethod...)
		 tt=append(tt,brainMethod)
	 }
	 if len(featureEnvy)>0{
		 featureEnvy=append([]string{"Feature Envy"},featureEnvy...)
		 tt=append(tt,featureEnvy)
	 }
	 if len(shortGun)>0{
		 shortGun=append([]string{"Shot Gun Surgery"},shortGun...)
		 tt=append(tt,shortGun)
	 }
	 if len(longparam)>0{
		 longparam=append([]string{"Long Parameter"},longparam...)
		 tt=append(tt,longparam)
	 }
	writer.WriteAll(tt)
}

