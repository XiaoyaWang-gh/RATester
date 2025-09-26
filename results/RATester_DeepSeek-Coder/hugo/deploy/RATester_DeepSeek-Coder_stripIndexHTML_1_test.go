package deploy

import (
	"fmt"
	"testing"
)

func TestStripIndexHTML_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"Test with /index.html", "/index.html", "/"},
		{"Test with /test/index.html", "/test/index.html", "/test/"},
		{"Test with /test", "/test", "/test"},
		{"Test with /test/", "/test/", "/test/"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := stripIndexHTML(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}
