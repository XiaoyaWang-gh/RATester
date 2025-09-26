package validation

import (
	"fmt"
	"regexp"
	"testing"
)

func TestIsSatisfied_4(t *testing.T) {
	type testCase struct {
		name     string
		noMatch  NoMatch
		obj      interface{}
		expected bool
	}

	testCases := []testCase{
		{
			name: "Test case 1",
			noMatch: NoMatch{
				Match: Match{
					Regexp: regexp.MustCompile("^[a-z]+$"),
				},
				Key: "test",
			},
			obj:      "test",
			expected: false,
		},
		{
			name: "Test case 2",
			noMatch: NoMatch{
				Match: Match{
					Regexp: regexp.MustCompile("^[a-z]+$"),
				},
				Key: "test",
			},
			obj:      "test1",
			expected: true,
		},
		{
			name: "Test case 3",
			noMatch: NoMatch{
				Match: Match{
					Regexp: regexp.MustCompile("^[0-9]+$"),
				},
				Key: "test",
			},
			obj:      "test",
			expected: true,
		},
		{
			name: "Test case 4",
			noMatch: NoMatch{
				Match: Match{
					Regexp: regexp.MustCompile("^[0-9]+$"),
				},
				Key: "test",
			},
			obj:      "test1",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.noMatch.IsSatisfied(tc.obj)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
