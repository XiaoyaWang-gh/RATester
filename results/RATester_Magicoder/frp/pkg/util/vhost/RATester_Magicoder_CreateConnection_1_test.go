package vhost

import (
	"fmt"
	"testing"
)

func TestCreateConnection_1(t *testing.T) {
	rp := &HTTPReverseProxy{
		vhostRouter: &Routers{
			indexByDomain: make(map[string]routerByHTTPUser),
		},
	}

	reqRouteInfo := &RequestRouteInfo{
		URL:        "/test",
		Host:       "example.com",
		HTTPUser:   "user",
		RemoteAddr: "127.0.0.1:8080",
		URLHost:    "example.com",
		Endpoint:   "endpoint",
	}

	tests := []struct {
		name       string
		byEndpoint bool
		wantErr    bool
	}{
		{
			name:       "Test case 1",
			byEndpoint: true,
			wantErr:    false,
		},
		{
			name:       "Test case 2",
			byEndpoint: false,
			wantErr:    false,
		},
		{
			name:       "Test case 3",
			byEndpoint: true,
			wantErr:    true,
		},
		{
			name:       "Test case 4",
			byEndpoint: false,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			_, err := rp.CreateConnection(reqRouteInfo, tt.byEndpoint)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateConnection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
