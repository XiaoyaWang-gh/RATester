package cookie

import (
	"fmt"
	"testing"
)

func TestSanitizeName_1(t *testing.T) {
	testCases := []struct {
		name     string
		backend  string
		expected string
	}{
		{"simple", "simple", "simple"},
		{"special characters", "simple@#$%^&*()_+`~-={}[]|\\:;\"',.<>?/", "simple_______________________________"},
		{"mixed case", "simpleMiXeD", "simpleMiXeD"},
		{"numbers", "simple123", "simple123"},
		{"mixed", "simple123@#$%^&*()_+`~-={}[]|\\:;\"',.<>?/MiXeD", "simple123_____________________________MiXeD"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := sanitizeName(tc.backend)
			if result != tc.expected {
				t.Errorf("Expected '%s', got '%s'", tc.expected, result)
			}
		})
	}
}
