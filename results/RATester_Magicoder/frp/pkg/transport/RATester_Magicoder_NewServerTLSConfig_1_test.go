package transport

import (
	"fmt"
	"testing"
)

func TestNewServerTLSConfig_1(t *testing.T) {
	tests := []struct {
		name     string
		certPath string
		keyPath  string
		caPath   string
		wantErr  bool
	}{
		{
			name:     "Test case 1",
			certPath: "path/to/cert",
			keyPath:  "path/to/key",
			caPath:   "path/to/ca",
			wantErr:  false,
		},
		{
			name:     "Test case 2",
			certPath: "",
			keyPath:  "",
			caPath:   "",
			wantErr:  false,
		},
		{
			name:     "Test case 3",
			certPath: "path/to/cert",
			keyPath:  "path/to/key",
			caPath:   "",
			wantErr:  false,
		},
		{
			name:     "Test case 4",
			certPath: "",
			keyPath:  "",
			caPath:   "path/to/ca",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			_, err := NewServerTLSConfig(tt.certPath, tt.keyPath, tt.caPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewServerTLSConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
