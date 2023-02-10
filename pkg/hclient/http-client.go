package hclient

import (
	"fmt"
	"io"
	"net/http"
)

type IClient interface {
	Get(uri string, params map[string]string) ([]byte, error)
}

func Get(uri string, params map[string]string) ([]byte, error) {
	queryString := getQueryString(params)
	uri = fmt.Sprintf("%s?%s", uri, queryString)

	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response failed with status code: %v", resp.StatusCode)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func getQueryString(params map[string]string) string {
	queryString := ""
	for key, value := range params {
		if len(queryString) > 0 {
			queryString += "&"
		}
		queryString = fmt.Sprintf("%s%s=%s", queryString, key, value)
	}
	return queryString
}
