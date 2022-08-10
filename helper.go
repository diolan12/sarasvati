package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/fatih/color"
)

func logo(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		color.Yellow(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func logThis(message any) {
	if debug {
		log.Print(message)
	}
}
func logThisln(message any) {
	if debug {
		log.Println(message)
	}
}
func logRedln(message string) {
	if debug {
		color.Red(message)
	}
}
func logGreenln(message string) {
	if debug {
		color.Green(message)
	}
}
func logBlueln(message string) {
	if debug {
		color.Blue(message)
	}
}

func helperLoads(pathToIndex string, ptr any) {
	logThisln("helperLoad")
	pathOfIndex := pathToIndex + "/" + indexMainFile
	if _, err := os.Stat(pathOfIndex); errors.Is(err, os.ErrNotExist) {
		color.Red("no main index file found!")
		color.Red("path = " + pathToIndex)
		os.RemoveAll(pathToIndex)
	} else {
		// Open our jsonFile
		jsonFile, err := os.Open(pathOfIndex)
		// if we os.Open returns an error then handle it
		if err != nil {
			fmt.Println(err)
		}
		if debug {
			color.Cyan("Successfully Opened " + pathOfIndex)
		}
		// defer the closing of our jsonFile so that we can parse it later on
		defer jsonFile.Close()
		byteValue, _ := ioutil.ReadAll(jsonFile)
		json.Unmarshal(byteValue, &ptr)
	}
}

func isOdd(x int) bool {
	return x%2 != 0
}
