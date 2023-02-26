package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"

	"github.com/jlhidalgo/nasa-daily-pic/pkg/client"
	"github.com/jlhidalgo/nasa-daily-pic/pkg/models"
	"github.com/jlhidalgo/nasa-daily-pic/pkg/rhandler"
	"github.com/jlhidalgo/nasa-daily-pic/pkg/server"
)

func homepageFunc(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./web/template/index.html")
	picture, _ := getPicOfDay()

	t.Execute(rw, picture)
}

func main() {
	// TODO: make these arguments configurable
	serv := server.NewServer("", "8081")
	serv.AddHandleFunc("/", homepageFunc)
	serv.Run()
}

func getPicOfDay() (models.Picture, error) {

	picture := models.Picture{}
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
		return models.Picture{}, err
	}

	return picture, nil

}
