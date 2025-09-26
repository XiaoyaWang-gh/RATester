package context

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestScheme_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	input := &BeegoInput{
		Context: &Context{
			Request: &http.Request{
				Header: http.Header{
					"X-Forwarded-Proto": []string{"https"},
				},
				URL: &url.URL{
					Scheme: "http",
				},
				TLS: &tls.ConnectionState{},
			},
		},
	}

	scheme := input.Scheme()
	if scheme != "https" {
		t.Errorf("Expected scheme 'https', got '%s'", scheme)
	}
}
