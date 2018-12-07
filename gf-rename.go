package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func main() {

	var currentFileCount int
	var destinationFilename string

	files, err := filepath.Glob("*")
	if err != nil {
		log.Fatal(err)
	}

	for _, currentFile := range files {

		if currentFile != "main" && IsFileOrDirectory(currentFile) == "f" {
			currentFileCount++

			//fmt.Printf("\n**\n\n")

			//fmt.Println(currentFile)
			//fmt.Println(extractFileExtension)

			destinationFilename = makeDestinationFilename(currentFile, currentFileCount, "n")

			fmt.Println("destinationFilename = " + destinationFilename)

			copyFile(currentFile, destinationFilename)
		}
	}
}

func extractFileExtension(fileName string) string {
	for i := len(fileName); i > 0; i-- {

		//fmt.Println(fileName[i-1 : i])

		if fileName[i-1:i] == "." {
			return fileName[i-1 : len(fileName)]
		}

	}

	return ""
}

func makeDestinationFilename(sourceFilename string, currentFileCount int, flags string) string {

	var destinationFilename string

	const padding = "000"

	countAsString := strconv.Itoa(currentFileCount)

	//fmt.Println(padding[0:len(padding)-len(countAsString)] + countAsString)

	destinationFilename = padding[0:len(padding)-len(countAsString)] + countAsString

	//fmt.Println("Extension = " + extractFileExtension(sourceFilename))

	return destinationFilename + extractFileExtension(sourceFilename)

}

func copyFile(sourceFilename string, destinationFilename string) {

	srcFile, err := os.Open(sourceFilename)
	check(err)
	defer srcFile.Close()

	destFile, err := os.Create(destinationFilename) // creates if file doesn't exist
	check(err)
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile) // check first var for number of bytes copied
	check(err)

	err = destFile.Sync()
	check(err)

}

func IsFileOrDirectory(fileOrDirName string) string {

	fi, err := os.Stat(fileOrDirName)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	switch mode := fi.Mode(); {
	case mode.IsDir():
		// do directory stuff
		fmt.Println("directory")
		return "d"
	case mode.IsRegular():
		// do file stuff
		fmt.Println("file")
		return "f"
	}
	return ""
}

func check(err error) {
	if err != nil {
		fmt.Println("Error : %s", err.Error())
		os.Exit(1)
	}
}
