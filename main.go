package main

import (
	"archive/zip"
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	reader:= bufio.NewReader(os.Stdin)
	var fileName string
	fmt.Fscan(reader,&fileName)

	files, err := Unzip(fileName, "output-folder")
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println("Unzipped:\n" + strings.Join(files, "\n"))
	ReadAllfiles(files)
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

		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)

		// Check for ZipSlip. More Info: http://bit.ly/2MsjAWE
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
func  ReadAllfiles(file []string)  {

 for i:=0;i<len(file);i++{
 	//if os.Stat(file[i])
 	f,err :=os.Stat(file[i])

	 if err != nil {
		 log.Fatal(err)
	 }
	 if f.IsDir() == false && filepath.Ext(file[i])==".go" {
		  structs, methods :=parseFile(file[i])
		 checklongParameterInMethod(methods,file[i])
		 for _, mlist := range methods {
			 for i, c := range structs {
				 if mlist.PkgName == c.PkgName && mlist.StructName == c.StructName {
					 structs[i].addMethod(mlist)
				 }
			 }
		 }

	 }

 }
}