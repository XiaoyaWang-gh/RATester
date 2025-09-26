package hstrings

import (
	"fmt"
	"regexp"
	"testing"
)

func TestSet_5(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name string
		key  string
		re   *regexp.Regexp
	}

	testCases := []testCase{
		{
			name: "Test case 1",
			key:  "key1",
			re:   regexp.MustCompile("regex1"),
		},
		{
			name: "Test case 2",
			key:  "key2",
			re:   regexp.MustCompile("regex2"),
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

			rc := &regexpCache{
				re: make(map[string]*regexp.Regexp),
			}

			rc.set(tc.key, tc.re)

			if r, ok := rc.re[tc.key]; !ok || r != tc.re {
				t.Errorf("Expected regexpCache to contain %v, but it did not", tc.re)
			}
		})
	}
}
