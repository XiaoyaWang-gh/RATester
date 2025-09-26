package filecache

import (
	"fmt"
	"testing"
)

func TestAssetsCache_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := Caches{}
	f[CacheKeyAssets] = &Cache{}
	if f.AssetsCache() != f[CacheKeyAssets] {
		t.Errorf("AssetsCache() = %v, want %v", f.AssetsCache(), f[CacheKeyAssets])
	}
}
