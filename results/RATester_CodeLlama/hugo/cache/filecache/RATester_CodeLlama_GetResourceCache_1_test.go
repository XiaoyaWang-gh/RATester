package filecache

import (
	"fmt"
	"testing"
)

func TestGetResourceCache_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := Caches{}
	f[CacheKeyGetResource] = &Cache{}
	if got := f.GetResourceCache(); got != f[CacheKeyGetResource] {
		t.Errorf("Caches.GetResourceCache() = %v, want %v", got, f[CacheKeyGetResource])
	}
}
