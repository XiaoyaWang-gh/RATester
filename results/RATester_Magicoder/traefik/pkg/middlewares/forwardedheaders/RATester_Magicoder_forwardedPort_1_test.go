package forwardedheaders

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"testing"
)

func TestForwardedPort_1(t *testing.T) {
	tests := []struct {
		name string
		req  *http.Request
		want string
	}{
		{
			name: "nil request",
			req:  nil,
			want: "",
		},
		{
			name: "host with port",
			req: &http.Request{
				Host: "example.com:8080",
			},
			want: "8080",
		},
		{
			name: "x-forwarded-proto https",
			req: &http.Request{
				Header: http.Header{
					"X-Forwarded-Proto": []string{"https"},
				},
			},
			want: "443",
		},
		{
			name: "x-forwarded-proto wss",
			req: &http.Request{
				Header: http.Header{
					"X-Forwarded-Proto": []string{"wss"},
				},
			},
			want: "443",
		},
		{
			name: "tls",
			req: &http.Request{
				TLS: &tls.ConnectionState{},
			},
			want: "443",
		},
		{
			name: "default",
			req:  &http.Request{},
			want: "80",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := forwardedPort(tt.req); got != tt.want {
				t.Errorf("forwardedPort() = %v, want %v", got, tt.want)
			}
		})
	}
}
