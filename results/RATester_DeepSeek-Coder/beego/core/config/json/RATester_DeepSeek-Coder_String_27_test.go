package json

import (
	"fmt"
	"testing"
)

func TestString_27(t *testing.T) {
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

	container := &JSONConfigContainer{
		data: map[string]interface{}{
			"key0": "value0",
			"key1": "value1",
			"key2": 123,
			"key3": []string{"val1", "val2"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			val, err := container.String(tc.key)
			if val != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, val)
			}
			if err != tc.err {
				t.Errorf("Expected error %v, got %v", tc.err, err)
			}
		})
	}
}
