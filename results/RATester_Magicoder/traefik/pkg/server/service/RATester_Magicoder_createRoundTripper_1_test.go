package service

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestCreateRoundTripper_1(t *testing.T) {
	tests := []struct {
		name    string
		cfg     *dynamic.ServersTransport
		wantErr bool
	}{
		{
			name:    "nil config",
			cfg:     nil,
			wantErr: true,
		},
		{
			name: "valid config",
			cfg: &dynamic.ServersTransport{
				ServerName: "example.com",
			},
			wantErr: false,
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			r := &RoundTripperManager{}
			_, err := r.createRoundTripper(tt.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("createRoundTripper() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
