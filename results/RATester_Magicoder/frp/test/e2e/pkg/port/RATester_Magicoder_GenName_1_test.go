package port

import (
	"fmt"
	"testing"
)

func TestGenName_1(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		options  []NameOption
		expected string
	}{
		{
			name:     "no options",
			input:    "test-name",
			options:  []NameOption{},
			expected: "testname",
		},
		{
			name:     "with option",
			input:    "test-name",
			options:  []NameOption{func(b *nameBuilder) *nameBuilder { b.name = "new-name"; return b }},
			expected: "newname",
		},
		// Add more test cases as needed
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := GenName(test.input, test.options...)
			if result != test.expected {
				t.Errorf("Expected %s, but got %s", test.expected, result)
			}
		})
	}
}
