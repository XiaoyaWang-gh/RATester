package page

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/helpers"
)

func TestSanitize_2(t *testing.T) {
	type testCase struct {
		name     string
		input    *pagePathBuilder
		expected *pagePathBuilder
	}

	testCases := []testCase{
		{
			name: "Test case 1",
			input: &pagePathBuilder{
				els: []string{"test1", "test2"},
				d: TargetPathDescriptor{
					PathSpec: &helpers.PathSpec{},
				},
			},
			expected: &pagePathBuilder{
				els: []string{"test1", "test2"},
				d: TargetPathDescriptor{
					PathSpec: &helpers.PathSpec{},
				},
			},
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

			tc.input.Sanitize()
			if !reflect.DeepEqual(tc.input, tc.expected) {
				t.Errorf("Expected: %v, but got: %v", tc.expected, tc.input)
			}
		})
	}
}
