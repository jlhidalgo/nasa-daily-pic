// Package server implements methods for adding handler functions
// and launching the listener.
package server

import (
	"fmt"
	"log"
	"net/http"
)

// ServerConf contains the hostname and the port number where
// server's listener will be operating.
type ServerConf struct {
	Hostname string
	Port     string
}

// NewServer creates and return a new variable of type ServerConf
func NewServer(host string, port string) ServerConf {
	return ServerConf{
		Hostname: host,
		Port:     port,
	}
}

// AddHandleFunc registers a handler function for the server, where
// pattern contains the path that will be associated to the action
// that will be performed by function.
func (s ServerConf) AddHandleFunc(pattern string, function func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(pattern, function)
}

// Run launches the server by using the hostname and port
// that are defined in ServerConf
func (s ServerConf) Run() {
	addr := s.Hostname + ":" + s.Port
	fmt.Println("Server is running at", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
