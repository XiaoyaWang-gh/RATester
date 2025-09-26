package glob

import (
	"fmt"
	"testing"
)

func TestNormalizePath_1(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"/", "/"},
		{"//", "/"},
		{"///", "/"},
		{"/a", "/a"},
		{"/a/", "/a"},
		{"/a//", "/a"},
		{"/a/b", "/a/b"},
		{"/a/b/", "/a/b"},
		{"/a/b//", "/a/b"},
		{"/a/b/c", "/a/b/c"},
		{"/a/b/c/", "/a/b/c"},
		{"/a/b/c//", "/a/b/c"},
		{"/a/b/c/d", "/a/b/c/d"},
		{"/a/b/c/d/", "/a/b/c/d"},
		{"/a/b/c/d//", "/a/b/c/d"},
		{"/a/b/c/d/e", "/a/b/c/d/e"},
		{"/a/b/c/d/e/", "/a/b/c/d/e"},
		{"/a/b/c/d/e//", "/a/b/c/d/e"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := NormalizePath(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}
