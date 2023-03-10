package rhandler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getQueryString(t *testing.T) {
	tests := []struct {
		name     string
		params   map[string]string
		expected string
	}{
		{
			name:     "Test empty params",
			params:   map[string]string{},
			expected: "",
		},
		{
			name: "Test one-element map",
			params: map[string]string{
				"key": "myKey",
			},
			expected: "key=myKey",
		},
		{
			name: "Test 2-element map",
			params: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			expected: "key1=value1&key2=value2",
		},
		{
			name: "Test empty value gets discarded",
			params: map[string]string{
				"key1": "value1",
				"key2": "",
			},
			expected: "key1=value1",
		},
		{
			name: "Test all values are empty",
			params: map[string]string{
				"key1": "",
				"key2": "",
			},
			expected: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := getQueryString(test.params)
			if actual != test.expected {
				t.Errorf("failed to get querystring, expected: %s, current: %s ", test.expected, actual)
			}
		})
	}
}

func TestGetUri(t *testing.T) {
	tests := []struct {
		name       string
		url        string
		parameters map[string]string
		expError   bool
		expOutput  string
	}{
		{
			name:      "Successful w/params",
			url:       "http://localhost",
			expError:  false,
			expOutput: "http://localhost",
		},
		{
			name:       "Successful w/empty params",
			url:        "http://localhost",
			parameters: map[string]string{},
			expError:   false,
			expOutput:  "http://localhost",
		},
		{
			name: "Successful w/one param",
			url:  "http://localhost",
			parameters: map[string]string{
				"key": "value",
			},
			expError:  false,
			expOutput: "http://localhost?key=value",
		},
		{
			name: "Successful w/several params",
			url:  "http://localhost",
			parameters: map[string]string{
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
			},
			expError:  false,
			expOutput: "http://localhost?key1=value1&key2=value2&key3=value3",
		},
		{
			name:     "Error missing uri",
			expError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			uri := &Uri{
				Uri:    tc.url,
				Params: tc.parameters,
			}
			output, err := uri.GetUri()

			if tc.expError {
				assert.NotNil(t, err)
				assert.Equal(t, "", output)
			} else {
				assert.Equal(t, tc.expOutput, output)
				assert.Nil(t, err)
			}
		})
	}
}
