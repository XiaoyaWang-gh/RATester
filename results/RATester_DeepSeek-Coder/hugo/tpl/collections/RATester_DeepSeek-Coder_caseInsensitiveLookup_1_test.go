package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCaseInsensitiveLookup_1(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name     string
		m        map[string]int
		k        string
		expected bool
	}

	testCases := []testCase{
		{
			name: "Test case-insensitive lookup with exact match",
			m: map[string]int{
				"Hello": 1,
				"WORLD": 2,
			},
			k:        "hello",
			expected: true,
		},
		{
			name: "Test case-insensitive lookup with partial match",
			m: map[string]int{
				"Hello": 1,
				"WORLD": 2,
			},
			k:        "WORLd",
			expected: true,
		},
		{
			name: "Test case-insensitive lookup with no match",
			m: map[string]int{
				"Hello": 1,
				"WORLD": 2,
			},
			k:        "goodbye",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			m := reflect.ValueOf(tc.m)
			k := reflect.ValueOf(tc.k)

			_, actual := caseInsensitiveLookup(m, k)

			if actual != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, actual)
			}
		})
	}
}
