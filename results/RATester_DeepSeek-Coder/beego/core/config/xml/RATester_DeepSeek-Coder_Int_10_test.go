package xml

import (
	"fmt"
	"strconv"
	"testing"
)

func TestInt_10(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name     string
		data     map[string]interface{}
		key      string
		expected int
		err      error
	}

	testCases := []testCase{
		{
			name:     "valid int",
			data:     map[string]interface{}{"key": "123"},
			key:      "key",
			expected: 123,
			err:      nil,
		},
		{
			name:     "invalid int",
			data:     map[string]interface{}{"key": "abc"},
			key:      "key",
			expected: 0,
			err:      strconv.ErrSyntax,
		},
		{
			name:     "key not exist",
			data:     map[string]interface{}{},
			key:      "key",
			expected: 0,
			err:      fmt.Errorf("key not exist"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &ConfigContainer{
				data: tc.data,
			}

			actual, err := c.Int(tc.key)
			if actual != tc.expected || err != tc.err {
				t.Errorf("expected %v, %v, got %v, %v", tc.expected, tc.err, actual, err)
			}
		})
	}
}
