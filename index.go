package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

var indexFile string

func index(args []string) {
	logThisln("index")
	fmt.Println(len(args))
	fmt.Println(args)
	indexFile = outputDir + "/" + indexMainFile

	if _, err := os.Stat(outputDir); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(outputDir, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
	if len(args) == 2 {
		indexShow(args)
	} else {
		switch args[2] {
		case "show":
			indexShow(args)
		case "map":
			indexMap(args)
		}
	}
}

func indexShow(args []string) {
	logThisln("indexShow")
	// color.Yellow("Yellow means not indexed yet")
	// color.Green("Green means indexed")
	if len(args) < 4 {
		showProvinces()
	} else {
		switch len(args[3]) {
		case 2:
			showRegencies(args[3])
		case 4:
			showDistricts(args[3])
		}
	}

}

func indexMap(args []string) {
	logThisln("indexMap")
	if len(args) < 4 {
		fmt.Println("Mapping all provinces...")
		mapMsifaProvinces()
	} else {
		switch len(args[3]) {
		case 2:
			fmt.Println("Mapping all regencies for " + args[3] + "...")
			mapMsifaRegencies(args)
		case 4:
			fmt.Println("Mapping all districts for " + args[3] + "...")
			mapMsifaDistricts(args)
		case 7:
			fmt.Println("Mapping all villages for " + args[3] + "...")
			mapMsifaVillages(args)
		}
	}
}
