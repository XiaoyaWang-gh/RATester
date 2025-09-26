package maps

import (
	"fmt"
	"testing"
)

func TestkeyPath_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	renamer := KeyRenamer{}

	testCases := []struct {
		k1, k2, expected string
	}{
		{"", "", ""},
		{"", "k2", "k2"},
		{"k1", "", "k1"},
		{"k1", "k2", "k1/k2"},
		{"K1", "K2", "k1/k2"},
	}

	for _, tc := range testCases {
		result := renamer.keyPath(tc.k1, tc.k2)
		if result != tc.expected {
			t.Errorf("Expected %s, got %s", tc.expected, result)
		}
	}
}
