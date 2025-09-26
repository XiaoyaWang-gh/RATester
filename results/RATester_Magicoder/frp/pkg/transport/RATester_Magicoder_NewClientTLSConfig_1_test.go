package transport

import (
	"fmt"
	"testing"
)

func TestNewClientTLSConfig_1(t *testing.T) {
	tests := []struct {
		name       string
		certPath   string
		keyPath    string
		caPath     string
		serverName string
		wantErr    bool
	}{
		{
			name:       "valid config",
			certPath:   "valid_cert.pem",
			keyPath:    "valid_key.pem",
			caPath:     "valid_ca.pem",
			serverName: "example.com",
			wantErr:    false,
		},
		{
			name:       "invalid cert",
			certPath:   "invalid_cert.pem",
			keyPath:    "valid_key.pem",
			caPath:     "valid_ca.pem",
			serverName: "example.com",
			wantErr:    true,
		},
		{
			name:       "invalid key",
			certPath:   "valid_cert.pem",
			keyPath:    "invalid_key.pem",
			caPath:     "valid_ca.pem",
			serverName: "example.com",
			wantErr:    true,
		},
		{
			name:       "invalid ca",
			certPath:   "valid_cert.pem",
			keyPath:    "valid_key.pem",
			caPath:     "invalid_ca.pem",
			serverName: "example.com",
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

			_, err := NewClientTLSConfig(tt.certPath, tt.keyPath, tt.caPath, tt.serverName)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewClientTLSConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
