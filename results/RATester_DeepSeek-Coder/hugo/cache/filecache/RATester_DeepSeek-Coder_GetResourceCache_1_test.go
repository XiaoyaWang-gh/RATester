package filecache

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/spf13/afero"
)

func TestGetResourceCache_1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		caches   Caches
		expected *Cache
	}{
		{
			name: "get resource cache",
			caches: Caches{
				CacheKeyGetResource: &Cache{
					Fs: afero.NewMemMapFs(),
				},
			},
			expected: &Cache{
				Fs: afero.NewMemMapFs(),
			},
		},
	}

	for _, tc := range testCases {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			t.Parallel()

			got := tc.caches.GetResourceCache()
			if diff := cmp.Diff(tc.expected, got); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
