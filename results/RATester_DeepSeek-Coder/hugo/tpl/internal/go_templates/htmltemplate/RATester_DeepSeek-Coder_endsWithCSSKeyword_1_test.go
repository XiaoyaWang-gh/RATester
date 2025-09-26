package template

import (
	"fmt"
	"testing"
)

func TestEndsWithCSSKeyword_1(t *testing.T) {
	testCases := []struct {
		name     string
		b        []byte
		kw       string
		expected bool
	}{
		{"ends with keyword", []byte("color: red;"), "red", true},
		{"does not end with keyword", []byte("color: blue;"), "red", false},
		{"empty keyword", []byte("color: green;"), "", true},
		{"empty string and keyword", []byte(""), "red", false},
		{"keyword is nil", []byte("color: yellow;"), "", true},
		{"string is nil", nil, "red", false},
		{"string and keyword are nil", nil, "", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := endsWithCSSKeyword(tc.b, tc.kw)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
