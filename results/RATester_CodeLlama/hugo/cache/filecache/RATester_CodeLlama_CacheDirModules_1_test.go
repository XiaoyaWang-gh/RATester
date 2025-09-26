package filecache

import (
	"fmt"
	"testing"
)

func TestCacheDirModules_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := Configs{}
	c[CacheKeyModules] = FileCacheConfig{
		Dir:         "test",
		DirCompiled: "test",
	}
	if got := c.CacheDirModules(); got != "test" {
		t.Errorf("CacheDirModules() = %v, want %v", got, "test")
	}
}
