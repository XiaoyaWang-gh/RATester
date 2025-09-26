package doctree

import (
	"fmt"
	"testing"
)

func TestIndex_1(t *testing.T) {
	tests := []struct {
		name     string
		d        DimensionFlag
		expected int
	}{
		{
			name:     "Test DimensionFlag 1",
			d:        1,
			expected: 0,
		},
		{
			name:     "Test DimensionFlag 2",
			d:        2,
			expected: 1,
		},
		{
			name:     "Test DimensionFlag 3",
			d:        3,
			expected: 2,
		},
		{
			name:     "Test DimensionFlag 4",
			d:        4,
			expected: 3,
		},
		{
			name:     "Test DimensionFlag 5",
			d:        5,
			expected: 4,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := test.d.Index()
			if result != test.expected {
				t.Errorf("Expected %d, but got %d", test.expected, result)
			}
		})
	}
}
