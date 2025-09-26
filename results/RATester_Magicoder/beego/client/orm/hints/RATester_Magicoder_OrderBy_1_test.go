package hints

import (
	"fmt"
	"testing"
)

func TestOrderBy_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected *Hint
	}{
		{
			name:  "Test case 1",
			input: "test",
			expected: &Hint{
				key:   KeyOrderBy,
				value: "test",
			},
		},
		{
			name:  "Test case 2",
			input: "another test",
			expected: &Hint{
				key:   KeyOrderBy,
				value: "another test",
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

			result := OrderBy(tc.input)
			if result.GetKey() != tc.expected.GetKey() || result.GetValue() != tc.expected.GetValue() {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
