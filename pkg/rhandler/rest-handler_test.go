package rhandler

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	mock_client "github.com/jlhidalgo/nasa-daily-pic/pkg/client/mock"
	"github.com/jlhidalgo/nasa-daily-pic/pkg/utils"
	"github.com/stretchr/testify/assert"
)

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
			uri := &utils.Uri{
				Uri:    "http://localhost",
				Params: map[string]string{},
			}
			resUri, _ := uri.GetUri()
			resp, err := rhandler.Get(resUri)
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
