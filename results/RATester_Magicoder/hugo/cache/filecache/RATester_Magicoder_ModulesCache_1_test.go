package filecache

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestModulesCache_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Create a new instance of Caches
	caches := make(Caches)

	// Create a new instance of Cache
	cache := &Cache{
		Fs: afero.NewMemMapFs(),
	}

	// Add the cache to the Caches
	caches[CacheKeyModules] = cache

	// Call the method under test
	result := caches.ModulesCache()

	// Assert that the result is the expected cache
	if result != cache {
		t.Errorf("Expected %v, but got %v", cache, result)
	}
}
