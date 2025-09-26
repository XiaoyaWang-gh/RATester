package ingress

import (
	"fmt"
	"testing"
)

func TestNewInClusterClient_1(t *testing.T) {
	tests := []struct {
		name      string
		endpoint  string
		wantErr   bool
		wantPanic bool
	}{
		{
			name:     "Test with empty endpoint",
			endpoint: "",
			wantErr:  false,
		},
		{
			name:     "Test with non-empty endpoint",
			endpoint: "http://localhost:8080",
			wantErr:  false,
		},
		{
			name:     "Test with invalid endpoint",
			endpoint: "http://localhost:8080/api",
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

			defer func() {
				if r := recover(); r != nil {
					if !tt.wantPanic {
						t.Errorf("newInClusterClient() = %v, wantPanic %v", r, tt.wantPanic)
					}
				}
			}()

			_, err := newInClusterClient(tt.endpoint)
			if (err != nil) != tt.wantErr {
				t.Errorf("newInClusterClient() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
