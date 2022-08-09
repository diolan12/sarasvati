package main

import (
	"errors"
	"os"

	"github.com/fatih/color"
)

func showProvinces() {
	logThisln("showProvinces")
	helperLoad(indexFile, &provinces)
	for _, province := range provinces {
		folder := province.ID + "-" + province.Name
		if _, err := os.Stat(outputDir + "/" + folder); errors.Is(err, os.ErrNotExist) {
			color.Yellow("[" + province.ID + "] " + province.Dapodik.KodeWilayah + " " + province.Name)
		} else {
			color.Green("[" + province.ID + "] " + province.Dapodik.KodeWilayah + " " + province.Name)
		}
	}
}
func showRegencies(id string) {
	logThisln("showRegencies")
	logThisln("id: " + id)
	helperLoad(indexFile, &provinces)
	for _, province := range provinces {
		logThis("match " + id + " == " + province.ID + " : ")
		logThisln(id == province.ID)
		if id == province.ID {
			folderProvince := outputDir + "/" + province.ID + "-" + province.Name
			if _, err := os.Stat(folderProvince); errors.Is(err, os.ErrNotExist) {
			} else {
				color.Blue("Result for [" + province.ID + "] " + province.Dapodik.KodeWilayah + " " + province.Name + ":\n")
				helperLoad(folderProvince+"/"+indexMainFile, &regencies)
				for _, regency := range regencies {
					if regency.ProvinceID == province.ID {
						folderRegency := folderProvince + "/" + regency.ID + "-" + regency.Name
						if _, err := os.Stat(folderRegency); errors.Is(err, os.ErrNotExist) {
							color.Yellow("[" + regency.ID + "] " + regency.Dapodik.KodeWilayah + " " + regency.Name)
						} else {
							color.Green("[" + regency.ID + "] " + regency.Dapodik.KodeWilayah + " " + regency.Name)
						}
					}

				}
				break
			}
			break
		}
	}

}
func showDistricts(id string) {
	logThisln("showDistricts")
	logThisln("id: " + id)
	provinceID := id[:2]
	helperLoad(indexFile, &provinces)
	for _, province := range provinces {
		logThis("match " + provinceID + " == " + province.ID + " : ")
		logThisln(provinceID == province.ID)
		if provinceID == province.ID {
			folderProvince := outputDir + "/" + province.ID + "-" + province.Name
			if _, err := os.Stat(folderProvince); errors.Is(err, os.ErrNotExist) {
			} else {
				color.Blue("Result for [" + province.ID + "] " + province.Dapodik.KodeWilayah + " " + province.Name + ":\n")
				helperLoad(folderProvince+"/"+indexMainFile, &regencies)
				for _, regency := range regencies {
					logThis("match " + provinceID + " == " + province.ID + " : ")
					logThisln(regency.ProvinceID == province.ID && regency.ID == id)
					folderRegency := folderProvince + "/" + regency.ID + "-" + regency.Name
					if regency.ProvinceID == province.ID && regency.ID == id {
						color.Blue("Result for [" + regency.ID + "] " + regency.Dapodik.KodeWilayah + " " + regency.Name + ":\n")
						helperLoad(folderRegency+"/"+indexMainFile, &districts)
						for _, districts := range districts {
							logThisln("this districts? " + districts.ID[:4] + " == " + id)
							logThis(districts.ID[:4] == id)
							if districts.ID[:4] == id {
								folderDistrict := folderRegency + "/" + districts.ID + "-" + districts.Name
								if _, err := os.Stat(folderDistrict); errors.Is(err, os.ErrNotExist) {
									color.Yellow("[" + districts.ID + "] " + districts.Dapodik.KodeWilayah + " " + districts.Name)
								} else {
									color.Green("[" + districts.ID + "] " + districts.Dapodik.KodeWilayah + " " + districts.Name)
								}
							}

						}
						break
					}

				}
			}
			break
		}
	}
}

// func showSchool()
