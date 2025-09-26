package tls

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/types"
)

func TestSanitizeDomains_2(t *testing.T) {
	type testCase struct {
		name     string
		domain   types.Domain
		expected []string
		err      error
	}

	testCases := []testCase{
		{
			name: "no domain",
			domain: types.Domain{
				Main: "",
				SANs: []string{},
			},
			expected: nil,
			err:      errors.New("no domain was given"),
		},
		{
			name: "one domain",
			domain: types.Domain{
				Main: "example.com",
				SANs: []string{},
			},
			expected: []string{"example.com"},
			err:      nil,
		},
		{
			name: "multiple domains",
			domain: types.Domain{
				Main: "example.com",
				SANs: []string{"www.example.com", "blog.example.com"},
			},
			expected: []string{"example.com", "www.example.com", "blog.example.com"},
			err:      nil,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			domains, err := sanitizeDomains(test.domain)
			if !reflect.DeepEqual(domains, test.expected) {
				t.Errorf("Expected %v, got %v", test.expected, domains)
			}
			if !reflect.DeepEqual(err, test.err) {
				t.Errorf("Expected error %v, got %v", test.err, err)
			}
		})
	}
}
