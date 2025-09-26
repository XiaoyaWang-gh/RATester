package security

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"
)

func TeststringSliceToWhitelistHook_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    any
		expected Whitelist
		err      error
	}{
		{
			name:  "Test case 1",
			input: []string{"pattern1", "pattern2"},
			expected: Whitelist{
				acceptNone:      false,
				patterns:        []*regexp.Regexp{regexp.MustCompile("pattern1"), regexp.MustCompile("pattern2")},
				patternsStrings: []string{"pattern1", "pattern2"},
			},
			err: nil,
		},
		{
			name:  "Test case 2",
			input: []string{"pattern3", "pattern4"},
			expected: Whitelist{
				acceptNone:      false,
				patterns:        []*regexp.Regexp{regexp.MustCompile("pattern3"), regexp.MustCompile("pattern4")},
				patternsStrings: []string{"pattern3", "pattern4"},
			},
			err: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			actual, err := stringSliceToWhitelistHook()(reflect.TypeOf(tc.input), reflect.TypeOf(tc.expected), tc.input)
			if err != tc.err {
				t.Errorf("Expected error %v, but got %v", tc.err, err)
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, actual)
			}
		})
	}
}
