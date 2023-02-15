package client

import (
	"net/http"
)

type (
	IHttpClient interface {
		Get(uri string) (*http.Response, error)
	}

	HttpClient struct {
		GetMethod func(string) (*http.Response, error)
	}
)

func NewHttpClient() HttpClient {
	return HttpClient{
		GetMethod: func(uri string) (*http.Response, error) {
			return http.Get(uri)
		},
	}
}

func (c HttpClient) Get(uri string) (*http.Response, error) {
	resp, err := c.GetMethod(uri)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
