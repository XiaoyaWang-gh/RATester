package images

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/disintegration/gift"
)

func TestSaturation_1(t *testing.T) {
	filters := &Filters{}

	testCases := []struct {
		name       string
		percentage any
		expected   gift.Filter
	}{
		{
			name:       "Test case 1",
			percentage: 0.5,
			expected:   filter{Options: newFilterOpts(0.5), Filter: gift.Saturation(0.5)},
		},
		{
			name:       "Test case 2",
			percentage: 1.0,
			expected:   filter{Options: newFilterOpts(1.0), Filter: gift.Saturation(1.0)},
		},
		{
			name:       "Test case 3",
			percentage: 0.0,
			expected:   filter{Options: newFilterOpts(0.0), Filter: gift.Saturation(0.0)},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := filters.Saturation(tc.percentage)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, but got: %v", tc.expected, result)
			}
		})
	}
}
