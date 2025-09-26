package filecache

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestGetCSVCache_1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		caches   Caches
		expected *Cache
	}{
		{
			name: "Returns the CSV cache when it exists",
			caches: Caches{
				CacheKeyGetCSV: &Cache{},
			},
			expected: &Cache{},
		},
		{
			name: "Returns nil when the CSV cache does not exist",
			caches: Caches{
				"other": &Cache{},
			},
			expected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.caches.GetCSVCache()
			assert.Equal(t, tc.expected, result)
		})
	}
}
