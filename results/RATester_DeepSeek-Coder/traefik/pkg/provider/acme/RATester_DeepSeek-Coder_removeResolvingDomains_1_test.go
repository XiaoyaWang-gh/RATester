package acme

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRemoveResolvingDomains_1(t *testing.T) {
	type testCase struct {
		name             string
		resolvingDomains []string
		expected         map[string]struct{}
	}

	testCases := []testCase{
		{
			name:             "remove existing domain",
			resolvingDomains: []string{"example.com"},
			expected: map[string]struct{}{
				"example2.com": {},
			},
		},
		{
			name:             "remove non-existing domain",
			resolvingDomains: []string{"nonexistent.com"},
			expected: map[string]struct{}{
				"example.com":  {},
				"example2.com": {},
			},
		},
		{
			name:             "remove multiple existing domains",
			resolvingDomains: []string{"example.com", "example2.com"},
			expected:         map[string]struct{}{},
		},
		{
			name:             "remove multiple non-existing domains",
			resolvingDomains: []string{"nonexistent1.com", "nonexistent2.com"},
			expected: map[string]struct{}{
				"example.com":  {},
				"example2.com": {},
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			p := &Provider{
				resolvingDomains: map[string]struct{}{
					"example.com":  {},
					"example2.com": {},
				},
			}

			p.removeResolvingDomains(test.resolvingDomains)

			if !reflect.DeepEqual(p.resolvingDomains, test.expected) {
				t.Errorf("Expected %v, got %v", test.expected, p.resolvingDomains)
			}
		})
	}
}
