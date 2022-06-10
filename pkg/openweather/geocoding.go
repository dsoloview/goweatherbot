package openweather

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type CoordinatesResponse []struct {
	Name       string            `json:"name"`
	LocalNames map[string]string `json:"local_names"`
	Lat        float64           `json:"lat"`
	Lon        float64           `json:"lon"`
	Country    string            `json:"country"`
	State      string            `json:"state,omitempty"`
}

type Coordinates struct {
	Latitude  float64
	Longitude float64
	City      string
}

const COORDINATES_URL_TEMPLATE = "http://api.openweathermap.org/geo/1.0/direct?q=%v&limit=1&appid=%v"

func getCoordinates(city string) (Coordinates, error) {

	url := fmt.Sprintf(COORDINATES_URL_TEMPLATE, city, os.Getenv("OPENWEATHER_API"))
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	if resp.StatusCode != 200 {
		return Coordinates{}, fmt.Errorf("status code %v", resp.StatusCode)
	}

	data := getCoordinatesJson(resp)

	if len(data) == 0 {
		return Coordinates{}, fmt.Errorf("city not found")
	}

	coordinates := Coordinates{
		Longitude: data[0].Lon,
		Latitude:  data[0].Lat,
		City:      data[0].LocalNames["ru"],
	}
	return coordinates, nil

}

func getCoordinatesJson(resp *http.Response) CoordinatesResponse {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}

	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err.Error())
	}

	var data CoordinatesResponse

	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println(err)
	}

	return data
}
