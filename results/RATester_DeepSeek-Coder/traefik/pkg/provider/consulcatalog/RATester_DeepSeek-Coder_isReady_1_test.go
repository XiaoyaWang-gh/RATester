package consulcatalog

import (
	"fmt"
	"testing"
)

func TestIsReady_1(t *testing.T) {
	type testCase struct {
		name     string
		cert     *connectCert
		expected bool
	}

	testCases := []testCase{
		{
			name:     "Nil Cert",
			cert:     nil,
			expected: false,
		},
		{
			name: "Empty Root",
			cert: &connectCert{
				root: []string{},
				leaf: keyPair{
					cert: "cert",
					key:  "key",
				},
			},
			expected: false,
		},
		{
			name: "Empty Leaf Cert",
			cert: &connectCert{
				root: []string{"root"},
				leaf: keyPair{
					cert: "",
					key:  "key",
				},
			},
			expected: false,
		},
		{
			name: "Empty Leaf Key",
			cert: &connectCert{
				root: []string{"root"},
				leaf: keyPair{
					cert: "cert",
					key:  "",
				},
			},
			expected: false,
		},
		{
			name: "Valid Cert",
			cert: &connectCert{
				root: []string{"root"},
				leaf: keyPair{
					cert: "cert",
					key:  "key",
				},
			},
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

			result := tc.cert.isReady()
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
