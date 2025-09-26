package fiber

import (
	"fmt"
	"testing"
)

func TestIsEtagStale_1(t *testing.T) {
	app := &App{
		getString: func(b []byte) string {
			return string(b)
		},
	}

	testCases := []struct {
		name           string
		etag           string
		noneMatchBytes []byte
		expected       bool
	}{
		{
			name:           "Test case 1",
			etag:           "etag1",
			noneMatchBytes: []byte("etag1"),
			expected:       false,
		},
		{
			name:           "Test case 2",
			etag:           "etag2",
			noneMatchBytes: []byte("etag2"),
			expected:       false,
		},
		{
			name:           "Test case 3",
			etag:           "etag3",
			noneMatchBytes: []byte("etag3"),
			expected:       false,
		},
		{
			name:           "Test case 4",
			etag:           "etag4",
			noneMatchBytes: []byte("etag4"),
			expected:       false,
		},
		{
			name:           "Test case 5",
			etag:           "etag5",
			noneMatchBytes: []byte("etag5"),
			expected:       false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := app.isEtagStale(tc.etag, tc.noneMatchBytes)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
