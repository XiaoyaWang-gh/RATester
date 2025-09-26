package page

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMergeByLanguageInterface_2(t *testing.T) {
	type testCase struct {
		name     string
		p1       Pages
		in       any
		expected any
		err      error
	}

	testCases := []testCase{
		{
			name:     "MergeByLanguageInterface_NilInput",
			p1:       Pages{},
			in:       nil,
			expected: Pages{},
			err:      nil,
		},
		{
			name:     "MergeByLanguageInterface_InvalidInput",
			p1:       Pages{},
			in:       "invalid",
			expected: nil,
			err:      fmt.Errorf("%T cannot be merged by language", "invalid"),
		},
		{
			name:     "MergeByLanguageInterface_ValidInput",
			p1:       Pages{},
			in:       Pages{},
			expected: Pages{},
			err:      nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result, err := tc.p1.MergeByLanguageInterface(tc.in)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("Expected error %v, but got %v", tc.err, err)
			}
		})
	}
}
