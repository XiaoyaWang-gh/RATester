package filecache

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestAssetsCache_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Create a new Caches instance
	caches := make(Caches)

	// Create a new Cache instance
	cache := &Cache{
		Fs: afero.NewMemMapFs(),
	}

	// Add the cache to the Caches instance
	caches[CacheKeyAssets] = cache

	// Call the method under test
	result := caches.AssetsCache()

	// Assert that the result is the expected cache
	if result != cache {
		t.Errorf("Expected %v, but got %v", cache, result)
	}
}
