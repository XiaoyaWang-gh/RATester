package filecache

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestGetCSVCache_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cache := Caches{
		CacheKeyGetCSV: &Cache{
			Fs: afero.NewMemMapFs(),
		},
	}

	result := cache.GetCSVCache()

	if result == nil {
		t.Error("Expected a cache, but got nil")
	}
}
