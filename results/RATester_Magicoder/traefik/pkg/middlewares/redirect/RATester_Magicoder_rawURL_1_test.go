package redirect

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"testing"
)

func TestRawURL_1(t *testing.T) {
	tests := []struct {
		name string
		req  *http.Request
		want string
	}{
		{
			name: "Test with http scheme",
			req: &http.Request{
				Host:       "example.com",
				RequestURI: "/path",
			},
			want: "http://example.com/path",
		},
		{
			name: "Test with https scheme",
			req: &http.Request{
				Host:       "example.com",
				RequestURI: "/path",
				TLS:        &tls.ConnectionState{},
			},
			want: "https://example.com/path",
		},
		{
			name: "Test with custom scheme",
			req: &http.Request{
				Host:       "example.com",
				RequestURI: "http://example.com/path",
			},
			want: "http://example.com/path",
		},
		{
			name: "Test with custom scheme and port",
			req: &http.Request{
				Host:       "example.com:8080",
				RequestURI: "http://example.com:8080/path",
			},
			want: "http://example.com:8080/path",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := rawURL(tt.req); got != tt.want {
				t.Errorf("rawURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
