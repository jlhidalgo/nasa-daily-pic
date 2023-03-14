// Package uri implements routines for validating and generating URLs.
// The uri package is not really necessary but it was added for educational purposes,
// especially to learn more about unit testing in Go.
package uri

import (
	"errors"
	"fmt"
)

// Uri is the representation of the parts of a URL.
// Path is a simplified version of schema + path, i.e. https://domain.com/.
// Params are the key-value pairs that will be used in the QueryString.
type Uri struct {
	Path   string
	Params map[string]string
}

// GetUri returns a concatenation of Path and query string parameters.
// If Path is empty then returns an error.
func (u *Uri) GetUri() (string, error) {
	if len(u.Path) <= 0 {
		return "", errors.New("failed to get the uri, property Uri is empty")
	}

	queryString := getQueryString(u.Params)

	if len(queryString) > 0 {
		return fmt.Sprintf("%s?%s", u.Path, queryString), nil
	}

	return u.Path, nil
}

// getQueryString it concatenates key-value pairs obtained from the params map
// in order to generate the query string to be used in the URL.
func getQueryString(params map[string]string) string {
	queryString := ""
	for key, value := range params {
		if len(value) > 0 {
			if len(queryString) > 0 {
				queryString += "&"
			}
			queryString = fmt.Sprintf("%s%s=%s", queryString, key, value)
		}

	}
	return queryString
}
