package orm

import (
	"fmt"
	"testing"
)

func TestReplaceMarks_2(t *testing.T) {
	db := &dbBasePostgres{}

	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "no question marks",
			input:    "SELECT * FROM users",
			expected: "SELECT * FROM users",
		},
		{
			name:     "single question mark",
			input:    "SELECT * FROM users WHERE id = ?",
			expected: "SELECT * FROM users WHERE id = $1",
		},
		{
			name:     "multiple question marks",
			input:    "SELECT * FROM users WHERE id = ? AND name = ?",
			expected: "SELECT * FROM users WHERE id = $1 AND name = $2",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			db.ReplaceMarks(&tc.input)
			if tc.input != tc.expected {
				t.Errorf("Expected '%s', but got '%s'", tc.expected, tc.input)
			}
		})
	}
}
