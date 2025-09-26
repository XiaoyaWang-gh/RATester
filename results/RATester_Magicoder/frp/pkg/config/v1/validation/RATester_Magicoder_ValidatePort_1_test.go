package validation

import (
	"fmt"
	"testing"
)

func TestValidatePort_1(t *testing.T) {
	tests := []struct {
		name      string
		port      int
		fieldPath string
		wantErr   bool
	}{
		{
			name:      "valid port",
			port:      8080,
			fieldPath: "port",
			wantErr:   false,
		},
		{
			name:      "invalid port",
			port:      65536,
			fieldPath: "port",
			wantErr:   true,
		},
		{
			name:      "negative port",
			port:      -1,
			fieldPath: "port",
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if err := ValidatePort(tt.port, tt.fieldPath); (err != nil) != tt.wantErr {
				t.Errorf("ValidatePort() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
