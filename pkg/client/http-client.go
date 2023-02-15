package client

import (
	"net/http"
)

type (
	IHttpClient interface {
		Get(uri string) (*http.Response, error)
	}

	HttpClientConfig struct {
		GetMethod func(string) (*http.Response, error)
	}
)

func NewHttpClient() HttpClientConfig {
	return HttpClientConfig{
		GetMethod: func(uri string) (*http.Response, error) {
			return http.Get(uri)
		},
	}
}

func (c HttpClientConfig) Get(uri string) (*http.Response, error) {
	resp, err := c.GetMethod(uri)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
