package filecache

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestGetJSONCache_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cache := Caches{
		CacheKeyGetJSON: &Cache{
			Fs: afero.NewMemMapFs(),
		},
	}

	result := cache.GetJSONCache()

	if result != cache[CacheKeyGetJSON] {
		t.Errorf("Expected %v, got %v", cache[CacheKeyGetJSON], result)
	}
}
