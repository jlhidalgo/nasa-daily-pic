package server

import (
	"fmt"
	"log"
	"net/http"
)

type ServerConf struct {
	Hostname string
	Port     string
}

func NewServer(host string, port string) ServerConf {
	return ServerConf{
		Hostname: host,
		Port:     port,
	}
}

func (s ServerConf) AddHandleFunc(pattern string, function func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(pattern, function)
}

func (s ServerConf) Run() {
	addr := s.Hostname + ":" + s.Port
	fmt.Println("Server is running at", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
