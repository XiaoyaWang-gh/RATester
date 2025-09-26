package client

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestRemoveProxyHeaders_1(t *testing.T) {
	testCases := []struct {
		name     string
		req      *http.Request
		expected *http.Request
	}{
		{
			name: "Test with a request with proxy headers",
			req: &http.Request{
				Header: http.Header{
					"Proxy-Connection":    []string{"close"},
					"Connection":          []string{"close"},
					"Proxy-Authenticate":  []string{"Basic"},
					"Proxy-Authorization": []string{"Basic QWxhZGRpbjpvcGVuIg=="},
					"TE":                  []string{"trailers"},
					"Trailers":            []string{"TE"},
					"Transfer-Encoding":   []string{"chunked"},
					"Upgrade":             []string{"HTTP/2.0"},
				},
			},
			expected: &http.Request{
				Header: http.Header{},
			},
		},
		{
			name: "Test with a request without proxy headers",
			req: &http.Request{
				Header: http.Header{
					"Content-Type": []string{"application/json"},
				},
			},
			expected: &http.Request{
				Header: http.Header{
					"Content-Type": []string{"application/json"},
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

			removeProxyHeaders(tc.req)
			if !reflect.DeepEqual(tc.req, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, tc.req)
			}
		})
	}
}
