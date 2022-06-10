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
	Name       string   `json:"name"`
	LocalNames struct{} `json:"local_names"`
	Lat        float64  `json:"lat"`
	Lon        float64  `json:"lon"`
	Country    string   `json:"country"`
	State      string   `json:"state,omitempty"`
}

type Coordinates struct {
	Latitude  float64
	Longitude float64
	City      string
}

const COORDINATES_URL = "http://api.openweathermap.org/geo/1.0/direct?q="

func getCoordinates(city string) (Coordinates, error) {

	resp, err := http.Get(COORDINATES_URL + city + "&limit=1" + APIKEY + os.Getenv("OPENWEATHER_API"))
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
		City:      data[0].Name,
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
