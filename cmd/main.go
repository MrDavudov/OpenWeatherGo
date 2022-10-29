package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/MrDavudov/OpenWeatherGO/internal/app/model"
)

const base = "http://api.openweathermap.org"
const pathCity = "/geo/1.0/direct?"
const pathWeather = "/data/2.5/forecast?"
const apiKeys = "&appid=90f2edc318c106c65581f4052ad16c6f"

type City struct {
	Name    string  `json:"name"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
	Country string  `json:"country"`
}

type DataTemp struct {
	List []struct {
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
		Data string `json:"dt_txt"`
	} `json:"list"`
}

func Response(city string) {
	resp, err := http.Get(base + pathCity + "q=" + city + apiKeys)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if string(body) == "[]" {
		log.Fatal(fmt.Errorf("no such city"))
	}

	objCity := []City{}
	err = json.Unmarshal(body, &objCity)
	if err != nil {
		log.Fatal(err)
	}

	Weather := model.Weather{
		Name:		objCity[0].Name,
		Lat:		objCity[0].Lat,
		Lon:		objCity[0].Lon,
		Country:	objCity[0].Country,
	}

	lat := fmt.Sprintf("lat=%f", Weather.Lat)
	lon := fmt.Sprintf("&lon=%f", Weather.Lon)

	resp, err = http.Get(base + pathWeather + lat + lon + "&units=metric" + apiKeys)
	if err != nil {
		log.Fatal(err)
	}

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	objTemp := DataTemp{}
	err = json.Unmarshal(body, &objTemp)
	if err != nil {
		log.Fatal(err)
	}

	for i := range objTemp.List {
		if strings.Contains(objTemp.List[i].Data, "12:00") {
			d := model.DtTemp {
				Dt: objTemp.List[i].Data,
				Temp: objTemp.List[i].Main.Temp,
			}
			Weather.DtTemp = append(Weather.DtTemp, d)
		}
	}

	// fmt.Println(Weather)
}

func main() {
	Response("London")

}