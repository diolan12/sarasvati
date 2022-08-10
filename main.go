package main

import (
	"os"

	"github.com/fatih/color"
	"github.com/gocolly/colly"
)

// emsifa rest api struct
type MsifaProvince struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type MsifaRegency struct {
	ID         string `json:"id"`
	ProvinceID string `json:"province_id"`
	Name       string `json:"name"`
}
type MsifaDistrict struct {
	ID        string `json:"id"`
	RegencyID string `json:"regency_id"`
	Name      string `json:"name"`
}

// dapodik resp api struct
type DapodikRegion struct {
	Name            string `json:"nama"`
	KodeWilayah     string `json:"kode_wilayah"`
	IDLevelWilayah  int    `json:"id_level_wilayah"`
	MstLevelWilayah int    `json:"mst_level_wilayah"`
}
type School struct {
	Nama      string `json:"nama"`
	ID        string `json:"sekolah_id"`
	IDEncrypt string `json:"sekolah_id_enkrip"`
	Alamat    string `json:"alamat"`
}

// app personalized struct
type Province struct {
	ID      string        `json:"id"`
	Name    string        `json:"name"`
	Dapodik DapodikRegion `json:"dapodik"`
}
type Regency struct {
	ID         string        `json:"id"`
	ProvinceID string        `json:"province_id"`
	Name       string        `json:"name"`
	Dapodik    DapodikRegion `json:"dapodik"`
}
type District struct {
	ID        string        `json:"id"`
	RegencyID string        `json:"regency_id"`
	Name      string        `json:"name"`
	Dapodik   DapodikRegion `json:"dapodik"`
}
type Village struct {
	ID         string `json:"id"`
	DistrictID string `json:"district_id"`
	Name       string `json:"name"`
}

// App constants
// var allowedDomains = [...]string{"sekolah.data.kemdikbud.go.id", "dapo.kemdikbud.go.id"}
// var allowedDomains = {"sekolah.data.kemdikbud.go.id", "dapo.kemdikbud.go.id"}
// provinces > regencies > districts > villages
const (
	debug         = false
	outputDir     = "output"
	indexMainFile = "index.json"

	urlMsifaProvinces = "http://www.emsifa.com/api-wilayah-indonesia/api/provinces.json"
	urlMsifaRegencies = "http://www.emsifa.com/api-wilayah-indonesia/api/regencies/{ID}.json"
	urlMsifaDistricts = "http://www.emsifa.com/api-wilayah-indonesia/api/districts/{ID}.json"
	urlMsifaVillages  = "http://www.emsifa.com/api-wilayah-indonesia/api/villages/{ID}.json"

	urlDapoRegions = "https://dapo.kemdikbud.go.id/rekap/dataSekolah?id_level_wilayah={idLevelWilayah}&kode_wilayah={kodeWilayah}&semester_id=20221"
	urlDapoSchool  = "https://dapo.kemdikbud.go.id/rekap/progresSP?id_level_wilayah=3&kode_wilayah={kodeWilayah}&semester_id=20221&bentuk_pendidikan_id={type}"
)

// App memory
var provincesMsifa []MsifaProvince
var regenciesMsifa []MsifaRegency
var districtsMsifa []MsifaDistrict

var dapodikRegions []DapodikRegion

var provinces []Province
var regencies []Regency
var districts []District
var villages []Village

var schools = []School{}

var appColly = colly.NewCollector(
	colly.AllowedDomains("sekolah.data.kemdikbud.go.id", "dapo.kemdikbud.go.id"),
	colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36"),
)

func main() {
	logo("logo.txt")
	logo("text.txt")
	args := os.Args

	// Before making a request print "Visiting ..."
	appColly.OnRequest(func(r *colly.Request) {
		color.Cyan("Fetching data from " + r.URL.String() + " ...")
	})

	if len(args) < 2 {
		color.Red("Please provide a command!")
		help(args)
		os.Exit(1)

	}
	command := args[1]
	switch command {
	case "work":
		work(args)
	case "buf":
		bufWrite()
	case "auto":
		auto(args)
	case "index":
		index(args)
	case "help":
		help(args)
	}
}
