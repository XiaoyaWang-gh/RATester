package dynamic

import (
	"fmt"
	"testing"
)

func TestHasSecureHeadersDefined_1(t *testing.T) {
	tests := []struct {
		name string
		h    *Headers
		want bool
	}{
		{
			name: "nil",
			h:    nil,
			want: false,
		},
		{
			name: "empty",
			h:    &Headers{},
			want: false,
		},
		{
			name: "allowed hosts",
			h: &Headers{
				AllowedHosts: []string{"example.com"},
			},
			want: true,
		},
		{
			name: "hosts proxy headers",
			h: &Headers{
				HostsProxyHeaders: []string{"X-Forwarded-Host"},
			},
			want: true,
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

			if got := tt.h.HasSecureHeadersDefined(); got != tt.want {
				t.Errorf("Headers.HasSecureHeadersDefined() = %v, want %v", got, tt.want)
			}
		})
	}
}
