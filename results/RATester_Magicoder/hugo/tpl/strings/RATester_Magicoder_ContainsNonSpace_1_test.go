package strings

import (
	"fmt"
	"testing"
)

func TestContainsNonSpace_1(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name     string
		input    any
		expected bool
	}{
		{"Empty string", "", false},
		{"String with spaces", " ", false},
		{"String with non-space characters", "Hello", true},
		{"String with non-space and space characters", " Hello ", true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result, err := ns.ContainsNonSpace(test.input)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if result != test.expected {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}
