package filecache

import (
	"fmt"
	"testing"
)

func TestAssetsCache_1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		caches Caches
		want   *Cache
	}{
		{
			name: "AssetsCache exists",
			caches: Caches{
				CacheKeyAssets: &Cache{},
			},
			want: &Cache{},
		},
		{
			name: "AssetsCache does not exist",
			caches: Caches{
				"other": &Cache{},
			},
			want: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := tc.caches.AssetsCache()
			if got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
