package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"

	"github.com/jlhidalgo/nasa-daily-pic/pkg/client"
	"github.com/jlhidalgo/nasa-daily-pic/pkg/rhandler"
	"github.com/jlhidalgo/nasa-daily-pic/pkg/server"
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
	t, _ := template.ParseFiles("./web/template/index.html")
	picture, _ := getPicOfDay()

	t.Execute(rw, picture)
}

func main() {
	http.HandleFunc("/", homepageFunc)
	serv := server.NewServer("", "8081")
	serv.Run()
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
