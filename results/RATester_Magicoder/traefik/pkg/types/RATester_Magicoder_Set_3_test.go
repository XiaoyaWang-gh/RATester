package types

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSet_3(t *testing.T) {
	tests := []struct {
		name     string
		domains  []string
		expected Domain
	}{
		{
			name:    "empty",
			domains: []string{},
			expected: Domain{
				Main: "",
				SANs: []string{},
			},
		},
		{
			name:    "one domain",
			domains: []string{"example.com"},
			expected: Domain{
				Main: "example.com",
				SANs: []string{},
			},
		},
		{
			name:    "multiple domains",
			domains: []string{"example.com", "www.example.com", "blog.example.com"},
			expected: Domain{
				Main: "example.com",
				SANs: []string{"www.example.com", "blog.example.com"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			d := &Domain{}
			d.Set(test.domains)

			if d.Main != test.expected.Main {
				t.Errorf("Expected Main to be %s, but got %s", test.expected.Main, d.Main)
			}

			if !reflect.DeepEqual(d.SANs, test.expected.SANs) {
				t.Errorf("Expected SANs to be %v, but got %v", test.expected.SANs, d.SANs)
			}
		})
	}
}
