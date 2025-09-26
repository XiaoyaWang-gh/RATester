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
			name:    "case 1",
			domains: []string{"example.com", "www.example.com", "blog.example.com"},
			expected: Domain{
				Main: "example.com",
				SANs: []string{"www.example.com", "blog.example.com"},
			},
		},
		{
			name:    "case 2",
			domains: []string{"example.net", "www.example.net"},
			expected: Domain{
				Main: "example.net",
				SANs: []string{"www.example.net"},
			},
		},
		{
			name:    "case 3",
			domains: []string{"example.org"},
			expected: Domain{
				Main: "example.org",
				SANs: []string{},
			},
		},
		{
			name:    "case 4",
			domains: []string{},
			expected: Domain{
				Main: "",
				SANs: []string{},
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
				t.Errorf("Expected Main to be %q, got %q", test.expected.Main, d.Main)
			}

			if !reflect.DeepEqual(d.SANs, test.expected.SANs) {
				t.Errorf("Expected SANs to be %q, got %q", test.expected.SANs, d.SANs)
			}
		})
	}
}
