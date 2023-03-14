// Package uri implements routines for manipulating URIs.
// The uri package allows to create a uri by providing both: uri and parameters.
package uri

import (
	"errors"
	"fmt"
)

type Uri struct {
	Path   string
	Params map[string]string
}

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
