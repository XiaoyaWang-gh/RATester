package security

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStringSliceToWhitelistHook_1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    any
		expected any
		err      error
	}{
		{
			name:  "Test case 1",
			input: []string{"test1", "test2", "test3"},
			expected: func() (any, error) {
				return NewWhitelist("test1", "test2", "test3")
			},
			err: nil,
		},
		{
			name:  "Test case 2",
			input: []string{},
			expected: func() (any, error) {
				return NewWhitelist()
			},
			err: nil,
		},
		{
			name:  "Test case 3",
			input: "not a slice",
			expected: func() (any, error) {
				return nil, nil
			},
			err: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			actual, err := stringSliceToWhitelistHook()(reflect.TypeOf(tc.input), reflect.TypeOf(tc.expected), tc.input)
			if err != tc.err {
				t.Errorf("Expected error %v, got %v", tc.err, err)
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
