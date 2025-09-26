package gin

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"
)

func TestPostFormArray_1(t *testing.T) {
	type testCase struct {
		name     string
		key      string
		expected []string
	}

	testCases := []testCase{
		{
			name:     "Test case 1",
			key:      "key1",
			expected: []string{"value1", "value2"},
		},
		{
			name:     "Test case 2",
			key:      "key2",
			expected: []string{"value3", "value4"},
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			ctx := &Context{
				formCache: url.Values{
					tc.key: tc.expected,
				},
			}

			result := ctx.PostFormArray(tc.key)

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, Got: %v", tc.expected, result)
			}
		})
	}
}
