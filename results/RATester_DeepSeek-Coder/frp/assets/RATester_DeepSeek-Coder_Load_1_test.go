package assets

import (
	"fmt"
	"net/http"
	"testing"
)

func TestLoad_1(t *testing.T) {
	testCases := []struct {
		name     string
		path     string
		expected http.FileSystem
	}{
		{
			name:     "Test with non-empty path",
			path:     "/test",
			expected: http.Dir("/test"),
		},
		{
			name:     "Test with empty path",
			path:     "",
			expected: http.FS(content),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			Load(tc.path)
			if FileSystem != tc.expected {
				t.Errorf("Expected FileSystem to be %v, but got %v", tc.expected, FileSystem)
			}
		})
	}
}
