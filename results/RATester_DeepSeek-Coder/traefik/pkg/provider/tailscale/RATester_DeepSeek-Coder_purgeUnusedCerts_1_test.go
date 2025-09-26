package tailscale

import (
	"fmt"
	"testing"

	traefiktls "github.com/traefik/traefik/v3/pkg/tls"
)

func TestPurgeUnusedCerts_1(t *testing.T) {
	type testCase struct {
		name     string
		domains  []string
		expected bool
	}

	testCases := []testCase{
		{
			name:     "Purge unused certs",
			domains:  []string{"example.com", "foo.com", "bar.com"},
			expected: true,
		},
		{
			name:     "No purge",
			domains:  []string{"example.com", "foo.com"},
			expected: false,
		},
		{
			name:     "Purge all",
			domains:  []string{},
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			p := &Provider{
				certByDomain: map[string]traefiktls.Certificate{
					"example.com": {},
					"foo.com":     {},
					"bar.com":     {},
				},
			}

			result := p.purgeUnusedCerts(tc.domains)

			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
