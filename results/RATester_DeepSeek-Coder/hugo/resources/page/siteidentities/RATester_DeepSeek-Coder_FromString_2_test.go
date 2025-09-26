package siteidentities

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/identity"
)

func TestFromString_2(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected identity.Identity
		ok       bool
	}{
		{
			name:     "Data",
			input:    "Data",
			expected: Data,
			ok:       true,
		},
		{
			name:     "Unknown",
			input:    "Unknown",
			expected: identity.Anonymous,
			ok:       false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, ok := FromString(test.input)
			if ok != test.ok {
				t.Errorf("Expected ok to be %v, got %v", test.ok, ok)
			}
			if !reflect.DeepEqual(got, test.expected) {
				t.Errorf("Expected identity to be %v, got %v", test.expected, got)
			}
		})
	}
}
