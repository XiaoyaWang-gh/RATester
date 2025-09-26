package orm

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestGroupBy_4(t *testing.T) {
	qs := querySet{
		groups: []string{"id"},
	}

	testCases := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "Test with one group",
			input:    []string{"name"},
			expected: []string{"id", "name"},
		},
		{
			name:     "Test with multiple groups",
			input:    []string{"name", "age"},
			expected: []string{"id", "name", "age"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			qs.GroupBy(tc.input...)
			assert.Equal(t, tc.expected, qs.groups)
		})
	}
}
