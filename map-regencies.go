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

func mapMsifaRegencies(args []string) {
	logThisln("mapperMsifaRegency")
	helperLoad(indexFile, &provinces)
	folderProvince := ""
	idLevelWilayah := 0
	kodeWilayah := ""
	for _, province := range provinces {
		if args[3] == province.ID {
			color.Magenta("Mapping " + province.Name + " ...")
			folderProvince = outputDir + "/" + province.ID + "-" + province.Name
			idLevelWilayah = province.Dapodik.IDLevelWilayah
			kodeWilayah = province.Dapodik.KodeWilayah
			if _, err := os.Stat(folderProvince); errors.Is(err, os.ErrNotExist) {
				err := os.Mkdir(folderProvince, os.ModePerm)
				if err != nil {
					log.Println(err)
				}

			}
			jsonByte := msifaGetRegencies(province.ID)
			jsonErr := json.Unmarshal(jsonByte, &regenciesMsifa)
			if jsonErr != nil {
				// log.Fatal(jsonErr)
				color.Red(jsonErr.Error())
				panic(jsonErr)
			}
			break
		}
	}
	if len(regenciesMsifa) == 0 {
		color.Red("Regency not found")
	} else {
		mapDapoRegencies(idLevelWilayah, kodeWilayah)
		mapRegenciesMerge(folderProvince)
	}
}
func mapDapoRegencies(idLevelWilayah int, kodeWilayah string) {
	logThisln("mapDapoRegencies")
	jsonByte := dapoGetAllRegions(idLevelWilayah, kodeWilayah)
	jsonErr := json.Unmarshal(jsonByte, &dapodikRegions)
	if jsonErr != nil {
		// log.Fatal(jsonErr)
		color.Red(jsonErr.Error())
		panic(jsonErr)
	}
}
func mapRegenciesMerge(folderProvince string) {
	targetFile := folderProvince + "/" + indexMainFile
	logThisln("mapperRegencyMerge")
	for _, regency := range regenciesMsifa {
		color.Cyan("Merging [" + regency.ID + "] (" + regency.Name + ")")
		for _, region := range dapodikRegions {
			replaceKab := strings.ReplaceAll(region.Name, "Kab. ", "Kabupaten ")
			current := strings.ToUpper(replaceKab)
			if regency.Name == current {
				region.KodeWilayah = strings.Replace(region.KodeWilayah, " ", "", -1)
				color.Green("\tMatch! [" + region.KodeWilayah + "] (" + current + ")")
				regencies = append(regencies, Regency{
					ID:         regency.ID,
					Name:       regency.Name,
					ProvinceID: regency.ProvinceID,
					Dapodik:    region,
				})
			}
		}
	}

	fmt.Println("Writing data to " + targetFile + " ...")
	fjson, _ := json.MarshalIndent(regencies, "", " ")
	err := ioutil.WriteFile(targetFile, fjson, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
