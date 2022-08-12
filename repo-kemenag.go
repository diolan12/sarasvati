package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/fatih/color"
)

func kemenagGetFromDistrict(id string) []byte {
	// https://api-pesantren-indonesia.vercel.app/pesantren/3510.json

	// https://ditpdpontren.kemenag.go.id/pdpp/searchpp?search=510035100183
	// https://ditpdpontren.kemenag.go.id/pdpp/profil/17763
	logThisln("kemenagGetFromDistrict")
	url := strings.Replace(urlKemenag, "{ID_Kab}", id, -1)
	color.Cyan("Fetching data from " + url + " ...")

	spaceClient := http.Client{
		Timeout: time.Second * 5, // Timeout after 2 seconds
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	color.Yellow(res.Status)
	if res.Body != nil {
		defer res.Body.Close()
	}
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	return body
}
