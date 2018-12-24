package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/urfave/cli"
)

func main() {

	//var currentFileCount int
	//var destinationFilename string

	app := cli.NewApp()
	app.Name = "gf-rename"
	app.Usage = "Global file rename"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "input, i",
			Value: "",
			Usage: "Input mask.",
		},
		cli.StringFlag{
			Name:  "output, o",
			Value: "",
			Usage: "Output mask.",
		},
	}

	app.Action = func(c *cli.Context) error {

		var currentFileCount int

		fmt.Println("Args = ", len(os.Args))

		files, err := filepath.Glob("*")
		if err != nil {
			log.Fatal(err)
		}

		for _, currentFile := range files {

			currentFileCount = processFile(currentFile, currentFileCount)

			fmt.Println(currentFile)
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func processFile(currentFile string, currentFileCount int) int {

	if currentFile != "gf-rename" && IsFileOrDirectory(currentFile) == "f" {
		currentFileCount++

		destinationFilename := makeDestinationFilename(currentFile, currentFileCount, "n")

		fmt.Println("destinationFilename = " + destinationFilename)

		copyFile(currentFile, destinationFilename)

	}
	return currentFileCount

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
