package rhandler

import "fmt"

type Uri struct {
	Uri    string
	Params map[string]string
}

func (u *Uri) GetUri() string {
	queryString := getQueryString(u.Params)
	return fmt.Sprintf("%s?%s", u.Uri, queryString)
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
