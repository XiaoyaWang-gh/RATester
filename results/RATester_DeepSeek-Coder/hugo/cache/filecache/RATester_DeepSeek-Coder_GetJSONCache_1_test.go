package filecache

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestGetJSONCache_1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		caches   Caches
		expected *Cache
	}{
		{
			name: "Returns the JSON cache when it exists",
			caches: Caches{
				CacheKeyGetJSON: &Cache{},
			},
			expected: &Cache{},
		},
		{
			name: "Returns nil when the JSON cache does not exist",
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

			result := tc.caches.GetJSONCache()
			assert.Equal(t, tc.expected, result)
		})
	}
}
