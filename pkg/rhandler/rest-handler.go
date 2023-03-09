package rhandler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/jlhidalgo/nasa-daily-pic/pkg/client"
)

type (
	RestHandler struct {
		HttpClient client.IHttpClient
	}
)

func NewRestHandler(hClient client.IHttpClient) RestHandler {
	return RestHandler{
		HttpClient: hClient,
	}
}

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

func getBody(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}
