package uri

import (
	"errors"
	"fmt"
)

type Uri struct {
	Uri    string
	Params map[string]string
}

func (u *Uri) GetUri() (string, error) {
	if len(u.Uri) <= 0 {
		return "", errors.New("failed to get the uri, property Uri is empty")
	}

	queryString := getQueryString(u.Params)

	if len(queryString) > 0 {
		return fmt.Sprintf("%s?%s", u.Uri, queryString), nil
	}

	return u.Uri, nil
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
