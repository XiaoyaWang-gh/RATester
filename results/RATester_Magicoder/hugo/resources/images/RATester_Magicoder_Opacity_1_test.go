package images

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/disintegration/gift"
)

func TestOpacity_1(t *testing.T) {
	filters := &Filters{}

	testCases := []struct {
		name     string
		opacity  any
		expected gift.Filter
	}{
		{
			name:    "Test case 1",
			opacity: 0.5,
			expected: filter{
				Options: newFilterOpts(0.5),
				Filter:  opacityFilter{opacity: 0.5},
			},
		},
		{
			name:    "Test case 2",
			opacity: 1.0,
			expected: filter{
				Options: newFilterOpts(1.0),
				Filter:  opacityFilter{opacity: 1.0},
			},
		},
		{
			name:    "Test case 3",
			opacity: 0.0,
			expected: filter{
				Options: newFilterOpts(0.0),
				Filter:  opacityFilter{opacity: 0.0},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := filters.Opacity(tc.opacity)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, but got: %v", tc.expected, result)
			}
		})
	}
}
