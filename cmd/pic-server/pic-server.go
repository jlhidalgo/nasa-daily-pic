package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jlhidalgo/nasa-daily-pic/pkg/client"
	"github.com/jlhidalgo/nasa-daily-pic/pkg/rhandler"
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

func homepageFunc(rw http.ResponseWriter, r *http.Request) {
	http.ServeFile(rw, r, "./web/template/index.html")
}

func main() {
	picture, err := getPicOfDay()
	fmt.Println(picture, err)

	http.HandleFunc("/", homepageFunc)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func getPicOfDay() (Picture, error) {

	picture := Picture{}
	uri := "https://api.nasa.gov/planetary/apod"
	params := map[string]string{"api_key": "DEMO_KEY"}

	httpClient := client.NewHttpClient()
	restHandler := rhandler.NewRestHandler(httpClient)

	body, err := restHandler.Get(uri, params)
	if err != nil {
		fmt.Println("There was an error:", err)
	}

	err = json.Unmarshal(body, &picture)
	if err != nil {
		return Picture{}, err
	}

	return picture, nil

}
