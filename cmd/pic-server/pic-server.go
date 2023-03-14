package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"

	"github.com/jlhidalgo/nasa-daily-pic/configs"
	"github.com/jlhidalgo/nasa-daily-pic/pkg/client"
	"github.com/jlhidalgo/nasa-daily-pic/pkg/models"
	"github.com/jlhidalgo/nasa-daily-pic/pkg/rhandler"
	"github.com/jlhidalgo/nasa-daily-pic/pkg/server"
	"github.com/jlhidalgo/nasa-daily-pic/pkg/uri"
)

func homepageFunc(w http.ResponseWriter, r *http.Request) {
	picture, err := getPicOfDay()

	if err != nil {
		// TODO: actually this should return an error in HTML
		fmt.Println("failed to obtain the picture of the day:", err)
		generateInternalErrorPage(w)
		return
	}

	t, err := template.ParseFiles("./web/template/index.html")
	if err != nil {
		// TODO: maybe parse and open different page for internal error
		fmt.Println("failed to parse file: ", err)
		generateInternalErrorPage(w)
		return
	}

	t.Execute(w, picture)
}

func main() {
	serv := server.NewServer(configs.SERVER_HOSTNAME, configs.SERVER_PORT)
	serv.AddHandleFunc("/", homepageFunc)
	serv.Run()
}

func getPicOfDay() (models.Picture, error) {
	// TODO: move this code out from this function
	picture := models.Picture{}
	httpClient := client.NewHttpClient()
	restHandler := rhandler.NewRestHandler(httpClient)

	uri := &uri.Uri{
		Path:   configs.CLIENT_APOD_URI,
		Params: configs.CLIENT_APOD_PARAMS,
	}

	resUri, _ := uri.GetUri()
	body, err := restHandler.Get(resUri)
	if err != nil {
		return models.Picture{}, err
	}

	err = json.Unmarshal(body, &picture)
	if err != nil {
		return models.Picture{}, err
	}

	return picture, nil

}

func generateInternalErrorPage(w http.ResponseWriter) {
	w.WriteHeader(500)
	w.Write([]byte("<html><body>Internal error!</body></html>"))
}
