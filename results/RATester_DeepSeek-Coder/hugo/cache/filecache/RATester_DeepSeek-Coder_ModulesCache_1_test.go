package filecache

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
	"gotest.tools/assert"
)

func TestModulesCache_1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		caches   Caches
		expected *Cache
	}{
		{
			name: "Returns the modules cache when it exists",
			caches: Caches{
				CacheKeyModules: &Cache{Fs: afero.NewMemMapFs()},
			},
			expected: &Cache{Fs: afero.NewMemMapFs()},
		},
		{
			name: "Returns nil when the modules cache does not exist",
			caches: Caches{
				"other": &Cache{Fs: afero.NewMemMapFs()},
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

			result := tc.caches.ModulesCache()
			assert.Equal(t, tc.expected, result)
		})
	}
}
