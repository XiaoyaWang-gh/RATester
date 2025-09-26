package gin

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"
)

func TestGetQueryMap_1(t *testing.T) {
	type testCase struct {
		name     string
		key      string
		expected map[string]string
		has      bool
	}

	testCases := []testCase{
		{
			name:     "Test case 1",
			key:      "test_key",
			expected: map[string]string{"test_key": "test_value"},
			has:      true,
		},
		{
			name:     "Test case 2",
			key:      "non_existing_key",
			expected: map[string]string{},
			has:      false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			ctx := &Context{
				queryCache: url.Values{
					tc.key: []string{"test_value"},
				},
			}

			result, has := ctx.GetQueryMap(tc.key)

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}

			if has != tc.has {
				t.Errorf("Expected has to be %v, but got %v", tc.has, has)
			}
		})
	}
}
