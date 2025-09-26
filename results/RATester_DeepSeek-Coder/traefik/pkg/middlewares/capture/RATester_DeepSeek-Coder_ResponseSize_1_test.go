package capture

import (
	"fmt"
	"testing"
)

func TestResponseSize_1(t *testing.T) {
	testCases := []struct {
		name string
		size int64
		want int64
	}{
		{
			name: "Test case 1",
			size: 1024,
			want: 1024,
		},
		{
			name: "Test case 2",
			size: 2048,
			want: 2048,
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

			c := &Capture{
				crw: &captureResponseWriter{
					size: tc.size,
				},
			}

			got := c.ResponseSize()

			if got != tc.want {
				t.Errorf("Expected ResponseSize() to return %d, but got %d", tc.want, got)
			}
		})
	}
}
