package openweather

import (
	"encoding/json"
	"github.com/dsoloview/gobot/pkg/helpers"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	WEATHER_URL = "https://api.openweathermap.org/data/2.5/weather?"
	LATITUDE    = "lat="
	LONGITUDE   = "&lon="
	APIKEY      = "&appid="
	LANG        = "&lang=ru"
	UNITS       = "&units=metric"
)

type WeatherResponse struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		Id          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int     `json:"type"`
		Id      int     `json:"id"`
		Message float64 `json:"message"`
		Country string  `json:"country"`
		Sunrise int     `json:"sunrise"`
		Sunset  int     `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

type Weather struct {
	Temperature float64
	Description string
	FeelsLike   float64
	City        Coordinates
}

func GetWeather(location string) (*Weather, error) {

	coordinates, err := getCoordinates(location)
	if err != nil {
		return nil, err
	}
	url := WEATHER_URL + LATITUDE + helpers.Floattostr(coordinates.Latitude) + LONGITUDE + helpers.Floattostr(coordinates.Longitude) + UNITS + LANG + APIKEY + os.Getenv("OPENWEATHER_API")

	resp := makeGetRequest(url)

	weatherResponse := getWeatherJson(&resp)

	return &Weather{
		Temperature: weatherResponse.Main.Temp,
		Description: weatherResponse.Weather[0].Description,
		FeelsLike:   weatherResponse.Main.FeelsLike,
		City:        coordinates,
	}, nil
}

func getWeatherJson(resp *http.Response) WeatherResponse {
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

	var data WeatherResponse

	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println(err)
	}

	return data
}

func makeGetRequest(url string) http.Response {

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	return *resp
}
