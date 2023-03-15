// Package rhandler implements functionality for processing HTTP client requests.
// In other words this is to consume a REST API.
package rhandler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/jlhidalgo/nasa-daily-pic/pkg/client"
)

type (
	// RestHandler implements methods for consuming REST APIs.
	// HttpClient takes an interface so different clients can be used
	// interchangeably.
	RestHandler struct {
		HttpClient client.IHttpClient
	}
)

// NewRestHandler creates a RestHandler and specifies the HTTP client to be used.
// The net/http client from the Go Standard Library will be usually utilized here,
// but this current implementation allows to replace it with a mock for
// testing purposes.
func NewRestHandler(hClient client.IHttpClient) RestHandler {
	return RestHandler{
		HttpClient: hClient,
	}
}

// Get implements the GET method for consuming a REST API.
// Returns the body if the response from the client is successful,
// otherwise it returns an error.
func (r RestHandler) Get(uri string) ([]byte, error) {
	resp, err := r.HttpClient.Get(uri)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response failed with status code: %v", resp.StatusCode)
	}

	return getBody(resp)
}

// getBody retrieves the body from the response by using a reader.
// Returns the body as a slice of bytes or an error if the body
// cannot be retrieved.
func getBody(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}
