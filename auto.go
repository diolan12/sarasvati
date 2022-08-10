package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/fatih/color"
)

var target = "0"
var targets []string

func auto(args []string) {
	count := 0
	color.Yellow("[system] automation process started")
	if len(args) > 2 {
		color.Yellow("limit: " + args[2])
	}
	// helperLoad(outputDir+"/"+indexMainFile, &provinces)
	// normal set rand 60
	for check() && count < 1000 {
		pickTarget()
		rand.Seed(time.Now().UnixNano())
		random := time.Duration(rand.Int31n(5))
		// random := time.Duration(rand.Int31n(60))
		color.Yellow("[bot] scrapping (" + target + ") in " + strconv.Itoa(int(random)) + " second")
		time.Sleep(random * time.Second)
		// time.Sleep(1 * time.Second)
		goIndex()
		if len(args) > 2 {
			count++
		}
	}
}
func pickTarget() {
	logThisln("pickTarget")
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(targets), func(i, j int) { targets[i], targets[j] = targets[j], targets[i] })
	// fmt.Println(targets)
	color.Yellow("seed size:" + strconv.Itoa(len(targets)))
	random := rand.Intn(len(targets)-0) + 0
	target = targets[random]
}
func goIndex() {
	logThisln("goIndex")
	provinces = []Province{}
	regencies = []Regency{}
	districts = []District{}
	villages = []Village{}
	schools = []School{}
	color.Green("[bot] scrapping " + target + " ...")
	args := []string{"bot", "index", "map", target}
	defer index(args)
}
func check() bool {
	logThisln("check")
	targets = []string{}
	unindexedProvinces := false
	unindexedRegencies := false
	unindexedDistricts := false
	folderRoot := outputDir
	helperLoads(folderRoot, &provinces)
	for _, province := range provinces {
		folderProvince := folderRoot + "/" + province.ID + "-" + province.Name
		if _, err := os.Stat(folderProvince); errors.Is(err, os.ErrNotExist) {
			targets = append(targets, province.ID)
			// color.Magenta(province.ID)
			unindexedProvinces = true
		} else {
			fmt.Println("\nhelper load regencies")
			helperLoads(folderProvince, &regencies)
			fmt.Println(folderProvince + "/" + indexMainFile)
			for _, regency := range regencies {
				folderRegency := folderProvince + "/" + regency.ID + "-" + regency.Name
				if _, err := os.Stat(folderRegency); errors.Is(err, os.ErrNotExist) {
					targets = append(targets, regency.ID)
					// color.Magenta(regency.ID)
					unindexedRegencies = true
				} else {
					fmt.Println("\nhelper load districts")
					helperLoads(folderRegency, &districts)
					fmt.Println(folderRegency + "/" + indexMainFile)
					for _, district := range districts {
						folderDistrict := folderRegency + "/" + district.ID + "-" + district.Name
						if _, err := os.Stat(folderDistrict + "/" + district.ID + "-" + district.Name); errors.Is(err, os.ErrNotExist) {
							targets = append(targets, district.ID)
							// color.Magenta(district.ID)
							unindexedDistricts = true
						}
					}
				}
			}
		}
	}
	fmt.Println(unindexedProvinces || unindexedRegencies || unindexedDistricts)
	return unindexedProvinces || unindexedRegencies || unindexedDistricts
}
