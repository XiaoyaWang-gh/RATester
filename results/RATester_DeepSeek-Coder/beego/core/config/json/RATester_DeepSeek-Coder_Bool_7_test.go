package json

import (
	"fmt"
	"testing"
)

func TestBool_7(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name     string
		key      string
		expected bool
		err      error
	}

	testCases := []testCase{
		{
			name:     "exist key",
			key:      "existKey",
			expected: true,
			err:      nil,
		},
		{
			name:     "not exist key",
			key:      "notExistKey",
			expected: false,
			err:      fmt.Errorf("not exist key: %q", "notExistKey"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			container := &JSONConfigContainer{
				data: map[string]interface{}{
					"existKey": true,
				},
			}

			actual, err := container.Bool(tc.key)
			if actual != tc.expected || err != tc.err {
				t.Errorf("expected %v, actual %v, err %v", tc.expected, actual, err)
			}
		})
	}
}
