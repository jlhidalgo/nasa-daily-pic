package hclient

import (
	"net/http"
)

func Get(uri string) (*http.Response, error) {
	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
