package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/fatih/color"
)

func mapMsifaProvinces() {
	logThisln("mapperMsifaProvinces")
	jsonByte := msifaGetAllProvinces()
	jsonErr := json.Unmarshal(jsonByte, &provincesMsifa)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	mapDapoProvinces()
}
func mapDapoProvinces() {
	logThisln("mapDapoProvinces")
	jsonByte := dapoGetAllRegions(0, "")
	jsonErr := json.Unmarshal(jsonByte, &dapodikRegions)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	mapProvincesMerge()
}

func mapProvincesMerge() {
	logThisln("indexMerge")
	for _, province := range provincesMsifa {
		color.Cyan("merging [" + province.ID + "] " + province.Name)
		for _, dapodik := range dapodikRegions {
			removeProv := strings.ReplaceAll(dapodik.Name, "Prov. ", "")
			removeDot := strings.ReplaceAll(removeProv, ".", "")
			current := strings.ToUpper(removeDot)
			if current == province.Name {
				dapodik.KodeWilayah = strings.Replace(dapodik.KodeWilayah, " ", "", -1)
				color.Green("\tMatch! [" + dapodik.KodeWilayah + "] " + province.Name)
				provinces = append(provinces, Province{
					ID:      province.ID,
					Name:    province.Name,
					Dapodik: dapodik,
				})
			}
		}
	}
	fmt.Println("Writing data to file...")
	fjson, _ := json.MarshalIndent(provinces, "", " ")
	err := ioutil.WriteFile(indexFile, fjson, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
