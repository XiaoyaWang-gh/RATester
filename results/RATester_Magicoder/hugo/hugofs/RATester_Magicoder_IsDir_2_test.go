package hugofs

import (
	"fmt"
	"testing"
	"time"
)

func TestIsDir_2(t *testing.T) {
	testCases := []struct {
		name     string
		fileInfo *dirNameOnlyFileInfo
		expected bool
	}{
		{
			name: "Test case 1: IsDir should return true",
			fileInfo: &dirNameOnlyFileInfo{
				name:    "test",
				modTime: time.Now(),
			},
			expected: true,
		},
		{
			name: "Test case 2: IsDir should return false",
			fileInfo: &dirNameOnlyFileInfo{
				name:    "test",
				modTime: time.Now(),
			},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.fileInfo.IsDir()
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
