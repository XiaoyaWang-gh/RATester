package echo

import (
	"fmt"
	"testing"
)

func TestNormalizePathSlash_1(t *testing.T) {
	type test struct {
		name     string
		path     string
		expected string
	}

	tests := []test{
		{"empty path", "", "/"},
		{"path with leading slash", "/test", "/test"},
		{"path without leading slash", "test", "/test"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := normalizePathSlash(tt.path)
			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}
