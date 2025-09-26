package server

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestEncodeQuerySemicolons_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.RawQuery != "a=b;c=d" {
			t.Errorf("expected req.URL.RawQuery to be 'a=b;c=d', got %q", req.URL.RawQuery)
		}
	})
	encodeQuerySemicolons(h).ServeHTTP(nil, &http.Request{
		URL: &url.URL{
			RawQuery: "a=b;c=d",
		},
	})
}
