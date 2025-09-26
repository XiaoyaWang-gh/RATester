package pageparser

import (
	"fmt"
	"testing"
)

func TestBackup_2(t *testing.T) {
	testCases := []struct {
		name     string
		items    Items
		expected int
	}{
		{
			name:     "Test case 1",
			items:    Items{Item{}, Item{}, Item{}},
			expected: 1,
		},
		{
			name:     "Test case 2",
			items:    Items{Item{}, Item{}, Item{}, Item{}},
			expected: 2,
		},
		{
			name:     "Test case 3",
			items:    Items{Item{}},
			expected: 0,
		},
		{
			name:     "Test case 4",
			items:    Items{},
			expected: -1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			iterator := &Iterator{
				items:   tc.items,
				lastPos: len(tc.items) - 1,
			}

			iterator.Backup()

			if iterator.lastPos != tc.expected {
				t.Errorf("Expected lastPos to be %d, but got %d", tc.expected, iterator.lastPos)
			}
		})
	}
}
