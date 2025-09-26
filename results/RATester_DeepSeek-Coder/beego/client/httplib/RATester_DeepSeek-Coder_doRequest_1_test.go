package httplib

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestDoRequest_1(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name    string
		request *BeegoHTTPRequest
		wantErr bool
	}

	testCases := []testCase{
		{
			name: "valid request",
			request: &BeegoHTTPRequest{
				url: "http://example.com",
				req: &http.Request{
					Method: "GET",
					URL:    &url.URL{},
				},
			},
			wantErr: false,
		},
		{
			name: "invalid url",
			request: &BeegoHTTPRequest{
				url: ":",
				req: &http.Request{
					Method: "GET",
					URL:    &url.URL{},
				},
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			_, err := tc.request.doRequest(context.Background())
			if (err != nil) != tc.wantErr {
				t.Errorf("doRequest() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
