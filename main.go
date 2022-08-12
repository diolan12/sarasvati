package main

import (
	"os"

	"github.com/fatih/color"
	"github.com/gocolly/colly"
)

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

	urlKemenag = "https://api-pesantren-indonesia.vercel.app/pesantren/{ID_Kab}.json"
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
var pondoks = []Pondok{}

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
		logCyanln("Fetching data from " + r.URL.String() + " ...")
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
	case "serve":
		serve(args)
	case "help":
		help(args)
	default:
		color.Red("Command not found!")
		help(args)
	}
}
