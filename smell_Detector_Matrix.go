package main

var WMC int = 47
var ATFD int = 5
var TCC float64 = 0.3
var LAA float64=0.3
var ATFD_M int = 5
var FDP int=5


func calculateNDC(st Struct) int {
	ndc := 0
	for i := 0; i < len(st.Methods)-1; i++ {
		for k := i + 1; k < len(st.Methods); k++ {
			if checkCommonAttributeAccess(st.Methods[i], st.Methods[k]) {
				ndc++
			}
		}
	}
	return ndc
}
func checkCommonAttributeAccess(m1 Method, m2 Method) bool {
	for _, i := range m1.SelfVarAccessed {
		for _, j := range m2.SelfVarAccessed {
			if i.right == j.right {
				return true
			}
		}
	}
	return false
}
func calculateNP(st Struct) int {
	n := len(st.Methods)
	if n <= 1 {
		return 0
	}
	return n * (n - 1) / 2
}
func calculateWMC(method []Method) int {
	wmc := 0
	for _, i := range method {
		wmc += i.Complexity
	}
	return wmc
}
func calculateATFD(st Struct) int {
	selectorList := uniqeSelectors{}
	for _, r := range st.Methods {
		for _, m := range r.OthersVarAccessed {
			if !selectorList.exists(m) {
				selectorList.add(m)
			}
		}
	}
	return len(selectorList.selectors)

}
func(m *Method)claculateMethodATFD(){
	selectorList := uniqeSelectors{}
	for _, r:= range m.OthersVarAccessed {
		if !selectorList.exists(r) {
			selectorList.add(r)
		}
	}
	m.ATFD=len(selectorList.selectors)
}
func calculateFDP(st []Struct,method Method) int{
	//count:=0
	structList := uniqueStruct{}
	//println(method.FuncName)
	for _,variable:= range method.OthersVarAccessed{
		for _,s:= range st{
			//println(variable.right)
			if containsVariable(variable.right,s.Attributes){
				/*if method.FuncName=="marshalObject"{
					println(s.StructName+"  "+ variable.right)
				}*/
				if ! structList.existStruct(s){
					structList.addStruct(s)
				}
			}
		}
	}
	return len(structList.structs)
}


func (m *Method)calculateMethodLAA()  {
	if len(m.OthersVarAccessed)==0{
		m.LAA=float64(TCC_Null)
	}else{
		m.LAA=float64(len(m.SelfVarAccessed))/float64(len(m.OthersVarAccessed))
	}
}
func calculateLCOM(st Struct, method []Method) int{
	var m = len(st.Totalmethods)
	var a = len(st.Methods)
	var v = countInstanceVariable(st)
	//println(m,"  | ",a," | ",v)
	var lcom int
    //println("--------- LCOM Calculation-------")
	//println("Total Methods: ", m," |", " Instance Methods: ",a," |", "Instance Variable: ",v)
	if(v==0) {
		//lcom = float64((m * (a / TCC_Null)) / (m - 1))
		t:=(float64(m)*(float64(a)/float64(TCC_Null))/float64(m-1))
		lcom=int(t)
		//lcom=((m*(a/TCC_Null))/(m-1))
	}else {
		t:=(float64(m)*(float64(a)/float64(v))/float64(m-1))
		lcom = int(t)
	}
	return lcom
}
func  countInstanceVariable(st Struct) int {
  var count =0
  //var i,j int
	selfVaraccessedSelector := uniqeSelectors{}
	otherVaraccessedSelector:= uniqeSelectors{}
	for _,s:= range st.Methods{
		for _,t := range s.SelfVarAccessed{
			if !selfVaraccessedSelector.exists(t){
				selfVaraccessedSelector.add(t)
			}
		}
		for _,i:=range s.OthersVarAccessed{
			if !otherVaraccessedSelector.exists(i){
				otherVaraccessedSelector.add(i)
			}
		}
	}
	count= len(selfVaraccessedSelector.selectors)+len(otherVaraccessedSelector.selectors)
	return count
}
func calculateTCC(st Struct) float64 {
	if st.NP == 0 {
		//println("debug")
		return TCC_Null
	}
	return float64(st.NDC) / float64(st.NP)
}
func checkGodStruct(st Struct) bool {
	if st.WMC > WMC && st.TCC < TCC && st.ATFD > ATFD {
		return true
	}

	return false
}
func checkFeatureEnvy(method Method) bool{
	if method.FDP <=FDP && method.LAA < LAA && method.ATFD > ATFD_M{
		return true
	}
	return false
}
func checkDataStruct(st Struct) bool{
	if st.LCOM>80 || st.WMC >50{
		return  true
	}
	return false
}
func checkBrainMethod(method Method)  bool{
	if method.LLOC>65 && method.Complexity>=3 && method.MAXNESTING >=3 && method.NOAV>7{
		return true
	}
	return false
}

func checkLongParameter(method Method) bool{
	if len(method.Parameters)>5 {
		return true
	}
	return false
}
func checkShortGunSurgery(method Method)bool{
	if method.CM>7 && method.CC>10{
		return true
	}
	return false
}