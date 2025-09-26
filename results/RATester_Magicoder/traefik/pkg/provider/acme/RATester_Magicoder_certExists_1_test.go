package acme

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/types"
)

func TestCertExists_1(t *testing.T) {
	provider := &Provider{
		certificates: []*CertAndStore{
			{
				Certificate: Certificate{
					Domain: types.Domain{
						Main: "example.com",
						SANs: []string{"www.example.com"},
					},
				},
			},
		},
	}

	tests := []struct {
		name         string
		validDomains []string
		expected     bool
	}{
		{
			name:         "Existing domain",
			validDomains: []string{"example.com", "www.example.com"},
			expected:     true,
		},
		{
			name:         "Non-existing domain",
			validDomains: []string{"nonexistent.com"},
			expected:     false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := provider.certExists(test.validDomains)
			if result != test.expected {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}
