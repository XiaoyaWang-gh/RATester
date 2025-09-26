package template

import (
	"fmt"
	"testing"
)

func TestIsCSSSpace_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    byte
		expected bool
	}{
		{"Tab", '\t', true},
		{"Newline", '\n', true},
		{"Formfeed", '\f', true},
		{"Return", '\r', true},
		{"Space", ' ', true},
		{"Other", 'a', false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := isCSSSpace(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
