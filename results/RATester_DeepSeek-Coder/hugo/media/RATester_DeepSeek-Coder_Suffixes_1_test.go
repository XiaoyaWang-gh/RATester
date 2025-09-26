package media

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSuffixes_1(t *testing.T) {
	type testCase struct {
		name     string
		mt       Type
		expected []string
	}

	testCases := []testCase{
		{
			name: "no suffixes",
			mt: Type{
				SuffixesCSV: "",
			},
			expected: nil,
		},
		{
			name: "one suffix",
			mt: Type{
				SuffixesCSV: "jpg",
			},
			expected: []string{"jpg"},
		},
		{
			name: "multiple suffixes",
			mt: Type{
				SuffixesCSV: "jpg,jpeg,png",
			},
			expected: []string{"jpg", "jpeg", "png"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.mt.Suffixes()
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
