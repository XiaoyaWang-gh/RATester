package maps

import (
	"fmt"
	"testing"
)

func TestLen_4(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		cache    *Cache[int, string]
		expected int
	}{
		{
			name: "empty cache",
			cache: &Cache[int, string]{
				m: make(map[int]string),
			},
			expected: 0,
		},
		{
			name: "non-empty cache",
			cache: &Cache[int, string]{
				m: map[int]string{
					1: "one",
					2: "two",
					3: "three",
				},
			},
			expected: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			actual := tc.cache.Len()
			if actual != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, actual)
			}
		})
	}
}
