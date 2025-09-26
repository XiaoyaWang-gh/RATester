package hstrings

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"
)

func TestGet_13(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name     string
		key      string
		expected *regexp.Regexp
		wantOk   bool
	}

	testCases := []testCase{
		{
			name:     "Existing key",
			key:      "existingKey",
			expected: regexp.MustCompile("existingRegexp"),
			wantOk:   true,
		},
		{
			name:     "Non-existing key",
			key:      "nonExistingKey",
			expected: nil,
			wantOk:   false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			rc := &regexpCache{
				re: map[string]*regexp.Regexp{
					"existingKey": regexp.MustCompile("existingRegexp"),
				},
			}

			got, gotOk := rc.get(tc.key)
			if !reflect.DeepEqual(got, tc.expected) || gotOk != tc.wantOk {
				t.Errorf("Expected (%v, %v), got (%v, %v)", tc.expected, tc.wantOk, got, gotOk)
			}
		})
	}
}
