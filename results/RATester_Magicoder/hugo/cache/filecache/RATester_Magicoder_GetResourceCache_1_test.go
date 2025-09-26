package filecache

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestGetResourceCache_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cache := &Cache{
		Fs: afero.NewMemMapFs(),
	}

	caches := Caches{
		CacheKeyGetResource: cache,
	}

	result := caches.GetResourceCache()

	if result != cache {
		t.Errorf("Expected %v, got %v", cache, result)
	}
}
