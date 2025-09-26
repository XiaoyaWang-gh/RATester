package maps

import (
	"fmt"
	"reflect"
	"testing"
)

func TestForEeach_1(t *testing.T) {
	t.Parallel()

	type K string
	type T int

	testCases := []struct {
		name     string
		cache    *Cache[K, T]
		expected map[K]T
	}{
		{
			name: "empty cache",
			cache: &Cache[K, T]{
				m: make(map[K]T),
			},
			expected: make(map[K]T),
		},
		{
			name: "non-empty cache",
			cache: &Cache[K, T]{
				m: map[K]T{
					"key1": 1,
					"key2": 2,
					"key3": 3,
				},
			},
			expected: map[K]T{
				"key1": 1,
				"key2": 2,
				"key3": 3,
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

			actual := make(map[K]T)
			tc.cache.ForEeach(func(k K, v T) {
				actual[k] = v
			})

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
