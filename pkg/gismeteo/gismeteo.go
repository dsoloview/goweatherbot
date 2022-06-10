package gismeteo

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	TOKEN       = "56b30cb255.3443075"
	WEATHER_URL = "https://api.gismeteo.net/v2/weather/current/"
	SEARCH_URL  = "https://api.gismeteo.net/v2/search/cities/?query="
)

type Object struct {
	id   int
	name string
	url  string
	kind string
}

type Country struct {
	code  string
	name  string
	nameP string
}

type District struct {
	name  string
	nameP string
}

type SubDistrict struct {
	name  string
	nameP string
}

func GetWeather(city string) {

	req, err := http.NewRequest("GET", WEATHER_URL, nil)
	if err != nil {
		log.Println(err)
	}
	req.Header.Set("X-Gismeteo-Token", TOKEN)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
}

func SearchCity(city string) {
	fmt.Println(SEARCH_URL + city)
	req, err := http.NewRequest("GET", SEARCH_URL+city, nil)
	if err != nil {
		log.Println(err)
	}
	req.Header.Set("X-Gismeteo-Token", "56b30cb255.3443075")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	if resp.StatusCode != 200 {
		s := string(body)
		log.Println(s)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(resp.Body)

}
