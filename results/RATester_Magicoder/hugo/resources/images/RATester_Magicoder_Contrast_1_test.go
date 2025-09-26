package images

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/disintegration/gift"
)

func TestContrast_1(t *testing.T) {
	filters := &Filters{}

	testCases := []struct {
		name       string
		percentage any
		expected   gift.Filter
	}{
		{
			name:       "Test case 1",
			percentage: 0.5,
			expected:   gift.Contrast(0.5),
		},
		{
			name:       "Test case 2",
			percentage: 1.0,
			expected:   gift.Contrast(1.0),
		},
		{
			name:       "Test case 3",
			percentage: 2.0,
			expected:   gift.Contrast(2.0),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := filters.Contrast(tc.percentage)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, but got: %v", tc.expected, result)
			}
		})
	}
}
