package echo

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestSetRequest_1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		req  *http.Request
	}{
		{
			name: "Nil Request",
			req:  nil,
		},
		{
			name: "Non-Nil Request",
			req: &http.Request{
				Method: "GET",
				URL: &url.URL{
					Scheme: "http",
					Host:   "example.com",
					Path:   "/path",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &context{}
			c.SetRequest(tc.req)

			if c.request != tc.req {
				t.Errorf("Expected request to be %v, got %v", tc.req, c.request)
			}
		})
	}
}
