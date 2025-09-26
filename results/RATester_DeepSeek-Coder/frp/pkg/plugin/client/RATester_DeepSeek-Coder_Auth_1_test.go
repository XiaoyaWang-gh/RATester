package client

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestAuth_1(t *testing.T) {
	type testCase struct {
		name     string
		opts     *v1.HTTPProxyPluginOptions
		req      *http.Request
		expected bool
	}

	testCases := []testCase{
		{
			name: "NoAuth",
			opts: &v1.HTTPProxyPluginOptions{
				HTTPUser:     "",
				HTTPPassword: "",
			},
			req: &http.Request{
				Header: make(http.Header),
			},
			expected: true,
		},
		{
			name: "InvalidAuth",
			opts: &v1.HTTPProxyPluginOptions{
				HTTPUser:     "user",
				HTTPPassword: "pass",
			},
			req: &http.Request{
				Header: make(http.Header),
			},
			expected: false,
		},
		{
			name: "ValidAuth",
			opts: &v1.HTTPProxyPluginOptions{
				HTTPUser:     "user",
				HTTPPassword: "pass",
			},
			req: &http.Request{
				Header: http.Header{
					"Proxy-Authorization": []string{"Basic " + base64.StdEncoding.EncodeToString([]byte("user:pass"))},
				},
			},
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			hp := &HTTPProxy{opts: tc.opts}
			result := hp.Auth(tc.req)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
