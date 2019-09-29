package main

var WMC int = 47
var ATFD int = 5
var TCC float64 = 0.3

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
func calculateWMC(st Struct) int {
	wmc := 0
	for _, i := range st.Methods {
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
