package config

import (
	"fmt"
	"testing"
)

func TestBool_3(t *testing.T) {
	type testCase struct {
		name     string
		key      string
		expected bool
		hasError bool
	}

	testCases := []testCase{
		{
			name:     "TestBool_0",
			key:      "key0",
			expected: true,
			hasError: false,
		},
		{
			name:     "TestBool_1",
			key:      "key1",
			expected: false,
			hasError: false,
		},
		{
			name:     "TestBool_2",
			key:      "key2",
			expected: true,
			hasError: true,
		},
		{
			name:     "TestBool_3",
			key:      "key3",
			expected: false,
			hasError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			fc := &fakeConfigContainer{
				data: map[string]string{
					"key0": "true",
					"key1": "false",
					"key2": "invalid",
					"key3": "",
				},
			}

			actual, err := fc.Bool(tc.key)
			if tc.hasError {
				if err == nil {
					t.Errorf("Expected error, but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if actual != tc.expected {
					t.Errorf("Expected %v, but got %v", tc.expected, actual)
				}
			}
		})
	}
}
