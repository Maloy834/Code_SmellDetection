package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	//reader := bufio.NewReader(os.Stdin)
	var fileName string
	flag.StringVar(&fileName, "file", "", "Usage")

	flag.Parse()
	//fmt.Fscan(reader, &fileName)

	files, err := Unzip(fileName, "output-folder")
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println("Unzipped:\n" + strings.Join(files, "\n"))
	ReadAllfiles(files,fileName)
	os.RemoveAll("output-folder")
}

// Unzip will decompress a zip archive, moving all files and folders
// within the zip file (parameter 1) to an output directory (parameter 2).
func Unzip(src string, dest string) ([]string, error) {

	var filenames []string

	readfile, err := zip.OpenReader(src)
	if err != nil {
		println("debug")
		return filenames, err
	}
	defer readfile.Close()
	//println("debug")
	for _, f := range readfile.File {
		fpath := filepath.Join(dest, f.Name)
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s: illegal file path", fpath)
		}

		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {
			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// Make File
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return filenames, err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return filenames, err
		}

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}

		_, err = io.Copy(outFile, rc)

		// Close the file without defer to close before next iteration of loop
		outFile.Close()
		rc.Close()

		if err != nil {
			return filenames, err
		}
	}
	return filenames, nil
}
func ReadAllfiles(file []string,projectName string) {
	var all_struct []Struct
	var all_method [] Method
	for i := 0; i < len(file); i++ {
		//if os.Stat(file[i])
		f, err := os.Stat(file[i])

		if err != nil {
			log.Fatal(err)
		}
		if f.IsDir() == false && filepath.Ext(file[i]) == ".go" {
			structs, methods := parseFile(file[i])
			/*for _,st:= range structs{
				all_struct=append(all_struct,st)
			}*/
			//all_struct=append(all_struct,structs)
			//println(len(structs))
			for _, mlist := range methods {
				for i, c := range structs {
					//structs[i].addAllMethods(methods)
					structs[i].Totalmethods = methods
					if mlist.PkgName == c.PkgName && mlist.StructName == c.StructName {
						structs[i].addUsedMethod(mlist)
						//structs[i].Totalmethods = methods
					}
				}
			}
			for _, st := range structs {

				all_struct = append(all_struct, st)
			}
			for _, mt := range methods {
				all_method = append(all_method, mt)
			}


		}

	}
	for i, st := range all_struct {
		st.NDC = calculateNDC(st)
		st.NP = calculateNP(st)
		st.ATFD = calculateATFD(st)
		st.TCC = calculateTCC(st)
		st.WMC = calculateWMC(st.Methods)
		st.ATFD = calculateATFD(st)
		st.TCC = calculateTCC(st)
		st.LCOM= calculateLCOM(st,st.Methods)
		st.GodStruct = checkGodStruct(st)
		//st.DataStruct=checkDataStruct(st)
		all_struct[i] = st
	}

	for j,m := range all_method{
		m.calculateCallerMethod(all_method)
		//println(m.FuncName,"  ",m.CM,"" ,m.CC)
		m.FDP=calculateFDP(all_struct,m)
		m.FeatureEnvy=checkFeatureEnvy(m)
		m.ShortGunSurgery=checkShortGunSurgery(m)
		m.BrainMethod=checkBrainMethod(m)
		m.LongParameter=checkLongParameter(m)
		all_method[j]=m
	}

	analyseStruct(all_struct,projectName)
	analyseMethods(all_method,projectName)
	writeCodeSmellSumarry(all_struct,all_method,projectName)
	writeCodeSmellList(all_struct,all_method,projectName)
	//checkMethodExtract(methods, file[i],structs)
}
