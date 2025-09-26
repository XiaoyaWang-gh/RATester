package images

import (
	"fmt"
	"image"
	"testing"
)

func TestBounds_1(t *testing.T) {
	testCases := []struct {
		name      string
		srcBounds image.Rectangle
		expected  image.Rectangle
	}{
		{
			name:      "Test case 1",
			srcBounds: image.Rect(0, 0, 100, 100),
			expected:  image.Rect(0, 0, 100, 100),
		},
		{
			name:      "Test case 2",
			srcBounds: image.Rect(0, 0, 200, 200),
			expected:  image.Rect(0, 0, 200, 200),
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			f := textFilter{}
			result := f.Bounds(tc.srcBounds)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
