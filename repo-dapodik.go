package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
	"github.com/gocolly/colly"
)

func dapoGetAllRegions(idLevelWilayah int, kodeWilayah string) []byte {
	logThisln("dapoGetAllProvinces")

	if kodeWilayah == "" {
		kodeWilayah = "000000"
	}
	replaceIdLevelWilayah := strings.Replace(urlDapoRegions, "{idLevelWilayah}", strconv.Itoa(idLevelWilayah), -1)
	url := strings.Replace(replaceIdLevelWilayah, "{kodeWilayah}", kodeWilayah, -1)
	data := []byte{}

	appColly.OnHTML("body", func(e *colly.HTMLElement) {
		// Initiate empty slice of Sekolah
		// var list []Sekolah

		fmt.Println("Parsing JSON...")
		// parse json to struct
		data = []byte(e.Text)
		err := json.Unmarshal(data, &dapodikRegions)
		if err != nil {
			// if error is not nil
			// print error
			fmt.Println(err)
		}
	})

	appColly.Visit(url)
	return data
}
func dapoGetSchool(kodeWilayah string, district string, bentuk string) []byte {
	logThisln("dapoGetAllProvinces")
	replaceKodeWilayah := strings.Replace(urlDapoSchool, "{kodeWilayah}", kodeWilayah, -1)
	url := strings.Replace(replaceKodeWilayah, "{type}", bentuk, -1)
	// fmt.Println(url)
	data := []byte{}

	appColly.OnHTML("body", func(e *colly.HTMLElement) {
		data = []byte(e.Text)
		logThisln(e.Text)
		err := json.Unmarshal(data, &schools)
		if err != nil {
			fmt.Println(err)
		}
		for index, school := range schools {
			almt := dapoGetAlamat(school.ID)
			// schools[index].IDEncrypt = strings.ReplaceAll(school.IDEncrypt, " ", "")
			schools[index].Alamat = almt
		}
	})
	appColly.Visit(url)
	return data
}
func dapoGetAlamat(id string) string {
	// Initiate blank address value
	address := ""
	// Base url
	// url := "https://sekolah.data.kemdikbud.go.id/index.php/chome/profil/" + id
	url := "https://sekolah.data.kemdikbud.go.id/index.php/Csanitasi/profil?id=" + id
	// url := "https://sekolah.data.kemdikbud.go.id/index.php/Ckesiapantik/profil?id=" + id
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: sekolah.data.kemdikbud.go.id, dapo.kemdikbud.go.id
		colly.AllowedDomains("sekolah.data.kemdikbud.go.id", "dapo.kemdikbud.go.id"),
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36"),
	)

	c.OnError(func(r *colly.Response, err error) {
		color.Red("\033[A\033[D\t\t\t\t\t\t\tRequest error: %d", r.StatusCode)
		doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(r.Body)))
		// doc.Find("font.small").Each(func(i int, s *goquery.Selection) {
		doc.Find("small").Each(func(i int, s *goquery.Selection) {
			text := s.Text()
			if strings.Contains(text, "(master referensi)") {
				add := strings.Split(text, "(")[0]
				// addr := add[1:]
				address = add[:len(add)-1]
				// Printing the address
				color.Green("Found [" + address + "]")
			}
		})

	})

	// c.OnHTML("font.small", func(e *colly.HTMLElement) {
	c.OnHTML("small", func(e *colly.HTMLElement) {
		// Pick the right element
		text := e.Text
		if strings.Contains(text, "(master referensi)") {
			add := strings.Split(text, "(")[0]
			// addr := add[1:]
			address = add[:len(add)-1]
			// Printing the address
			color.Green("Found [" + address + "]")
		}
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		color.Cyan("Fetching [" + id + "]")
	})

	// Start scraping on sekolah.data.kemdikbud.go.id
	c.Visit(url)
	return address
}
