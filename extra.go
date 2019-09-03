package main

import (
	"archive/zip"
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main()  {
	/*args :=os.Args[1:]
	fmt.Println(args)*/
	reader:= bufio.NewReader(os.Stdin)
	var fileName string
	/*var zipReader zip.File
	var err error*/
	fmt.Fscan(reader,&fileName)
	fmt.Println(fileName)
	zipReader,err  :=zip.OpenReader(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer zipReader.Close()
	for _, file := range zipReader.Reader.File {

		zippedFile, err := file.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer zippedFile.Close()

		targetDir := "./"
		extractedFilePath := filepath.Join(
			targetDir,
			file.Name,
		)

		if file.FileInfo().IsDir() {
			log.Println("Directory Created:", extractedFilePath)
			os.MkdirAll(extractedFilePath, file.Mode())
		} else {
			log.Println("File extracted:", file.Name)

			outputFile, err := os.OpenFile(
				extractedFilePath,
				os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
				file.Mode(),
			)
			if err != nil {
				log.Fatal(err)
			}
			defer outputFile.Close()

			_, err = io.Copy(outputFile, zippedFile)
			if err != nil {
				log.Fatal(err)
			}

		}
	}
}