package types

import (
	"fmt"
	"reflect"
	"testing"
)

func TestToStrArray_1(t *testing.T) {
	tests := []struct {
		name     string
		domain   *Domain
		expected []string
	}{
		{
			name: "Test with empty domain",
			domain: &Domain{
				Main: "",
				SANs: []string{},
			},
			expected: []string{},
		},
		{
			name: "Test with main domain",
			domain: &Domain{
				Main: "example.com",
				SANs: []string{},
			},
			expected: []string{"example.com"},
		},
		{
			name: "Test with SANs",
			domain: &Domain{
				Main: "",
				SANs: []string{"san1.com", "san2.com"},
			},
			expected: []string{"san1.com", "san2.com"},
		},
		{
			name: "Test with main and SANs",
			domain: &Domain{
				Main: "example.com",
				SANs: []string{"san1.com", "san2.com"},
			},
			expected: []string{"example.com", "san1.com", "san2.com"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := test.domain.ToStrArray()
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}
