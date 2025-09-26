package tailscale

import (
	"fmt"
	"testing"
)

func TestIsTailscaleDomain_1(t *testing.T) {

	tests := []struct {
		name     string
		domain   string
		expected bool
	}{
		{
			name:     "valid tailscale domain",
			domain:   "test.ts.net",
			expected: true,
		},
		{
			name:     "invalid tailscale domain",
			domain:   "test.ts.com",
			expected: false,
		},
		{
			name:     "invalid tailscale domain",
			domain:   "test.ts.net.com",
			expected: false,
		},
		{
			name:     "invalid tailscale domain",
			domain:   "test.ts.net.",
			expected: false,
		},
		{
			name:     "invalid tailscale domain",
			domain:   ".ts.net",
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := isTailscaleDomain(test.domain)
			if result != test.expected {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}
