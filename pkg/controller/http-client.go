package controller

import (
	"net/http"
)

type (
	IHttpController interface {
		Get(uri string) (*http.Response, error)
	}

	HttpConfig struct {
		GetMethod func(string) (*http.Response, error)
	}
)

func NewHttpClient() HttpConfig {
	return HttpConfig{
		GetMethod: func(uri string) (*http.Response, error) {
			return http.Get(uri)
		},
	}
}

func (c HttpConfig) Get(uri string) (*http.Response, error) {
	resp, err := c.GetMethod(uri)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
