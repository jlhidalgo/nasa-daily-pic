// package client allows sending requests to APIs, it's an HTTP client.
package client

import (
	"net/http"
)

type (
	// IHttpClient is implemented by any value that has a Get method.
	IHttpClient interface {
		Get(uri string) (*http.Response, error)
	}

	// HttpClient is a type that implements IHttpClient interface
	HttpClient struct {
		GetMethod func(string) (*http.Response, error)
	}
)

// NewHttpClient returns a new value of HttpClient type.
// The GetMethod function is initialized to return the result
// of http.Get which is a direct invocation to the standard
// library implementation for an HTTP client
func NewHttpClient() HttpClient {
	return HttpClient{
		GetMethod: func(uri string) (*http.Response, error) {
			return http.Get(uri)
		},
	}
}

// Get actually performs the execution of the http.get method
// implemented in the standard library.
func (c HttpClient) Get(uri string) (*http.Response, error) {
	resp, err := c.GetMethod(uri)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
