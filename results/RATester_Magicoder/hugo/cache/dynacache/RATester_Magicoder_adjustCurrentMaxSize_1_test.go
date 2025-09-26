package dynacache

import (
	"fmt"
	"testing"
)

func TestadjustCurrentMaxSize_1(t *testing.T) {
	testCases := []struct {
		name             string
		opts             Options
		currentMaxSize   int
		adjustmentFactor float64
		expected         bool
	}{
		{
			name: "Test case 1: MaxSize is less than MinMaxSize",
			opts: Options{
				MaxSize:    10,
				MinMaxSize: 20,
			},
			currentMaxSize:   10,
			adjustmentFactor: 0.5,
			expected:         true,
		},
		{
			name: "Test case 2: MaxSize is greater than MinMaxSize",
			opts: Options{
				MaxSize:    20,
				MinMaxSize: 10,
			},
			currentMaxSize:   20,
			adjustmentFactor: 0.5,
			expected:         false,
		},
		{
			name: "Test case 3: MaxSize is equal to MinMaxSize",
			opts: Options{
				MaxSize:    20,
				MinMaxSize: 20,
			},
			currentMaxSize:   20,
			adjustmentFactor: 0.5,
			expected:         false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			s := &stats{
				opts:             tc.opts,
				currentMaxSize:   tc.currentMaxSize,
				adjustmentFactor: tc.adjustmentFactor,
			}

			result := s.adjustCurrentMaxSize()
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
