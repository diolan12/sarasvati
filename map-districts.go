package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
)

func mapMsifaDistricts(args []string) {
	logThisln("mapMsifaDistricts")
	helperLoad(indexFile, &provinces)
	folderProvince := ""
	folderDistrict := ""
	idLevelWilayah := 0
	kodeWilayah := ""
	for _, province := range provinces {
		if args[3][:2] == province.ID {
			color.Magenta("Mapping [" + province.ID + "] " + province.Name + " ...")
			folderProvince = outputDir + "/" + province.ID + "-" + province.Name
			helperLoad(folderProvince+"/"+indexMainFile, &regencies)
			if _, err := os.Stat(folderProvince); errors.Is(err, os.ErrNotExist) {
				err := os.Mkdir(folderProvince, os.ModePerm)
				if err != nil {
					log.Println(err)
				}
			}
			for _, regency := range regencies {
				if args[3] == regency.ID {
					color.Magenta("Mapping [" + regency.ID + "] " + regency.Dapodik.KodeWilayah + " " + regency.Name + " ...")
					folderDistrict = folderProvince + "/" + regency.ID + "-" + regency.Name
					idLevelWilayah = regency.Dapodik.IDLevelWilayah
					kodeWilayah = regency.Dapodik.KodeWilayah
					if _, err := os.Stat(folderDistrict); errors.Is(err, os.ErrNotExist) {
						err := os.Mkdir(folderDistrict, os.ModePerm)
						if err != nil {
							log.Println(err)
						}
					}
					jsonByte := msifaGetDistricts(regency.ID)
					jsonErr := json.Unmarshal(jsonByte, &districtsMsifa)
					if jsonErr != nil {
						log.Fatal(jsonErr)
					}
					break
				}
			}
			break
		}
	}
	if len(regencies) == 0 {
		color.Red("District not found")
		color.Red("Make sure the province is indexed or regency ID is correct")
	} else {
		mapDapoDistrict(idLevelWilayah, kodeWilayah)
		mapDistrictsMerge(folderDistrict)
	}
}
func mapDapoDistrict(idLevelWilayah int, kodeWilayah string) {
	logThisln("mapDapoDistrict")
	jsonByte := dapoGetAllRegions(idLevelWilayah, kodeWilayah)
	jsonErr := json.Unmarshal(jsonByte, &dapodikRegions)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
}
func mapDistrictsMerge(folderDistrict string) {
	targetFile := folderDistrict + "/" + indexMainFile
	fmt.Println(targetFile)
	logThisln("mapDistrictsMerge")
	for _, district := range districtsMsifa {
		color.Cyan("Merging [" + district.ID + "] (" + district.Name + ")")
		for _, region := range dapodikRegions {
			replaceKec := strings.ReplaceAll(region.Name, "Kec. ", "")
			current := strings.ToUpper(replaceKec)
			if district.Name == current {
				region.KodeWilayah = strings.Replace(region.KodeWilayah, " ", "", -1)
				color.Green("\tMatch! [" + region.KodeWilayah + "] (" + region.Name + ")")
				districts = append(districts, District{
					ID:        district.ID,
					Name:      district.Name,
					RegencyID: district.RegencyID,
					Dapodik:   region,
				})
			}
		}
	}

	fmt.Println("Writing data to " + targetFile + " ...")
	fjson, _ := json.MarshalIndent(districts, "", " ")
	err := ioutil.WriteFile(targetFile, fjson, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
