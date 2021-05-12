package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Payload struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
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
		Type    int    `json:"type"`
		ID      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

const APP_ID string = "140ac99e142ea063e7ca6be8e4e0613f"

var URI string = "https://api.openweathermap.org/data/2.5/weather?zip=%d,us&units=imperial&appid=%s"
var zip = 92126

func main() {
	var uri = fmt.Sprintf(URI, zip, APP_ID)
	log.Printf("Calling weather API: %s", uri)
	res, err := http.Get(uri)
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	body, _ := ioutil.ReadAll(res.Body)
	log.Printf("weather: %s", body)
	// deserialize json result
	var weather Payload
	err = json.Unmarshal(body, &weather)
	if err != nil {
		log.Fatal(err)
	}
	// format and print some fields
	var w = weather.Main
	log.Printf("temperature (F): min: %.1f, max: %.1f, humidity: %v%%", w.TempMin, w.TempMax, w.Humidity)
	var s = weather.Sys
	log.Printf("sunrise: %v, sunset: %v", formatTime(s.Sunrise), formatTime(s.Sunset))
}

func formatTime(ts int) time.Time {
	return time.Unix(int64(ts), 0)
}
