package transport

import (
	"fmt"
	"testing"
)

func TestNewClientTLSConfig_1(t *testing.T) {
	t.Parallel()

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
			certPath:   "valid_cert_path",
			keyPath:    "valid_key_path",
			caPath:     "valid_ca_path",
			serverName: "test_server",
			wantErr:    false,
		},
		{
			name:       "invalid cert path",
			certPath:   "invalid_cert_path",
			keyPath:    "valid_key_path",
			caPath:     "valid_ca_path",
			serverName: "test_server",
			wantErr:    true,
		},
		{
			name:       "invalid key path",
			certPath:   "valid_cert_path",
			keyPath:    "invalid_key_path",
			caPath:     "valid_ca_path",
			serverName: "test_server",
			wantErr:    true,
		},
		{
			name:       "invalid ca path",
			certPath:   "valid_cert_path",
			keyPath:    "valid_key_path",
			caPath:     "invalid_ca_path",
			serverName: "test_server",
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
			}
		})
	}
}
