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

func mapMsifaVillages(args []string) {
	logThisln("mapMsifaVillages")
	helperLoads(indexFile, &provinces)
	folderProvince := ""
	folderRegency := ""
	folderDistrict := ""
	districtName := ""
	kodeWilayah := ""
	for _, province := range provinces {
		if args[3][:2] == province.ID {
			color.Magenta("Mapping [" + province.ID + "] " + province.Name + " ...")
			folderProvince = outputDir + "/" + province.ID + "-" + province.Name
			helperLoads(folderProvince, &regencies)
			if _, err := os.Stat(folderProvince); errors.Is(err, os.ErrNotExist) {
				err := os.Mkdir(folderProvince, os.ModePerm)
				if err != nil {
					log.Println(err)
				}
			}
			for _, regency := range regencies {
				if args[3][:4] == regency.ID {
					color.Magenta("Mapping [" + regency.ID + "] " + regency.Dapodik.KodeWilayah + " " + regency.Name + " ...")
					folderRegency = folderProvince + "/" + regency.ID + "-" + regency.Name
					helperLoads(folderRegency, &districts)
					for _, district := range districts {
						if args[3] == district.ID {
							color.Magenta("Mapping [" + district.ID + "] " + district.Dapodik.KodeWilayah + " " + district.Name + " ...")
							folderDistrict = folderRegency + "/" + district.ID + "-" + district.Name
							kodeWilayah = district.Dapodik.KodeWilayah
							districtName = district.ID + "-" + district.Name
							fmt.Println(folderDistrict)
							if _, err := os.Stat(folderDistrict); errors.Is(err, os.ErrNotExist) {
								err := os.Mkdir(folderDistrict, os.ModePerm)
								if err != nil {
									log.Println(err)
								}
							}
							jsonByte := msifaGetVillages(district.ID)
							jsonErr := json.Unmarshal(jsonByte, &villages)
							if jsonErr != nil {
								log.Fatal(jsonErr)
							}
							break
						}
					}
					break
				}
			}
			break
		}
	}
	if len(regencies) == 0 {
		color.Red("Regency not found")
		color.Red("Make sure the regency is indexed or regency ID is correct")
	} else {
		mapDapoSchool(kodeWilayah, districtName, "kb")
		mapVillagesMerge(folderDistrict, districtName, "kb")

		mapDapoSchool(kodeWilayah, districtName, "tk")
		mapVillagesMerge(folderDistrict, districtName, "tk")

		mapDapoSchool(kodeWilayah, districtName, "sd")
		mapVillagesMerge(folderDistrict, districtName, "sd")

		mapDapoSchool(kodeWilayah, districtName, "smp")
		mapVillagesMerge(folderDistrict, districtName, "smp")

		mapDapoSchool(kodeWilayah, districtName, "sma")
		mapVillagesMerge(folderDistrict, districtName, "sma")

		mapDapoSchool(kodeWilayah, districtName, "smk")
		mapVillagesMerge(folderDistrict, districtName, "smk")

		mapDapoSchool(kodeWilayah, districtName, "slb")
		mapVillagesMerge(folderDistrict, districtName, "slb")
	}
}
func mapDapoSchool(kodeWilayah string, district string, bentuk string) {
	jsonByte := dapoGetSchool(kodeWilayah, district, bentuk)
	err := json.Unmarshal(jsonByte, &schools)
	if err != nil {
		// log.Fatal(err)
		color.Red(err.Error())
		color.Red("Error byte: " + string(jsonByte))
		panic(err)
	}
}

func mapVillagesMerge(folderDistrict string, district string, bentuk string) {
	logThisln("mapVillagesMerge")
	targetIndex := folderDistrict + "/" + indexMainFile
	targetFile := folderDistrict + "/" + district + "-" + bentuk + ".json"
	fmt.Println(targetIndex)
	// for index, school := range schools {
	// 	alamat := dapoGetAlamat(school.ID)
	// 	color.Cyan("Writing [" + school.ID + "] (" + school.Nama + ")")
	// 	schools[index].Alamat = alamat
	// 	fmt.Println(alamat)
	// }

	fmt.Println("Writing data to " + targetIndex + " ...")
	fjsonIndex, _ := json.MarshalIndent(villages, "", " ")
	errIndex := ioutil.WriteFile(targetIndex, fjsonIndex, 0644)
	if errIndex != nil {
		log.Fatal(errIndex)
	}

	for index, school := range schools {
		schools[index].IDEncrypt = strings.ReplaceAll(school.IDEncrypt, " ", "")
	}
	fmt.Println("Writing data to " + targetFile + " ...")
	fjson, _ := json.MarshalIndent(schools, "", " ")
	err := ioutil.WriteFile(targetFile, fjson, 0644)
	if err != nil {
		log.Fatal(err)
	}
	schools = []School{}
}
