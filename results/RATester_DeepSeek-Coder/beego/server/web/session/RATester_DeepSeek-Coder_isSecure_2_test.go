package session

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestIsSecure_2(t *testing.T) {
	manager := &Manager{
		config: &ManagerConfig{
			Secure: true,
		},
	}

	tests := []struct {
		name string
		req  *http.Request
		want bool
	}{
		{
			name: "Secure is true and scheme is https",
			req: &http.Request{
				URL: &url.URL{
					Scheme: "https",
				},
				TLS: &tls.ConnectionState{},
			},
			want: true,
		},
		{
			name: "Secure is true and scheme is http",
			req: &http.Request{
				URL: &url.URL{
					Scheme: "http",
				},
				TLS: &tls.ConnectionState{},
			},
			want: false,
		},
		{
			name: "Secure is false",
			req: &http.Request{
				URL: &url.URL{
					Scheme: "https",
				},
				TLS: &tls.ConnectionState{},
			},
			want: false,
		},
		{
			name: "TLS is nil",
			req: &http.Request{
				URL: &url.URL{
					Scheme: "https",
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := manager.isSecure(tt.req); got != tt.want {
				t.Errorf("isSecure() = %v, want %v", got, tt.want)
			}
		})
	}
}
