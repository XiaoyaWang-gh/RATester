package template

import (
	"fmt"
	"testing"
)

func TestAsciiAlphaNum_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    byte
		expected bool
	}{
		{"lowercase letter", 'a', true},
		{"uppercase letter", 'A', true},
		{"number", '0', true},
		{"special character", '@', false},
		{"space", ' ', false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := asciiAlphaNum(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
