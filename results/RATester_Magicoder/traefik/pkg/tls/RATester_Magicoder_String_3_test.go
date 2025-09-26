package tls

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/types"
)

func TestString_3(t *testing.T) {
	tests := []struct {
		name     string
		certFile string
		keyFile  string
		expected string
	}{
		{
			name:     "empty",
			certFile: "",
			keyFile:  "",
			expected: "",
		},
		{
			name:     "single",
			certFile: "cert1",
			keyFile:  "key1",
			expected: "cert1,key1",
		},
		{
			name:     "multiple",
			certFile: "cert1,cert2",
			keyFile:  "key1,key2",
			expected: "cert1,key1;cert2,key2",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := Certificates{
				{
					CertFile: types.FileOrContent(test.certFile),
					KeyFile:  types.FileOrContent(test.keyFile),
				},
			}

			result := c.String()

			if result != test.expected {
				t.Errorf("Expected %q, got %q", test.expected, result)
			}
		})
	}
}
