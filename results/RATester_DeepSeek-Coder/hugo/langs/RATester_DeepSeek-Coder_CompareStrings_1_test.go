package langs

import (
	"fmt"
	"testing"

	"golang.org/x/text/collate"
	"golang.org/x/text/language"
)

func TestCompareStrings_1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		a        string
		b        string
		expected int
	}{
		{
			name:     "strings are equal",
			a:        "abc",
			b:        "abc",
			expected: 0,
		},
		{
			name:     "strings are not equal",
			a:        "abc",
			b:        "def",
			expected: 1,
		},
		{
			name:     "strings are not equal 2",
			a:        "def",
			b:        "abc",
			expected: -1,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			t.Parallel()

			c := &Collator{
				c: collate.New(language.English, collate.Loose),
			}

			result := c.CompareStrings(tc.a, tc.b)

			if result != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, result)
			}
		})
	}
}
