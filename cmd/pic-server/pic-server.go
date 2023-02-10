package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
	resp, err := http.Get("https://api.nasa.gov/planetary/apod?api_key=DEMO_KEY")
	if err != nil {
		return picture, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return Picture{}, err
	}

	err = json.Unmarshal(body, &picture)
	if err != nil {
		return Picture{}, err
	}

	return picture, nil

}
