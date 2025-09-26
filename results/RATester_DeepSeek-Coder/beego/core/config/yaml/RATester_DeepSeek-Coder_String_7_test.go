package yaml

import (
	"fmt"
	"testing"
)

func TestString_7(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name     string
		key      string
		expected string
		err      error
	}

	testCases := []testCase{
		{
			name:     "TestString_0",
			key:      "key0",
			expected: "value0",
			err:      nil,
		},
		{
			name:     "TestString_1",
			key:      "key1",
			expected: "value1",
			err:      nil,
		},
		{
			name:     "TestString_2",
			key:      "key2",
			expected: "",
			err:      nil,
		},
		{
			name:     "TestString_3",
			key:      "key3",
			expected: "",
			err:      nil,
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
					tc.key: tc.expected,
				},
			}

			actual, err := c.String(tc.key)
			if actual != tc.expected || err != tc.err {
				t.Errorf("Expected (%v, %v), got (%v, %v)", tc.expected, tc.err, actual, err)
			}
		})
	}
}
