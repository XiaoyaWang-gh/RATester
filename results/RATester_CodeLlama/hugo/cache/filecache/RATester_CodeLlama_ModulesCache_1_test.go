package filecache

import (
	"fmt"
	"testing"
)

func TestModulesCache_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := Caches{}
	f[CacheKeyModules] = &Cache{}
	if f.ModulesCache() != f[CacheKeyModules] {
		t.Errorf("ModulesCache() = %v, want %v", f.ModulesCache(), f[CacheKeyModules])
	}
}
