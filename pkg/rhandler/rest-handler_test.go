package rhandler

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	mock_client "github.com/jlhidalgo/nasa-daily-pic/pkg/client/mock"
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

func Test_Get(t *testing.T) {
	ctrl := gomock.NewController(t)

	tests := []struct {
		name             string
		httpClientResult func(string) (*http.Response, error)
		expectedError    bool
		expectedResult   []byte
	}{
		{
			name: "Error in HTTP Client",
			httpClientResult: func(uri string) (*http.Response, error) {
				return nil, errors.New("Boom!")
			},
			expectedError: true,
		},
		{
			name: "Status code not 200",
			httpClientResult: func(uri string) (*http.Response, error) {
				return &http.Response{
					Status:     "400 Bad Request",
					StatusCode: 400,
				}, nil
			},
			expectedError: true,
		},
		{
			name: "Successful response",
			httpClientResult: func(uri string) (*http.Response, error) {
				resp := generateSuccessfulResponse()
				return resp, nil
			},
			expectedError: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			hclientMock := mock_client.NewMockIHttpClient(ctrl)
			hclientMock.EXPECT().Get(gomock.Any()).Return(tc.httpClientResult(""))

			rhandler := NewRestHandler(hclientMock)
			uri := &Uri{
				"http://localhost",
				map[string]string{},
			}
			resp, err := rhandler.Get(uri.GetUri())
			if tc.expectedError {
				assert.NotNil(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NotNil(t, resp)
				assert.Nil(t, err)
			}
		})
	}

}

func generateSuccessfulResponse() *http.Response {
	body := "Hello world!"
	resp := &http.Response{
		Status:        "200 OK",
		StatusCode:    200,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Body:          ioutil.NopCloser(bytes.NewBufferString(body)),
		ContentLength: int64(len(body)),
		Header:        make(http.Header, 0),
	}
	return resp
}
