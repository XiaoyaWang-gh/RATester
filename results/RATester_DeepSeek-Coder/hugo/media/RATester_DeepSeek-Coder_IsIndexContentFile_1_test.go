package media

import (
	"fmt"
	"testing"
)

func TestIsIndexContentFile_1(t *testing.T) {
	contentTypes := ContentTypes{
		HTML: Type{
			MainType: "text",
		},
	}

	testCases := []struct {
		name     string
		filename string
		expected bool
	}{
		{
			name:     "index file",
			filename: "index.html",
			expected: true,
		},
		{
			name:     "underscore index file",
			filename: "_index.html",
			expected: true,
		},
		{
			name:     "non-index file",
			filename: "content.html",
			expected: false,
		},
		{
			name:     "non-content file",
			filename: "non-content.txt",
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

			result := contentTypes.IsIndexContentFile(tc.filename)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
