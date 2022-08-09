package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/fatih/color"
)

func msifaGetAllProvinces() []byte {
	logThisln("msifaGetAllProvince")
	color.Cyan("Fetching data from " + urlMsifaProvinces + " ...")

	spaceClient := http.Client{
		Timeout: time.Second * 5, // Timeout after 2 seconds
	}
	req, err := http.NewRequest(http.MethodGet, urlMsifaProvinces, nil)
	if err != nil {
		log.Fatal(err)
	}
	// req.Header.Set("User-Agent", "sarasvati")

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
func msifaGetRegencies(id string) []byte {
	logThisln("msifaGetRegency")
	url := strings.Replace(urlMsifaRegencies, "{ID}", id, -1)
	spaceClient := http.Client{
		Timeout: time.Second * 5, // Timeout after 2 seconds
	}
	color.Cyan("Fetching data from " + url + " ...")
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	// req.Header.Set("User-Agent", "sarasvati")

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

func msifaGetDistricts(id string) []byte {
	logThisln("msifaGetDistricts")
	url := strings.Replace(urlMsifaDistricts, "{ID}", id, -1)
	spaceClient := http.Client{
		Timeout: time.Second * 5, // Timeout after 2 seconds
	}
	color.Cyan("Fetching data from " + url + " ...")
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	// req.Header.Set("User-Agent", "sarasvati")

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

func msifaGetVillages(id string) []byte {
	logThisln("msifaGetVillages")
	url := strings.Replace(urlMsifaVillages, "{ID}", id, -1)
	spaceClient := http.Client{
		Timeout: time.Second * 5, // Timeout after 2 seconds
	}
	color.Cyan("Fetching data from " + url + " ...")
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	// req.Header.Set("User-Agent", "sarasvati")

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
