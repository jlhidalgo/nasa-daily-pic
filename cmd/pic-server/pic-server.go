package main

import (
	"encoding/json"
	"fmt"

	"github.com/jlhidalgo/nasa-daily-pic/pkg/hclient"
)

type (
	Picture struct {
		Copyright      string `json:"copyright,omitempty"`
		Date           string `json:"date,omitempty"`
		Explanation    string `json:"explanation,omitempty"`
		Hdurl          string `json:"hdurl,omitempty"`
		MediaType      string `json:"media_type,omitempty"`
		ServiceVersion string `json:"service_version,omitempty"`
		Title          string `json:"title,omitempty"`
		Url            string `json:"url,omitempty"`
	}
)

func main() {
	picture, err := getPicOfDay()
	fmt.Println(picture, err)
}

func getPicOfDay() (Picture, error) {

	picture := Picture{}
	uri := "https://api.nasa.gov/planetary/apod"
	params := map[string]string{"api_key": "DEMO_KEY"}

	body, err := hclient.Get(uri, params)
	if err != nil {
		fmt.Println("There was an error:", err)
	}

	err = json.Unmarshal(body, &picture)
	if err != nil {
		return Picture{}, err
	}

	return picture, nil

}
