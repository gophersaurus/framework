package cityweather

import (
	"encoding/json"

	"git.target.com/gophersaurus/framework"
)

func Find(city, country string) (*Result, error) {

	url := "http://api.openweathermap.org/data/2.5/weather?q=" + city + "," + country

	req, err := gf.Client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	bytes, err := gf.Client.Send(req)
	if err != nil {
		return nil, err
	}

	result := &Result{}
	err = json.Unmarshal(bytes, result)
	return result, err

}

type Result struct {
	Coord   `json:"coord"`
	Sys     `json:"sys"`
	Weather []Data `json:"weather"`
	Base    string `json:"base"`
	Main    `json:"main"`
	Wind    `json:"wind"`
	Clouds  `json:"clouds"`
	Dt      int    `json:"dt"`
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Cod     int    `json:"cod"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Sys struct {
	Type    int     `json:"type"`
	Id      int     `json:"id"`
	Message float64 `json:"message"`
	Country string  `json:"country"`
	Sunrise int     `json:"sunrise"`
	Sunset  int     `json:"sunset"`
}

type Data struct {
	Id          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Main struct {
	Temp     float64 `json:"temp"`
	Pressure int     `json:"pressure"`
	Humidity int     `json:"humidity"`
	TempMin  float64 `json:"temp_min"`
	TempMax  float64 `json:"temp_max"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
}

type Clouds struct {
	All int `json:"all"`
}
