package client

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestRemoveProxyHeaders_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	req := &http.Request{
		Method: "GET",
		URL: &url.URL{
			Scheme: "http",
			Host:   "example.com",
			Path:   "/path",
		},
		Header: http.Header{
			"Proxy-Connection":    []string{"close"},
			"Connection":          []string{"keep-alive"},
			"Proxy-Authenticate":  []string{"Basic"},
			"Proxy-Authorization": []string{"Basic dXNlcjpwYXNzd29yZA=="},
			"TE":                  []string{"trailers"},
			"Trailers":            []string{"TE"},
			"Transfer-Encoding":   []string{"chunked"},
			"Upgrade":             []string{"HTTP/2.0"},
		},
	}

	removeProxyHeaders(req)

	expectedHeaders := http.Header{
		"Proxy-Connection":    []string{},
		"Connection":          []string{},
		"Proxy-Authenticate":  []string{},
		"Proxy-Authorization": []string{},
		"TE":                  []string{},
		"Trailers":            []string{},
		"Transfer-Encoding":   []string{},
		"Upgrade":             []string{},
	}

	if !reflect.DeepEqual(req.Header, expectedHeaders) {
		t.Errorf("Expected headers to be removed, but they are not.")
	}
}
