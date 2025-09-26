package session

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestisSecure_2(t *testing.T) {
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
			name: "https",
			req: &http.Request{
				URL: &url.URL{
					Scheme: "https",
				},
			},
			want: true,
		},
		{
			name: "http",
			req: &http.Request{
				URL: &url.URL{
					Scheme: "http",
				},
			},
			want: false,
		},
		{
			name: "https with TLS",
			req: &http.Request{
				TLS: &tls.ConnectionState{},
			},
			want: true,
		},
		{
			name: "http without TLS",
			req:  &http.Request{},
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
