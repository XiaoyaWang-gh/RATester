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

	configs := Configs{
		CacheKeyModules: FileCacheConfig{
			DirCompiled: "/path/to/cache",
		},
	}

	expected := "/path/to/cache"
	actual := configs.CacheDirModules()

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
