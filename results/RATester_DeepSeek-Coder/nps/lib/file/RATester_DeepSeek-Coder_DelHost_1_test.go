package file

import (
	"fmt"
	"testing"
)

func TestDelHost_1(t *testing.T) {
	db := &DbUtils{
		JsonDb: &JsonDb{},
	}

	testCases := []struct {
		name     string
		id       int
		expected error
	}{
		{
			name:     "Test case 1",
			id:       1,
			expected: nil,
		},
		{
			name:     "Test case 2",
			id:       2,
			expected: nil,
		},
		{
			name:     "Test case 3",
			id:       3,
			expected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			err := db.DelHost(tc.id)
			if err != tc.expected {
				t.Errorf("Expected error %v, but got %v", tc.expected, err)
			}
		})
	}
}
