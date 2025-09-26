package capture

import (
	"fmt"
	"testing"
)

func TestSize_1(t *testing.T) {
	testCases := []struct {
		name     string
		size     int64
		expected int64
	}{
		{"Test case 1", 1024, 1024},
		{"Test case 2", 2048, 2048},
		{"Test case 3", 4096, 4096},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			crw := &captureResponseWriter{
				size: tc.size,
			}

			result := crw.Size()
			if result != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, result)
			}
		})
	}
}
