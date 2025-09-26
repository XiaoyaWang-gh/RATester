package gin

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"
)

func TestGetPostFormArray_1(t *testing.T) {
	type testCase struct {
		name     string
		key      string
		expected []string
		setup    func(c *Context)
	}

	testCases := []testCase{
		{
			name:     "Test case 1",
			key:      "key1",
			expected: []string{"value1", "value2"},
			setup: func(c *Context) {
				c.formCache = url.Values{"key1": {"value1", "value2"}}
			},
		},
		{
			name:     "Test case 2",
			key:      "key2",
			expected: []string{"value3", "value4"},
			setup: func(c *Context) {
				c.formCache = url.Values{"key2": {"value3", "value4"}}
			},
		},
		{
			name:     "Test case 3",
			key:      "key3",
			expected: []string{},
			setup: func(c *Context) {
				c.formCache = url.Values{}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			ctx := &Context{}
			tc.setup(ctx)
			values, ok := ctx.GetPostFormArray(tc.key)
			if !ok {
				t.Errorf("Expected ok to be true, got false")
			}
			if !reflect.DeepEqual(values, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, values)
			}
		})
	}
}
