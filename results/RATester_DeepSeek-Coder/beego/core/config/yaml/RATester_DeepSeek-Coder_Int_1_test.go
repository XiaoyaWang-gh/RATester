package yaml

import (
	"errors"
	"fmt"
	"testing"
)

func TestInt_1(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name     string
		key      string
		expected int
		err      error
	}

	testCases := []testCase{
		{
			name:     "TestInt_ExistingKey",
			key:      "existingKey",
			expected: 1,
			err:      nil,
		},
		{
			name:     "TestInt_NonExistingKey",
			key:      "nonExistingKey",
			expected: 0,
			err:      errors.New("not int value"),
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
				data: map[string]interface{}{
					"existingKey": 1,
				},
			}

			actual, err := c.Int(tc.key)
			if actual != tc.expected || err != tc.err {
				t.Errorf("Expected (%v, %v), got (%v, %v)", tc.expected, tc.err, actual, err)
			}
		})
	}
}
