package passtlsclientcert

import (
	"fmt"
	"testing"
)

func TestSanitize_1(t *testing.T) {
	tests := []struct {
		name string
		cert []byte
		want string
	}{
		{
			name: "Test case 1",
			cert: []byte("-----BEGIN CERTIFICATE-----\nMIIB...\n-----END CERTIFICATE-----"),
			want: "MIIB...",
		},
		{
			name: "Test case 2",
			cert: []byte("-----BEGIN CERTIFICATE-----\nMIIB...\n-----END CERTIFICATE-----"),
			want: "MIIB...",
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

			if got := sanitize(tt.cert); got != tt.want {
				t.Errorf("sanitize() = %v, want %v", got, tt.want)
			}
		})
	}
}
